package window

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

// Processor implements the window basic mechanism.
type Processor struct {
	theme     *material.Theme
	component Component
}

// NewProcessor creates window processor.
func NewProcessor(component Component) *Processor {
	return &Processor{
		theme:     material.NewTheme(gofont.Collection()),
		component: component,
	}
}

// Run implements window events and render components.
func (p *Processor) Run(window *app.Window) error {
	var ops op.Ops
	for evt := range window.Events() {
		switch e := evt.(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if err := p.component.Update(gtx); err != nil {
				return err
			}
			p.component.Render(gtx, p.theme)

			e.Frame(gtx.Ops)

		default:
			// log.Printf("window event: %#v\n", e)
		}
	}
	return nil
}
