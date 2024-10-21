package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type Logger struct {
	Logger zerolog.Logger
	conf   Config
}

type Config struct {
	Dir      string
	Level    string
	MaxAge   int
	Duration int
	MaxSize  int
}

var logger *Logger

func (l *Logger) setlevel() {
	switch strings.ToLower(l.conf.Level) {
	case "debug":
		l.Logger = l.Logger.Level(zerolog.DebugLevel)
	case "warn":
		l.Logger = l.Logger.Level(zerolog.WarnLevel)
	case "error":
		l.Logger = l.Logger.Level(zerolog.ErrorLevel)
	default:
		l.Logger = l.Logger.Level(zerolog.InfoLevel)
	}
}

func findLastLogFile(dir string) (string, bool) {
	var latestFile string
	var latestModTime time.Time
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只考虑普通文件
		if !info.IsDir() {
			if info.ModTime().After(latestModTime) {
				latestFile = path
			}
		}
		return nil
	}); err != nil {
		return "", false
	}

	pices := strings.Split(filepath.Base(latestFile), "-")
	if len(pices) != 2 {
		return "", false
	}

	latestModTime, err := time.ParseInLocation("20060102150405", strings.TrimSuffix(pices[1], filepath.Ext(pices[1])), time.Local)
	if err != nil {
		return "", false
	}

	if time.Now().After(latestModTime) {
		return "", false
	}

	return latestFile, true
}

// 选择是使用那个日志文件
func (l *Logger) loadfile() (string, error) {
	duration := time.Duration(l.conf.Duration) * time.Hour
	oldPath, ok := findLastLogFile(l.conf.Dir)
	if ok {
		return oldPath, nil
	}

	logName := time.Now().Format("20060102150405") + "-" + time.Now().Add(duration).Format("20060102150405") + ".log"
	logPath := filepath.Join(l.conf.Dir, logName)
	file, err := os.Create(logPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return logPath, nil
}

func (l *Logger) setWriter() {
	go func() {
		// 终端输出始终是要有的
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		// 如果这个文件夹不存在
		if _, err := os.Stat(l.conf.Dir); err != nil && os.IsNotExist(err) {
			// 就创建这个文件夹
			if err := os.Mkdir(l.conf.Dir, os.ModePerm); err != nil {
				panic(err)
			}
		}
		tracker := time.NewTicker(time.Hour * time.Duration(l.conf.Duration))
		defer tracker.Stop()
		for {
			logFile, err := l.loadfile()
			if err != nil {
				continue
			}

			// 设置新的输出目标为 lumberjack 管理的日志文件
			logFileWriter := &lumberjack.Logger{
				Filename: logFile,
				MaxAge:   l.conf.MaxAge, // 日志文件的保留天数
				MaxSize:  l.conf.MaxSize,
			}
			multiWriter := io.MultiWriter(logFileWriter, consoleWriter)
			// 更新 logger 输出
			logger.Logger = zerolog.New(multiWriter).With().Timestamp().Logger()
			logger.setlevel()
			<-tracker.C
		}
	}()
}

func SetupLogger(conf Config) {
	if conf.MaxAge <= 0 {
		conf.MaxAge = 30 // 默认30天清除一次日志
	}

	if conf.MaxSize <= 0 {
		conf.MaxSize = 500 // 默认采用500m为最大日志大小
	}

	if conf.Duration <= 0 {
		conf.Duration = 24 // 默认日志间隔为24小时
	}

	logger = &Logger{
		Logger: zerolog.Logger{},
		conf:   conf,
	}

	if strings.TrimSpace(conf.Dir) == "" {
		absPath, err := filepath.Abs("./log")
		if err != nil {
			panic(err)
		}
		logger.conf.Dir = absPath
	}

	logger.setWriter()
}

func GetLogger() *Logger {
	return logger
}

func (l *Logger) Debug(msg string, args ...any) {
	// 使用 strings.Builder 来高效拼接字符串
	var builder strings.Builder
	// 拼接键值对字符串
	for i := 0; i < len(args)-1; i += 2 {
		key, _ := args[i].(string) // 假设每个键都是字符串，简化错误处理
		value := args[i+1]
		builder.WriteString(fmt.Sprintf("%s=%v ", key, value))
	}
	logger.Logger.Debug().Msg(msg + " " + strings.TrimSpace(builder.String()))
}

func (l *Logger) Info(msg string, args ...any) {
	// 使用 strings.Builder 来高效拼接字符串
	var builder strings.Builder
	// 拼接键值对字符串
	for i := 0; i < len(args)-1; i += 2 {
		key, _ := args[i].(string) // 假设每个键都是字符串，简化错误处理
		value := args[i+1]
		builder.WriteString(fmt.Sprintf("%s=%v ", key, value))
	}
	logger.Logger.Info().Msg(msg + " " + strings.TrimSpace(builder.String()))
}

func (l *Logger) Warn(msg string, args ...any) {
	// 使用 strings.Builder 来高效拼接字符串
	var builder strings.Builder
	// 拼接键值对字符串
	for i := 0; i < len(args)-1; i += 2 {
		key, _ := args[i].(string) // 假设每个键都是字符串，简化错误处理
		value := args[i+1]
		builder.WriteString(fmt.Sprintf("%s=%v ", key, value))
	}
	logger.Logger.Warn().Msg(msg + " " + strings.TrimSpace(builder.String()))
}

func (l *Logger) Error(msg string, args ...any) {
	//使用 strings.Builder 来高效拼接字符串
	var builder strings.Builder
	//拼接键值对字符串
	for i := 0; i < len(args)-1; i += 2 {
		key, _ := args[i].(string) // 假设每个键都是字符串，简化错误处理
		value := args[i+1]
		builder.WriteString(fmt.Sprintf("%s=%v ", key, value))
	}
	logger.Logger.Error().Msg(msg + " " + strings.TrimSpace(builder.String()))
}

func Debug(msg string, arg ...any) {
	logger.Debug(msg, arg...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
