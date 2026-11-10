package main

import "github.com/egoist/mygo/ui"

// toastMsg is a queued toast: some sources (the run goroutine) have no
// build context, so toasts drain at the top of the view.
type toastMsg struct {
	title, desc string
	bad         bool
}

func (a *app) toastOK(title string) {
	a.toasts = append(a.toasts, toastMsg{title: title})
}

func (a *app) toastErr(title string) {
	a.toasts = append(a.toasts, toastMsg{title: title, bad: true})
}

// toastIcon is the Lucide mark Sonner shows for a toast type.
func toastIcon(ty string) *ui.SVG {
	switch ty {
	case "success":
		return iconCircleCheck
	case "error":
		return iconOctagonX
	case "warning":
		return iconTriangle
	default:
		return iconInfo
	}
}

// drainToasts fires queued toasts and builds the toast viewport, bottom
// right as on the web. The look follows Sonner: a 356px popover card with
// 16px padding, a 16px Lucide mark, and no close button (toasts dismiss
// themselves, as Sonner's do by default).
func (a *app) drainToasts(c *ui.Context, p palette) {
	for _, t := range a.toasts {
		ty := "success"
		if t.bad {
			ty = "error"
		}
		c.AddToast(ui.Toast{Title: t.title, Description: t.desc, Type: ty})
	}
	a.toasts = nil
	ui.ToastViewportBase(c, func(viewport ui.Element, toasts []ui.Toast) {
		viewport.Padding(32).AlignItems(ui.End).Gap(14)
		for _, t := range toasts {
			toast := ui.ToastBase(c, t)
			toast.Root.Row().Gap(6).Padding(16).Radius(8).Width(356).
				AlignItems(ui.Center).Background(p.popover.color()).
				Border(1, p.border.color()).
				Shadow(0, 4, 12, 0, ui.RGBA(0, 0, 0, 0.1))
			toast.Root.Children(func() {
				// Sonner's icon box: 16px, pulled 3px left and 4px from the
				// content, which the row's 6px gap adds to.
				ui.Icon(c, toastIcon(t.Type)).Size(16, 16).
					TextColor(p.foreground.color()).Margin(0, 4, 0, -3)
				if t.Description == "" {
					ui.Text(c, t.Title).FontSize(13).FontWeight(500).
						Grow(1).Basis(0).MinWidth(0).
						TextColor(p.foreground.color())
				} else {
					desc := ui.RGBA(0x3f, 0x3f, 0x3f, 1)
					if p.isDark {
						desc = ui.RGBA(0xe8, 0xe8, 0xe8, 1)
					}
					ui.Column(c).Grow(1).Basis(0).MinWidth(0).Gap(2).Children(func() {
						ui.Text(c, t.Title).FontSize(13).FontWeight(500).
							TextColor(p.foreground.color())
						ui.Text(c, t.Description).FontSize(13).
							TextColor(desc)
					})
				}
			})
		}
	})
}
