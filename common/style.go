package common

import (
	"github.com/aybabtme/rgbterm"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

const (
	Title = iota
	Normal
	Alert
)

var Light = []Color{
	{30, 144, 255},
	{89, 90, 78},
	{64, 56, 54},
}

var Dark = []Color{
	{247, 68, 97},
	{173, 195, 192},
	{226, 211, 172},
}

func ColorIt(content string, position int, style ...string) string {
	s := "dark"
	if len(style) >= 1 {
		s = style[0]
	}

	colorIdx := Normal
	switch position {
	case Title:
		colorIdx = Title
	case Normal:
		colorIdx = Normal
	case Alert:
		colorIdx = Alert
	}

	return color(content, s, colorIdx)
}

func color(content, style string, rgbIndex int) string {
	switch style {
	case "light":
		return rgbterm.FgString(content, Light[rgbIndex].R, Light[rgbIndex].G, Light[rgbIndex].B)
	case "dark":
		fallthrough
	default:
		return rgbterm.FgString(content, Dark[rgbIndex].R, Dark[rgbIndex].G, Dark[rgbIndex].B)
	}
}
