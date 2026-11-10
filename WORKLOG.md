# 工作日志

> 最后更新：2026-11-10。当前版本 **v0.2.5**，`main` 与 `origin/main` 同步，CI 绿。
> 本文件记录交接所需的状态、决策与验证；待办清单见 [TODO.md](TODO.md)。

## 2026-11-10 交接摘要

本轮从迁移收尾开始，完成 0.2.0 → 0.2.5 六个发布（0.2.0 为本仓库首发），主要工作：

1. **迁移收尾**：安装 0.2.0、清理遗留 0.1.2 目录、桌面/开始菜单快捷方式改为 Reqbox Go。
2. **品牌改名**：对外显示名改为 **Reqbox Go**；内部标识、可执行文件名、安装目录与 slug
   保持 `Reqbox Native` / `reqbox-native`，用户数据目录固定回 `%APPDATA%\Reqbox Native`，
   保证已有安装的数据与自动更新连续。
3. **标题栏版本号**：应用名后显示小字 `v0.2.5`，由构建注入的版本自动生成。
4. **CI**：GitHub Actions 在 push/PR 跑 `go build`、`go vet`、`mygo vet`、`go test`
   与无签名打包冒烟。
5. **大工作区性能（0.2.3/0.2.4）**：
   - 侧栏改读 `RequestSummary`，集合计数缓存，筛选大小写预计算；发送后只刷新历史
   - 10k 请求加载：521ms / 114MB → 76–103ms / 9MB
   - 每帧 100 个集合计数：21.7ms → ~5µs
   - 侧栏虚拟化（`ui.List`，不设 `Selected` 以保留行内 kebab/折叠按钮点击）
   - 单条请求变更就地更新：10k 下保存约 5.9ms（原全量 80–100ms）
6. **UI 对齐（0.2.5）**：骨架 `animate-pulse`、Sonner toast（16px Lucide 图标、356px 宽、
   16px 内边距、32px 视口偏移、无关闭按钮）、mygo 内置文案中文化。

## 当前状态

- **仓库/分支**：`iamartinlong/Reqbox-go`，`main`，最新提交 `1e563b0`；工作区干净。
- **版本/发布**：`mygo.json` 为 0.2.5；GitHub Release `v0.2.5` 已发布（非草稿），8 个附件；
  latest feed 已指向 0.2.5。
- **许可证**：MIT，见 `LICENSE`。
- **安装版**：`D:\soft\Reqbox Native\Reqbox Native.exe`，FileVersion **0.2.5**；
  桌面与开始菜单快捷方式为 `Reqbox Go.lnk`。
- **用户数据**：`%APPDATA%\Reqbox Native\reqbox.db`（SQLite）与同目录 `updater.json`。
- **更新源**：`mygo.json` 的 `updates.github` 指向本仓库，`tagPrefix` 为 `v`，Ed25519
  公钥内置；签名私钥在仓库外
  `%APPDATA%\mygo\update-keys\reqbox-native\mygo-update.key`，不入库。
- **CI**：`.github/workflows/ci.yml`，`windows-latest`；基准测试不在 CI 默认运行。
- **工具**：本机没有 `gh` CLI。发版当前用 GitHub API 创建 draft、上传附件、发布（API
  token 从 `git credential fill` 读取），或安装 `gh` 后用 `mygo build -upload`。

## 版本记录（2026-11-10）

| 版本 | 内容 | 增量包（旧版 → 该版） |
| --- | --- | --- |
| 0.2.0 | 本仓库首发，切换更新源；无 delta | — |
| 0.2.1 | 显示名改 Reqbox Go；本仓库首个 delta | 0.2.0 → 185 KB |
| 0.2.2 | 更新弹窗名称同步；数据目录固定 | 0.2.1 → 384 KB；0.2.0 → 368 KB |
| 0.2.3 | 大工作区加载/计数/筛选优化 | 0.2.2 → 420 KB；0.2.1 → 501 KB；0.2.0 → 499 KB |
| 0.2.4 | 侧栏虚拟化、增量刷新、标题栏版本号 | 0.2.3 → 400 KB；0.2.2 → 470 KB；0.2.1 → 566 KB |
| 0.2.5 | 骨架 pulse、Sonner toast、mygo 中文文案 | 0.2.4 → 566 KB；0.2.3 → 591 KB；0.2.2 → 636 KB |

## 关键改动与文件

- `main.go`：`App.SetName("Reqbox Go")`、`PathUserData` 固定、标题栏版本号、
  `locales/zh-Hans` 空导入。
- `sidebar.go`：`RequestSummary`、`colCounts`、`reqSearch`、`applyRequest` /
  `removeRequest` 增量刷新、`ui.List` 侧栏虚拟化与 `rowKey`。
- `pkg/store/requests.go`：新增 `ListRequestSummaries`（侧栏轻量读取）。
- `toast.go` + `icons/*.svg`：Sonner 对齐的图标、尺寸、间距与配色。
- `response.go`：等待骨架的 `Keyframes` + `RepeatForever` pulse。
- `pkg/store/bench_test.go`、`perf_bench_test.go`：数据层与侧栏基准。
- `.github/workflows/ci.yml`：CI 验证。
- `CHANGELOG.md`：发布说明；`TODO.md`：待办。

## 验证记录

- 每个发版提交的 CI 均为 success。
- 每个 Release 均核对 8 个附件（安装器、整包、3 个 delta、`update-windows-amd64.json`、
  `SHA256SUMS`、`install.ps1`）；latest feed 与 delta SHA256 均与本地构建一致。
- 应用内升级：0.2.0 → 0.2.2、0.2.3 → 0.2.4 已验证；0.2.4 → 0.2.5 已完成安装
  （安装版 FileVersion 0.2.5）。
- 侧栏：2000 请求 / 50 集合的合成工作区，验证滚动、集合折叠、kebab 菜单、行点击；
  10k/50k 数据见基准测试。
- 骨架与 toast：本地 3 秒延迟接口实机截图确认 pulse 明暗变化；toast 截图放大核对图标、
  文字、圆角与间距。
- 已知缺口：GUI 滚动帧率未做量化测量；基准未纳入 CI。

## 交接注意 / 风险

- **不要改内部身份**：`mygo.json` 的 `name` / `identifier`、可执行文件名、安装目录与
  数据目录是更新与数据连续性的基础；对外名称只走 `windows.displayName`、
  `App.SetName` 和 UI 文案。
- **侧栏 `ui.List` 故意不设 `Selected`**：行内 kebab 与折叠按钮需要点击；若设置
  `Selected`，这些控件会被行选择吃掉。
- **增量刷新范围**：只覆盖保存/新建/删除/复制/重命名/发布单条请求；集合增删、导入与
  启动仍走 `refresh()` 全量刷新。
- **共享包双份**：`pkg/store`、`pkg/runner` 在旧仓库 `reqbox-desktop` 仍有一份，共享逻辑
  改动需两边同步。
- **发布必须带私钥**：不设置 `MYGO_UPDATER_PRIVATE_KEY` 的构建装得上，但收不到自动更新。
- **发布后核对**：附件、latest feed、delta SHA256；发布为 draft 时 `/releases/tags` 查不到，
  需按 release id 查询。

## 新会话入口

1. 先读 `README.md`（项目与构建）、`TODO.md`（待办）、本文件（状态与交接）。
2. 开发/检查：

   ```sh
   go run .
   go build ./...
   go vet ./...
   go run github.com/egoist/mygo/cmd/mygo vet .
   go test ./...
   ```

3. 性能基准（手动）：

   ```sh
   go test ./pkg/store -run '^$' -bench List -benchmem
   go test . -run '^$' -bench RebuildRows -benchmem
   ```

4. 发版：改 `mygo.json` 版本与 `CHANGELOG.md` → 签名构建 → 提交并打 `v<version>` tag →
   创建 GitHub Release 并上传 `build/windows-amd64/` 的附件与 `build/SHA256SUMS` →
   核对 latest feed 与 delta。
