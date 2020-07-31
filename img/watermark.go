package img

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
)

func WriteText(im image.Image, text string, fontobj *truetype.Font) error {
	c := freetype.NewContext()
	c.SetDPI(float64(im.Bounds().Dy() / len([]rune(text))))
	c.SetFont(fontobj)
	c.SetFontSize(100)
	c.SetClip(im.Bounds())
	c.SetDst(im.(*image.RGBA))
	c.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))
	count := 10
	if len([]rune(text)) > 10 {
		count = len([]rune(text))
	}
	pt := freetype.Pt(im.Bounds().Dx()/count, im.Bounds().Dy()/2)
	_, err := c.DrawString(text, pt)
	return err
}

// 图片写入水印
// @params img: 水印
// @params im: 需要写入水印的原图片
func WriteImageWatermark(img image.Image, watermark image.Image, dest io.Writer, x, y int) error {
	offset := image.Pt(x, y)
	b := img.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)
	return jpeg.Encode(dest, m, &jpeg.Options{100})
}
