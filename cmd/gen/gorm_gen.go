package main

import (
	"flag"

	"github.com/sirius2001/loon/config"
	"github.com/sirius2001/loon/pkg/db"
	"github.com/sirius2001/loon/pkg/log"

	"gorm.io/gen"
)

var confPath = flag.String("conf", "./config.json", "path/to/your/config.json")

func main() {
	flag.Parse()
	if confPath == nil {
		panic(flag.PanicOnError)
	}

	if err := config.LoadConfig(*confPath); err != nil {
		panic(err)
	}

	log.SetupLogger(log.Config(config.Conf().Log))

	if err := db.NewDB(config.Conf().Merge, config.Conf().DB.DSN); err != nil {
		panic(err)
	}

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../pkg/persistence",                                            // 生成代码的输出目录
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	})

	// 设置数据库连接
	g.UseDB(db.GetDB())

	// 将 NotifyChannel 模型转换为数据库操作方法
	//	g.ApplyBasic(model.User{})
	// 执行生成
	g.Execute()
}
