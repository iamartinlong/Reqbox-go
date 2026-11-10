# Changelog

## 0.2.1 — 2026-11-10

Maintenance release: the app now shows as **Reqbox Go** — in the window
title, the Windows shortcuts, installer and app entry. The executable,
install folder, slug and user data directory keep their `Reqbox Native` /
`reqbox-native` names, so existing installs keep their data and update in
place.

It also carries the first delta update from this repository
(`0.2.0 → 0.2.1`), so an installed 0.2.0 downloads only what changed
instead of the whole archive.

## 0.2.0 — 2026-11-10

First release from the app's own repository: the same code as 0.1.2, moved
out of the mygo-example monorepo (see the migration record there). The
update feed now points at this repository, so install this version once by
hand if you run 0.1.2 or earlier; automatic updates resume after that.

No delta update this time: this repository has no 0.1.2 release to diff
against. The next release will carry one for 0.2.0.

## 0.1.2 — 2026-11-10

Maintenance release: the app is unchanged from 0.1.1. It exists to exercise
the update chain end to end, and carries the first delta update
(`0.1.1 → 0.1.2`), so an installed 0.1.1 downloads only what changed
instead of the whole archive.

## 0.1.1 — 2026-11-09

Native UI port of Reqbox (mygo ui), aligned to the web baseline panel by
panel (batches 19–31 of the migration record):

- Request editor: KV grid (params/headers/form), Body/Auth panels, publish
  panel, tab panel scrolling, xl side-by-side layout.
- Response pane rebuilt on the web's cF: status bar, four line tabs
  (headers/request/preview/raw), collapsible General/headers/timing, editors
  that grow with the body.
- History tab with its own search and relative times; sidebar tree with
  foldable collections, kebab menus and published dots; empty state on the
  web's SF.
- Dialogs on the shadcn baseline: confirm, prompt, curl, run settings, and
  the 768px two-pane environment manager (save visibility fixed).
- Keyboard paths: input focus rings and a keyboard-reachable row delete.
- Dark window: DWM frame border follows the app theme (no white edge in
  dark mode).
- Packaging: mygo.json, signed updates (tagPrefix `reqbox-native-v`) and the
  NSIS installer.

## 0.1.0 — 2026-11-05

First native shell: theme base, sidebar, request bar, send pipeline,
response, history, mocks, environments, dialogs, i18n, settings, dark mode.
