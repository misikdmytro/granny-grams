package processor

import (
	"context"
	"image/png"
	"os"

	"github.com/fogleman/gg"
)

type ImageProcessor interface {
	AddTextToImage(ctx context.Context, originalpath, resultpath, text string) error
}

type imageProcessor struct {
}

// AddTextToImage implements ImageProcessor.
func (i *imageProcessor) AddTextToImage(ctx context.Context, originalpath string, resultpath string, text string) error {
	f, err := os.Open(originalpath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return err
	}

	dc := gg.NewContextForImage(img)
	if err = dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 48); err != nil {
		return err
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(text, 512, 50, 0.5, 0.5)

	err = dc.SavePNG(resultpath)
	if err != nil {
		return err
	}

	return nil
}

func NewImageProcessor() ImageProcessor {
	return &imageProcessor{}
}
