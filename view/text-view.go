package view

import (
	"term-ex/model"
	"term-ex/position"

	"github.com/gdamore/tcell"
)

// A view that renders the model as colored ASCII text
type TextView struct {
	screen     tcell.Screen // the screen to render to
	fogEnabled bool         // whether to obscure tiles with fog-of-war
	fogRadius  uint8        // the radius of the fog-of-war that obscures unexplored terrain
	noIntro    bool         // whether to render an intro animation
	noOutro    bool         // whether to render an outro animation
}

// Creates a new view that renders the model as colored text
func NewTextView(screen tcell.Screen, fogEnabled bool, fogRadius uint8, noIntro bool, noOutro bool) *TextView {
	return &TextView{
		screen:     screen,
		fogEnabled: fogEnabled,
		fogRadius:  fogRadius,
		noIntro:    noIntro,
		noOutro:    noOutro,
	}
}

// Renders the provided state
func (v *TextView) RenderModel(model model.Model) {
	width, height := v.screen.Size()
	modelCenter := model.GetPosition()
	// View coordinate + offset = model coordinate
	xOffset := int(modelCenter.X) - width/2
	yOffset := int(modelCenter.Y) - height/2
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, ok := model.GetTile(position.Position{X: int64(x + xOffset), Y: int64(y + yOffset)})
			//TODO: Fog-of-war, tile set, height
			if x == width/2 && y == height/2 {
				v.screen.SetContent(x, y, 'c', []rune{}, tcell.StyleDefault.Foreground(tcell.ColorMediumVioletRed))
				continue
			}
			if ok {
				v.screen.SetContent(x, y, 't', []rune{}, tcell.StyleDefault)
			} else {
				v.screen.SetContent(x, y, ' ', []rune{}, tcell.StyleDefault)
			}
		}
	}
	v.screen.Show()
}

// Renders an intro screen
func (v *TextView) RenderIntro() {
	if v.noIntro {
		return
	}
	//TODO: Implement
	v.screen.Fill('i', tcell.StyleDefault)
}

// Renders an outro screen
func (v *TextView) RenderOutro() {
	//TODO: Implement
	if v.noOutro {
		return
	}
	v.screen.Fill('o', tcell.StyleDefault)
}

// Shuts down the view
func (v *TextView) ShutDown() {
	v.screen.Fini()
}
