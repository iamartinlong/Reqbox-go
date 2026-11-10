# 待办

> 更新：2026-11-10。当前版本 **v0.2.0**（本仓库首个发布）。完成一项就勾掉并补一句结果；
> 文末「有意差异」是设计取舍，别当缺陷改。

## 立即（迁移收尾）

- [x] **手动安装一次 v0.2.0**：已装在 `D:\soft\Reqbox Native`（FileVersion 0.2.0；
      开始菜单与桌面快捷方式均指向此处，2026-11-10 核对），之后更新走本仓库
- [x] **清理遗留的 0.1.2 安装目录**：`%LOCALAPPDATA%\Programs\Reqbox Native` 已整体删除
      （2026-11-10，含 `.old-*` 更新备份与 Uninstall.exe），快捷方式均已指向 0.2.0
- [x] **旧仓库已推送**：迁移标注/知识库同步/待办入口/收尾交接 4 次提交已推
      （`6a1bb39`…`862b5fb`），`mygo-example` 的 `main` 与 `origin/main` 齐平（2026-11-10）
- [ ] **0.2.1 发版时**：确认构建生成 `0.2.0→0.2.1` delta（本仓库首次带 delta），并做一次应用内升级验证
- [x] **v0.2.0 Release 资产已核对**：`reqbox-native-setup-0.2.0-amd64.exe`、
      `reqbox-native-0.2.0-windows-amd64.tar.gz`、`update-windows-amd64.json`（签名）、
      `SHA256SUMS`、`install.ps1` 共 5 个附件齐全（2026-11-10）

## 功能收尾（非阻塞，按需）

- [ ] **历史重放保存实际发送值**：现在 `history` 存的是发送前草稿快照，重放时请求头/请求体与当时实际发送的不完全一致；
      需给共享 store 的 history 表加字段（影响旧仓库 webview 版 DB schema，单独评估后再动）
- [ ] **等待骨架 pulse 动画**：目前是静态骨架，mygo 有动画能力可补
- [ ] **toast 图标细节对齐 sonner**：形状已对，图标与间距未逐像素

## 维护

- [ ] **`pkg/store`、`pkg/runner` 双份同步**：本仓库一份、旧仓库 `reqbox-desktop` 一份；共享逻辑改动要两边一起改
- [ ] **评估消除双份**：让旧仓库 webview 版改为依赖本仓库模块，或把共享包抽成独立模块
- [ ] 若以后要预发布（beta），在插件侧配 `SetPrereleases`
- [ ] **每次发版核对 Release 附件**：安装器、整包、签名 `update-windows-amd64.json`、
      `SHA256SUMS` 齐全；带 delta 的版本另核对增量包能被旧版拉取

## 有意差异（别当缺陷改）

- KV 行 id：web 用 `crypto.randomUUID()`，原生用 8 字节随机 hex——都唯一，库间互换无碍
- Mock URL：原生显示网关实际端口（web 显示页面 origin；5178 被占时两者不同）
- CSS 百分比基准差 ~19px：原生拿不到窗口高，`MaxHeightPercent` 按父容器解析，滚动起点略早
- 桌面单用户：没有 web 的登录/token 流与 chi API 服务；store/runner 进程内直连，mock 用真实监听

## 参考

- 完整批次记录、web 基线对照数据：旧仓库 `docs/01-research/reqbox-migration.md`
- 更新链路原理与多轮 E2E（含 delta、回退演练）：旧仓库 `docs/04-verification/update-e2e.md`
- 框架级待办（IME 组字人工实测、macOS/Linux 真机、Azure Trusted Signing 证书）：
  旧仓库 `docs/01-research/mygo-production-gaps.md`
- 发版流程见 [README](README.md)「运行 / 构建」；签名密钥在仓库外
  `%APPDATA%/mygo/update-keys/reqbox-native/mygo-update.key`，不入库
