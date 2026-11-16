package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

// GeneratePNG создаёт одноцветное PNG-изображение заданного размера.
func GeneratePNG(width, height int) []byte {
	rgb := make([]byte, 3)
	if _, err := rand.Read(rgb); err != nil {
		panic(fmt.Sprintf("Не удалось сгенерировать цвет изображения: %s", err.Error()))
	}
	fill := color.RGBA{R: rgb[0], G: rgb[1], B: rgb[2], A: 255}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, fill)
		}
	}

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, img); err != nil {
		panic(fmt.Sprintf("Не удалось закодировать PNG: %s", err.Error()))
	}

	return buffer.Bytes()
}
