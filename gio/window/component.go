package window

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// Component implements widget with logic.
type Component interface {
	/// Update implements update logic.
	Update(gtx layout.Context) error
	/// Render implements draw widget.
	Render(gtx layout.Context, theme *material.Theme) layout.Dimensions
}
