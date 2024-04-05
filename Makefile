.PHONY: win64 win64_tools ubuntu_server ubuntu_server_tools clean hello_world

APP_BINARY_WIN=rkvs.exe
APP_BINARY_UBUNTU=rkvs

win64:
	@echo Building for windows/amd64
	@SET GOOS=windows
	@SET GOARCH=amd64
	protoc --go_out=. --go-grpc_out=. ./proto/rkvs.proto
	go build -o ./bin/$(APP_BINARY_WIN) ./cmd/rkvs

win64_tools:
	@echo Building for windows/amd64
	@SET GOOS=windows
	@SET GOARCH=amd64
	go build -o ./bin/install.exe ./cmd/install
	go build -o ./bin/uninstall.exe ./cmd/uninstall

ubuntu_server:
	@echo Building for linux/amd64
	GOOS=linux GOARCH=amd64 protoc --go_out=. --go-grpc_out=. ./proto/rkvs.proto
	GOOS=linux GOARCH=amd64 go build -o ./bin/rkvs ./cmd/rkvs

ubuntu_server_tools:
	@echo Building for linux/amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/install ./cmd/install
	GOOS=linux GOARCH=amd64 go build -o ./bin/uninstall ./cmd/uninstall

clean:
	@echo Cleaning up...
	@if exist ./bin/$(APP_BINARY_WIN) del /Q ./bin/$(APP_BINARY_WIN)
	@if exist ./bin/install.exe del /Q ./bin/install.exe
	@if exist ./bin/uninstall.exe del /Q ./bin/uninstall.exe
	@if [ -f ./bin/$(APP_BINARY_UBUNTU) ]; then rm ./bin/$(APP_BINARY_UBUNTU); fi
	@if [ -f ./bin/install ]; then rm ./bin/install; fi
	@if [ -f ./bin/uninstall ]; then rm ./bin/uninstall; fi

