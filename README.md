# CleanC

基于 Wails v3 + Vue 3 的 Windows 系统工具。

## 项目结构

```
cleanC/
├── main.go                          # 应用入口，注册 Service，创建窗口
├── go.mod
├── Taskfile.yml                     # 构建任务（Task Runner）
├── build/
│   ├── config.yml                   # Wails 项目配置（开发模式、应用信息）
│   ├── Taskfile.yml                 # 公共构建任务
│   ├── appicon.png                  # 应用图标
│   └── windows/                     # Windows 平台构建资源
│       ├── Taskfile.yml
│       ├── icon.ico
│       ├── info.json
│       ├── wails.exe.manifest
│       └── nsis/                    # NSIS 安装包脚本
├── internal/                        # 内部服务层
│   └── service/
│       └── systemservice.go         # 系统信息服务
├── pkg/                             # 公共工具包
│   ├── util/
│   │   └── common.go               # 通用工具函数
│   ├── constant/                    # 常量定义
│   └── error/                       # 错误定义
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
        └── views/                   # 页面组件
            └── Home.vue
```

## 分层说明

- **入口层** (`main.go`) — 创建 Application，注册 Service，配置窗口
- **服务层** (`internal/service/`) — 业务逻辑，通过 Wails 绑定暴露给前端
- **工具层** (`pkg/`) — 公共工具函数、常量、错误定义
- **前端层** (`frontend/`) — Vue 3 + Vite + TypeScript

## 开发

```bash
# 启动开发模式（热重载）
wails3 task dev

# 构建生产版本（当前架构）
wails3 task build:amd64
wails3 task build:arm64

# 打包安装程序
wails3 task package:amd64
wails3 task package:arm64

# 重新生成前端绑定
wails3 generate bindings -ts
```

## 依赖

- Go 1.25+
- Node.js 18+
- Wails v3 CLI
- NSIS（打包安装程序时需要）
