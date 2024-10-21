# 创建一个新的轻量级镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY ./pack/service /app/service 

# 确保二进制文件可执行
RUN chmod +x /app/service

# 运行二进制文件，指定配置文件
CMD ["./service", "-conf", "./config.json"]
