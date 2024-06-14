package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"time"
)

// 字符点阵定义
var charMap = map[rune][]string{
	'0': {
		"   #######   ",
		"  ##     ##  ",
		" ##       ## ",
		"##         ##",
		"##         ##",
		"##         ##",
		"##         ##",
		" ##       ## ",
		"  ##     ##  ",
		"   #######   ",
	},
	'1': {
		"     ##      ",
		"    ###      ",
		"   ####      ",
		"     ##      ",
		"     ##      ",
		"     ##      ",
		"     ##      ",
		"     ##      ",
		"     ##      ",
		"   ######    ",
	},
	'2': {
		"   #######   ",
		"  ##     ##  ",
		"         ##  ",
		"         ##  ",
		"   #######   ",
		"  ##         ",
		" ##          ",
		"##           ",
		"##           ",
		"###########  ",
	},
	'3': {
		"   #######   ",
		"  ##     ##  ",
		"         ##  ",
		"         ##  ",
		"   #######   ",
		"         ##  ",
		"         ##  ",
		"         ##  ",
		"  ##     ##  ",
		"   #######   ",
	},
	'4': {
		"##           ",
		"##     ##    ",
		"##     ##    ",
		"##     ##    ",
		"##     ##    ",
		"###########  ",
		"        ##   ",
		"        ##   ",
		"        ##   ",
		"        ##   ",
	},
	'5': {
		"###########  ",
		"##           ",
		"##           ",
		"##           ",
		"##########   ",
		"         ##  ",
		"         ##  ",
		"         ##  ",
		"##      ##   ",
		" ########    ",
	},
	'6': {
		"   #######   ",
		"  ##     ##  ",
		" ##       ## ",
		"##           ",
		"##########   ",
		"##       ##  ",
		"##       ##  ",
		"##       ##  ",
		" ##     ##   ",
		"   #######   ",
	},
	'7': {
		"###########  ",
		"        ##   ",
		"       ##    ",
		"      ##     ",
		"     ##      ",
		"    ##       ",
		"   ##        ",
		"  ##         ",
		" ##          ",
		"##           ",
	},
	'8': {
		"   #######   ",
		"  ##     ##  ",
		" ##       ## ",
		" ##       ## ",
		"  ##     ##  ",
		"   #######   ",
		"  ##     ##  ",
		" ##       ## ",
		" ##       ## ",
		"  ##     ##  ",
		"   #######   ",
	},
	'9': {
		"   #######   ",
		"  ##     ##  ",
		" ##       ## ",
		" ##       ## ",
		"  ##     ##  ",
		"   ######### ",
		"         ##  ",
		"         ##  ",
		" ##      ##  ",
		"  #######    ",
	},
}

func drawText(img *image.RGBA, x, y int, text string, col color.RGBA) {
	for _, c := range text {
		if bitmap, ok := charMap[c]; ok {
			for row, line := range bitmap {
				for colIdx, pixel := range line {
					if pixel == '#' {
						img.Set(x+colIdx, y+row, col)
					}
				}
			}
		}
		x += len(charMap['0'][0]) + 2 // Move to the next character position
	}
}

func GenerateCaptchaImage(text string) (string, error) {
	const width, height = 180, 40

	// 创建一个新的RGBA图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 设置背景颜色为白色
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: white}, image.Point{}, draw.Src)

	// 绘制干扰线
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		col := color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}
		x1 := rand.Intn(width)
		y1 := rand.Intn(height)
		x2 := rand.Intn(width)
		y2 := rand.Intn(height)
		drawLine(img, x1, y1, x2, y2, col)
	}

	// 设置文本颜色为黑色
	black := color.RGBA{A: 255}

	// 计算文本的起始位置以居中对齐
	charWidth := len(charMap['0'][0])
	textWidth := len(text) * (charWidth + 2)
	x := (width - textWidth) / 2
	y := (height - len(charMap['0'])) / 2

	// 绘制验证码文本
	drawText(img, x, y, text, black)

	// 将图像编码为PNG格式
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	// 将图像转换为Base64编码字符串
	base64Img := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Img, nil
}

// drawLine 在图像上绘制一条线
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, col color.RGBA) {
	dx := x2 - x1
	dy := y2 - y1
	steps := abs(dx)
	if abs(dy) > steps {
		steps = abs(dy)
	}
	xIncrement := float64(dx) / float64(steps)
	yIncrement := float64(dy) / float64(steps)
	x := float64(x1)
	y := float64(y1)
	for i := 0; i <= steps; i++ {
		img.Set(int(x), int(y), col)
		x += xIncrement
		y += yIncrement
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
