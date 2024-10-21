# 定义变量
PACK_PATH =./cmd/server/
Config  = ./config/config.json
BINARY = service
OUTPUT_DIR = pack
GIT_TAG := $(shell git describe --tags --abbrev=0)
ZIP_FILE = $(BINARY)_$(GIT_TAG).zip
# 默认目标
all: build
build:
	@rm -rf $(OUTPUT_DIR) 
	@go build  -o $(BINARY) $(PACK_PATH)
	@mkdir -p $(OUTPUT_DIR)
	@mv $(BINARY) $(OUTPUT_DIR)/
	@cp $(Config) $(OUTPUT_DIR)/ 
	@echo "Build . Output located in $(OUTPUT_DIR)/."
# 构建目标
build/linux:
	GOOS=linux GOARCH=amd64  go build  -o $(BINARY) $(PACK_PATH)
	@rm -rf $(OUTPUT_DIR) 
	@mkdir -p $(OUTPUT_DIR)
	@mv $(BINARY) $(OUTPUT_DIR)/
	@cp $(Config) $(OUTPUT_DIR)/
	@echo "Build Linux. Output located in $(OUTPUT_DIR)/."
# Windows 打包目标
build/windows: 
	GOOS=windows GOARCH=amd64  go build  -o $(BINARY).exe $(PACK_PATH)
	@rm -rf $(OUTPUT_DIR) 
	@mkdir -p $(OUTPUT_DIR)
	@mv $(BINARY).exe $(OUTPUT_DIR)/
	@cp $(Config) $(OUTPUT_DIR)/
	@echo "Build Windows. Output located in $(OUTPUT_DIR)/."

pack:build/linux
	@echo "Creating zip file for $(OUTPUT_DIR)..."
	@zip -r $(ZIP_FILE) $(OUTPUT_DIR)/*
	@echo "Zip file created: $(ZIP_FILE)"
	@rm -rf ./pack

upload:build/linux
	docker build -t $(BINARY):latest .
	docker login --username=aliyun4541033394 crpi-ldcb2d9ge4cmfxso.cn-hangzhou.personal.cr.aliyuncs.com
	docker tag  $(BINARY):latest crpi-ldcb2d9ge4cmfxso.cn-hangzhou.personal.cr.aliyuncs.com/sirius-hub/hub:latest
	docker push crpi-ldcb2d9ge4cmfxso.cn-hangzhou.personal.cr.aliyuncs.com/sirius-hub/hub:latest
# 清理目标
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR) $(BINARY) $(ZIP_FILE)
	@echo "Clean up completed."

.PHONY: all build pack  upload clean
