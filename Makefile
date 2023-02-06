# 忽略后面的target文件是否存在，直接执行target下面的command命令
.PHONY: all build run gotool clean help

BINARY="webApp"

# 单独输入make时候, 执行all, 就是先执行gotool, 再执行build
all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

# make run 就会执行下面的命令, 加一个@ 符号，这行命令不会在终端输出
run:
	@go run ./

# make gotool 执行下面两个命令
gotool:
	go fmt ./
	go vet ./

# clean : 如果${BINARY}文件存在, 删除之
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# make help 执行下面的命令
help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"