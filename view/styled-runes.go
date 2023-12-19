package view

import (
	"term-ex/tile"

	"github.com/gdamore/tcell"
)

// A Rune with a visual style
type StyledRune struct {
	Rune  rune        // the rune to draw
	Style tcell.Style // the style of the rune
}

var (
	styledRunes = map[tile.Tile]StyledRune{
		tile.Ocean:    {'/', tcell.StyleDefault.Foreground(tcell.ColorBlue)},
		tile.Sand:     {'.', tcell.StyleDefault.Foreground(tcell.ColorTan)},
		tile.Grass:    {'"', tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)},
		tile.Dirt:     {'.', tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown)},
		tile.Path:     {':', tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown)},
		tile.Mountain: {'^', tcell.StyleDefault.Foreground(tcell.ColorGray)},
		tile.Wall:     {'#', tcell.StyleDefault.Foreground(tcell.ColorTan)},
	}
)
