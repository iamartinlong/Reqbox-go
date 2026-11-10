# Reqbox Go

Reqbox 的原生桌面版：mygo 原生 UI（Go 直接绘制，无 WebView），与 Reqbox 共用
SQLite 数据格式。2026-11-10 从
[iamartinlong/mygo-example](https://github.com/iamartinlong/mygo-example) 的
`reqbox-native/` 分离到本仓库（v0.1.2）。

> 显示名（窗口标题、Windows 快捷方式/安装器/应用列表）为 **Reqbox Go**；可执行文件、
> 安装目录、slug 与用户数据目录仍沿用 `Reqbox Native` / `reqbox-native`，以保证已有
> 安装的数据与自动更新连续。详见 [TODO.md](TODO.md) 的「有意差异」。

> 剩余待办、已知差距与有意差异见 [TODO.md](TODO.md)。

## 功能

- 请求编辑：方法/URL、Params / Headers / Form KV、Body、Auth、cURL 导入
- 响应面板：状态行、Headers / Request / Preview / Raw、分阶段计时、自增长编辑器
- 历史、发布/Mock、环境管理与 `{{var}}` 替换、导入导出
- 明暗主题、en/zh、快捷键（Ctrl+Enter 发送、Ctrl+S 保存、Ctrl+N 新建）
- 打包与更新：签名构建、NSIS 安装器、GitHub Releases 增量更新（delta）

## 运行 / 构建

```sh
go run .          # 开发运行（Go ≥ 1.27.3；低版本 GOTOOLCHAIN=auto 自动拉取）
go test ./...     # pkg/store 与 pkg/runner 的测试
```

签名构建（发布用；密钥在仓库外，不入库）：

```sh
MYGO_UPDATER_PRIVATE_KEY="$(cat "$APPDATA/mygo/update-keys/reqbox-native/mygo-update.key")" \
  go run github.com/egoist/mygo/cmd/mygo build
```

产物在 `build/windows-amd64/`；`update-windows-amd64.json` 由构建写入并签名，
安装版据此更新。不传私钥也能出包，但安装版收不到更新。

## 更新源与迁移

`mygo.json` 的 `updates.github` 指向本仓库，`tagPrefix` 为 `v`。**0.1.2 及更早的
安装版是从旧仓库发布的（feed 指向 mygo-example），需要手动安装一次本仓库的版本**，
之后由本仓库自动更新。v0.2.0 是本仓库首个发布（无 delta）；v0.2.1 起恢复增量更新
（首个 delta 185 KB）。

## 目录

```
main.go / *.go      应用与各面板（即时模式 ui）
pkg/store           SQLite 存储（请求/集合/历史/环境/设置），与旧仓库同源
pkg/runner          HTTP 执行器（变量替换、超时、重定向、TLS 选项）
icons/ fonts/       图标与内嵌字体
mygo.json           打包与更新配置
```

## 沿革

- 2026-11-05 路线 A：`reqbox-desktop`（webview 复用，100% 保真基线，仍在旧仓库）
- 2026-11-06 起路线 B：原生重写，批次与对照记录见旧仓库
  `docs/01-research/reqbox-migration.md`
- 2026-11-10 发版 v0.1.2（首个 delta 22KB），同日迁到本仓库
- 2026-11-10 v0.2.0：本仓库首个发布（同 0.1.2 代码，切换更新源；无 delta）
- 2026-11-10 v0.2.1：显示名改为 Reqbox Go；本仓库首个 delta（0.2.0→0.2.1，23 KB）
