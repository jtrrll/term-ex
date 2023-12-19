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
}

// Creates a new view that renders the model as colored text
func NewTextView(screen tcell.Screen, fogEnabled bool, fogRadius uint8) *TextView {
	return &TextView{
		screen:     screen,
		fogEnabled: fogEnabled,
		fogRadius:  fogRadius,
	}
}

// Renders an intro screen
func (v *TextView) RenderIntro() {
	//TODO: Implement
	v.screen.Fill(' ', tcell.StyleDefault.Background(tcell.ColorTan))
	// width, height := v.screen.Size()
	// if width < 80 || height < 24 {
	// 	// TODO: Small mode
	// } else {
	// 	// TODO: Regular mode
	// }
	v.screen.Show()
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
			//TODO: Fog-of-war, tile set, height
			ti, _ := model.GetTile(position.Position{X: int64(x + xOffset), Y: int64(y + yOffset)})
			var char rune
			var style tcell.Style

			styledRune := styledRunes[ti]
			char = styledRune.Rune
			style = styledRune.Style
			if x == width/2 && y == height/2 {
				char = 'c'
				style = style.Foreground(tcell.ColorMediumVioletRed).Bold(true)
			}
			v.screen.SetContent(x, y, char, []rune{}, style)

		}
	}
	v.screen.Show()
}

// Renders an outro screen
func (v *TextView) RenderOutro() {
	//TODO: Implement
	v.screen.Fill('o', tcell.StyleDefault.Background(tcell.ColorDarkRed))
	v.screen.Show()
}

// Shuts down the view. This must be called before the application exits
func (v *TextView) ShutDown() {
	v.screen.Fini()
}
