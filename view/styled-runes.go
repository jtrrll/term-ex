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
		tile.Fog:      {' ', tcell.StyleDefault},
		tile.Ocean:    {'~', tcell.StyleDefault.Foreground(tcell.ColorDodgerBlue)},
		tile.Water:    {'~', tcell.StyleDefault.Foreground(tcell.ColorDeepSkyBlue)},
		tile.Sand:     {'.', tcell.StyleDefault.Foreground(tcell.ColorTan)},
		tile.Dirt:     {'.', tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown)},
		tile.Grass:    {'"', tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)},
		tile.Path:     {':', tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown)},
		tile.Tree:     {'*', tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)},
		tile.Mountain: {'^', tcell.StyleDefault.Foreground(tcell.ColorGray)},
	}
)
