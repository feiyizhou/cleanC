.PHONY: dev dev-frontend build build-windows build-frontend generate clean

# 完整开发模式（需要 GUI 环境，Windows/macOS 下使用）
dev:
	@wails3 dev

# 仅启动前端开发服务器（Linux 下使用）
dev-frontend:
	@cd frontend && npm run dev

# 本地构建
build: build-frontend
	@go build -o build/bin/cleanc .

# 交叉编译 Windows 版本
build-windows: build-frontend
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/bin/cleanc.exe .

build-frontend:
	@cd frontend && npm install && npm run build

generate:
	@wails3 generate bindings

clean:
	@rm -rf frontend/dist
	@rm -rf build/bin
