package main

import (
	"image"
	"image/color"
	"strings"
)

func getASCIIColor(charset string, character rune) color.Color {
	index := strings.Index(charset, string(character))

	ratio := float32(index) / float32(len(charset))
	colorValue := uint8(ratio * 255)

	col := color.RGBA{
		R: colorValue,
		G: colorValue,
		B: colorValue,
		A: uint8(255),
	}

	return col
}

func processASCII(charset string, scale uint, input string) image.Image {
	i, j := uint(0), uint(0)

	lines := uint(0)
	for _, letter := range input {
		if letter == '\n' {
			lines++
		}
	}

	baseImg := image.NewRGBA(image.Rectangle{
		Min: image.Pt(0, 0),
		Max: image.Pt(len(input)/int(lines)*int(scale), int(lines)*int(scale)), // len(input) / lines is the number of columns
	})

	for _, letter := range input {
		if letter == '\n' {
			j++
			i = 0
			continue
		}
		col := getASCIIColor(charset, letter)
		for x := range scale {
			for y := range scale {
				baseImg.Set(int(i*scale+x), int(j*scale+y), col)
			}
		}

		i++
	}
	return baseImg
}
