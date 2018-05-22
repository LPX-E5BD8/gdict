package common

import "github.com/aybabtme/rgbterm"

type Color struct {
	R uint8
	G uint8
	B uint8
}

var Light = []Color{
	{23, 50, 7},
	{119, 52, 96},
	{78, 29, 76},
}

var Dark = []Color{
	{247, 68, 97},
	{173, 195, 192},
	{226, 211, 172},
}

func ColorTitle(content string, style ...string) string {
	s := "dark"
	if len(style) > 0 {
		s = style[0]
	}

	return color(content, s, 0)
}

func ColorNormal(content string, style ...string) string {
	s := "dark"
	if len(style) > 0 {
		s = style[0]
	}

	return color(content, s, 1)
}

func ColorAlert(content string, style ...string) string {
	s := "dark"
	if len(style) > 0 {
		s = style[0]
	}

	return color(content, s, 2)
}

func color(content, style string, rgbIndex int) string {
	switch style {
	case "dark":
		return rgbterm.FgString(content, Dark[rgbIndex].R, Dark[rgbIndex].G, Dark[rgbIndex].B)
	case "light":
		return rgbterm.FgString(content, Light[rgbIndex].R, Light[rgbIndex].G, Light[rgbIndex].B)
	default:
		return content
	}
}
