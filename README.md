# CleanC

基于 Wails v3 + Vue 3 的 Windows 系统工具。

## 项目结构

```
cleanC/
├── main.go                          # 应用入口，注册 Service，创建窗口
├── go.mod
├── Makefile                         # 构建任务
├── build/
│   └── config.yml                   # Wails 项目配置（开发模式、应用信息）
├── services/                        # 后端服务层
│   └── system/
│       └── system.go                # 系统信息服务
├── internal/                        # 内部工具包
│   └── util/
│       └── util.go
└── frontend/                        # Vue 3 前端
    ├── index.html
    ├── package.json
    ├── vite.config.ts
    ├── tsconfig.json
    ├── bindings/                    # Wails 自动生成的绑定（勿手动编辑）
    └── src/
        ├── main.ts
        ├── App.vue
        ├── style.css
        ├── views/                   # 页面组件
        │   └── Home.vue
        └── components/              # 通用组件
            └── Greet.vue
```

## 分层说明

- **入口层** (`main.go`) — 创建 Application，注册 Service，配置窗口
- **服务层** (`services/`) — 业务逻辑，通过 Wails 绑定暴露给前端
- **内部工具层** (`internal/`) — 不对外暴露的辅助函数
- **前端层** (`frontend/`) — Vue 3 + Vite + TypeScript

## 开发

```bash
# 安装前端依赖
cd frontend && npm install && cd ..

# 生成前端绑定
make generate

# 启动开发模式（热重载）
make dev

# 构建生产版本
make build

# 清理构建产物
make clean
```

## 依赖

- Go 1.21+
- Node.js 18+
- Wails v3 CLI
