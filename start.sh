#!/bin/bash

# 定义常量
CONFIG_FILE="./config.json"                 # 配置文件路径
DOCKER_COMPOSE_FILE="./docker-compose.yml"  # Docker Compose 文件路径
BINARY_PATH="service"                 # 二进制程序路径
PID_FILE="$BINARY_PATH.pid"                # PID 文件路径
LOG_FILE="./service.log"                 # 日志文件路径

# 确保日志目录存在
mkdir -p ./log

# 函数：记录日志
log_message() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $1" >> $LOG_FILE
}



# 函数：启动二进制程序
start() {
    log_message "开始服务..."
    nohup ./$BINARY_PATH -conf $CONFIG_FILE > $LOG_FILE 2>&1 &
    PID=$!
    echo $PID > $PID_FILE
    log_message "服务已启动，PID: $PID"
    log_message "PID 已保存到 $PID_FILE"
}

# 函数：停止二进制程序
stop() {
    if [ -f $PID_FILE ]; then
        PID=$(cat $PID_FILE)
        log_message "停止服务，PID: $PID"
        kill -9 $PID
        rm -f $PID_FILE
        log_message "服务已停止"
    else
        log_message "PID 文件不存在，无法停止服务"
    fi
}

# 主程序
main() {
    case "$1" in
        start)
            start
            ;;
        stop)
            stop
            ;;
        restart)
            stop
            start
            ;;
        *)
            echo "用法: $0 {start|stop|restart}"
            exit 1
            ;;
    esac
}

# 执行主程序
main "$@"
