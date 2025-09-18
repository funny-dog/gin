# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 开发命令

### 构建和测试
```bash
# 运行测试 (生成覆盖率报告)
make test

# 代码格式化
make fmt

# 检查代码格式 (不修改)
make fmt-check

# 静态代码分析
make vet

# 代码检查 (需要golint工具)
make lint

# 检查拼写错误
make misspell-check

# 安装工具
make tools
```

### 代码质量工具
项目使用 `golangci-lint` 进行静态代码分析，配置在 `.golangci.yml` 中：
- 启用的 linter: `golangci-lint`
- 包含 `revive`, `gosec`, `testifylint` 等现代 Go 工具
- 格式化工具: `gofmt`, `goimports`, `gofumpt`, `gci`

## 架构概览

### 核心组件

Gin 是一个高性能的 Go Web 框架，主要组件包括：

**1. 引擎 (`gin.go`)** - 核心引擎实现
- Engine 结构体包含路由、中间件、上下文池等
- 实现了 `gin.Engine` 和 `gin.IRouter` 接口
- 支持多种 HTTP 协议 (HTTP/1.1, HTTP/2, HTTP/3)

**2. 上下文 (`context.go`)** - 请求上下文
- `gin.Context` 提供请求处理的所有方法
- 继承自 `gin.Context`，提供 JSON、XML、HTML 等响应方法
- 支持中间件链和请求参数绑定

**3. 路由系统**
- **树形路由 (`tree.go`)** - 基于 radix tree 的路由匹配
- **路由组 (`routergroup.go`)** - 支持路由分组和中间件
- 基于 `httprouter` 的高效路由匹配

**4. 绑定系统 (`binding/`)** - 数据绑定
- JSON、XML、YAML、Form、Multipart 等数据绑定
- 集成了 `go-playground/validator` 进行数据验证
- 支持自定义绑定器

**5. 渲染系统 (`render/`)** - 响应渲染
- JSON、XML、HTML、YAML、TOML 等格式渲染
- 支持多种 JSON 引擎 (`sonic`, `jsoniter`, `std`)
- 自定义渲染器支持

**6. 中间件框架**
- 内置 `recovery` 和 `logger` 中间件
- 支持链式中间件调用
- 错误处理机制

### 模块化设计

```
gin/
├── binding/     - 数据绑定和验证
├── codec/       - 编码器 (JSON, YAML等)
├── render/      - 渲染器
└── internal/    - 内部工具
```

### 性能特点

- 零分配路由器
- 基于 `httprenchmark` 的高性能路由匹配
- 内置连接池优化
- 支持 HTTP/2 和 HTTP/3 协议

### 依赖管理

- Go 1.23+ 版本要求
- 使用 Go Modules 管理依赖
- 核心依赖：`sonic` (高性能JSON), `validator`, `goccy/go-json`

### 测试策略

- 单元测试覆盖所有核心组件
- 集成测试 (`gin_integration_test.go`)
- 基准测试 (`benchmarks_test.go`)
- 示例测试 (`examples/` 目录)

## 关键设计模式

1. **链式调用** - 支持流畅的 API 设计
2. **中间件模式** - 请求处理管道
3. **上下文传播** - 请求生命周期管理
4. **路由树** - 高效的路由匹配算法
5. **接口抽象** - 可扩展的组件设计