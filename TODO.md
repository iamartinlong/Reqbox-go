# 待办

> 更新：2026-11-10。当前版本 **v0.2.5**（已发布，安装版已升级）。
> 详细状态、版本记录与交接说明见 [WORKLOG.md](WORKLOG.md)。
> 完成一项就勾掉并补一句结果；文末「有意差异」是设计取舍，别当缺陷改。

## 功能

- [ ] **历史重放保存实际发送值**：现在 `history` 存的是发送前草稿快照，重放时请求头/请求体
      与当时实际发送的不完全一致；需给共享 store 的 history 表加字段（影响旧仓库 webview
      版 DB schema，单独评估后再动）
- [ ] （可选）**GUI 滚动帧率量化**：侧栏虚拟化已用 2000 请求实机冒烟，但未测 FPS；
      若遇到 10k+ 请求滚动卡顿，先量帧时间再决定是否继续优化
- [ ] （可选）**性能基准纳入 CI/夜间任务**：目前 `go test -bench` 只手动运行

## 维护与发版

- [ ] **每次发版核对 Release 附件**：安装器、整包、签名 `update-windows-amd64.json`、
      `SHA256SUMS` 齐全；带 delta 的版本另核对增量包能被旧版拉取
- [ ] **`pkg/store`、`pkg/runner` 双份同步**：本仓库一份、旧仓库 `reqbox-desktop` 一份；
      共享逻辑改动要两边一起改
- [ ] **评估消除双份**：让旧仓库 webview 版改为依赖本仓库模块，或把共享包抽成独立模块
- [ ] 若以后要预发布（beta），在插件侧配 `SetPrereleases`
- [ ] （可选）**发布自动化**：本机没有 `gh` CLI；可安装 gh 后用 `mygo build -upload`，
      或把现有 GitHub API 发版流程脚本化

## 本轮已完成

- [x] 迁移收尾：安装 0.2.0、清理遗留 0.1.2 目录、快捷方式改名、0.2.0 → 0.2.5 发布链
- [x] 对外显示名 Reqbox Go（内部身份与数据目录不变，保证数据与自动更新连续）
- [x] 更新弹窗名称同步；标题栏显示小字版本号
- [x] CI 接入：build、vet、mygo vet、test 与无签名打包冒烟
- [x] 大工作区性能：摘要加载、集合计数缓存、筛选预计算、侧栏虚拟化、增量刷新
- [x] 骨架 pulse 动画、Sonner toast 对齐、mygo 内置文案中文化

版本记录、性能数据与验证细节见 [WORKLOG.md](WORKLOG.md)。

## 有意差异（别当缺陷改）

- 显示名 Reqbox Go：窗口标题、更新弹窗、Windows 快捷方式/安装器/应用列表用新名；可执行文件、
  安装目录、slug 与用户数据目录仍是 `Reqbox Native` / `reqbox-native`，保证已有安装
  的数据与自动更新连续（改名只动 `windows.displayName`、启动时 `App.SetName` 与 UI 文案）
- KV 行 id：web 用 `crypto.randomUUID()`，原生用 8 字节随机 hex——都唯一，库间互换无碍
- Mock URL：原生显示网关实际端口（web 显示页面 origin；5178 被占时两者不同）
- CSS 百分比基准差 ~19px：原生拿不到窗口高，`MaxHeightPercent` 按父容器解析，滚动起点略早
- 桌面单用户：没有 web 的登录/token 流与 chi API 服务；store/runner 进程内直连，mock 用真实监听

## 参考

- 本仓库当前状态与交接：`WORKLOG.md`
- 完整批次记录、web 基线对照数据：旧仓库 `docs/01-research/reqbox-migration.md`
- 更新链路原理与多轮 E2E（含 delta、回退演练）：旧仓库 `docs/04-verification/update-e2e.md`
- 框架级待办（IME 组字人工实测、macOS/Linux 真机、Azure Trusted Signing 证书）：
  旧仓库 `docs/01-research/mygo-production-gaps.md`
- 发版流程见 [README](README.md)「运行 / 构建」；签名密钥在仓库外
  `%APPDATA%/mygo/update-keys/reqbox-native/mygo-update.key`，不入库
