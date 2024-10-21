# 部署文档

## 目录结构

```cmd
.
├── config.json         # 服务配置文件
├── docker-compose.yml   # Docker Compose 配置文件
├── service              # 二进制可执行文件
└── start.sh             # 启动脚本
```

- `config.json`: 配置文件，确保其内容正确。
- `docker-compose.yml`: Docker Compose 配置文件。
- `service`: 你的二进制文件，确保它是可执行的。
- `start.sh`: 启动脚本。

## 构建与运行

### 使用 Docker Compose 启动服务

1. **构建镜像**:

   ```shell
   docker-compose build
   ```

2. **启动服务**:

   ```shell
   docker-compose up -d
   ```

   这将会在后台启动服务。可以使用以下命令查看服务状态：

   ```sh
   docker-compose ps
   ```

3. **查看日志**: 可以使用以下命令查看服务日志：

   ```
   docker-compose logs -f
   ```

### 使用启动脚本

1. **赋予启动脚本可执行权限**:

   ```shell
   chmod +x start.sh
   ```

2. **运行启动脚本**:

   ```shell
   ./start.sh start.sh
   ```

   启动脚本将会使用 Docker Compose 启动服务，并在后台运行。

## 停止服务

1. 使用 Docker Compose 停止服务:

   ```shell
   docker-compose down
   ```

   这将停止并删除所有与服务相关的容器。

## 配置文件说明

### `config.json`

- **log**: 日志配置。
  - `level`: 日志级别（如 info、error）。
  - `dir`: 日志文件目录。
  - `max_age`: 最大保留天数。
  - `duration`: 轮转频率。
  - `max_size`: 最大文件大小（MB）。
- **web**: Web 服务器配置。
  - `enable`: 是否启用 Web 服务器。
  - `addr`: 监听地址（如 `0.0.0.0:8081`）。
- **grpc**: gRPC 服务器配置。
  - `enable`: 是否启用 gRPC 服务器。
  - `addr`: 监听地址（如 `0.0.0.0:8181`）。
- **kafka**: Kafka 配置。
  - `addr`: Kafka 地址（如 `127.0.0.1:9091`）。
  - `topic`: 使用的主题名称。