package window

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// Component implements widget with logic.
type Component interface {
	/// UpdateFrame implements update frame logic.
	UpdateFrame(gtx layout.Context) error
	/// UpdateEvent implements update event logic.
	UpdateEvent(evt interface{}) error
	/// Render implements draw widget.
	Render(gtx layout.Context, theme *material.Theme) layout.Dimensions
}
