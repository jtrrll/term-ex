package view

import (
	"term-ex/model"
	"term-ex/model/world"

	"github.com/gdamore/tcell"
)

// A view that renders the model as colored ASCII text
type TextView struct {
	screen     tcell.Screen // the screen to render to
	fogEnabled bool         // whether to obscure tiles with fog-of-war
	fogRadius  uint8        // the radius of the fog-of-war that obscures unexplored terrain
}

// Creates a new view that renders the model as colored text
func NewTextView(screen tcell.Screen, fogEnabled bool, fogRadius uint8) TextView {
	return TextView{
		screen:     screen,
		fogEnabled: fogEnabled,
		fogRadius:  fogRadius,
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
			_, ok := model.GetTile(world.Position{X: int64(x + xOffset), Y: int64(y + yOffset)})
			//TODO: Fog-of-war, tile set, height
			if x == width/2 && y == height/2 {
				v.screen.SetContent(x, y, 'c', []rune{}, tcell.StyleDefault.Foreground(tcell.ColorMediumVioletRed))
				continue
			}
			if ok {
				v.screen.SetContent(x, y, 't', []rune{}, tcell.StyleDefault)
			}
		}
	}

	v.screen.Show()
}

// Renders a splash screen
func (v *TextView) RenderSplash() {
	//TODO: Implement
}

// Renders an exit screen
func (v *TextView) RenderExit() {
	//TODO: Implement
}
