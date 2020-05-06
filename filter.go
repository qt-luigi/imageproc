package imageproc

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/disintegration/gift"
)

// DecreaseBrightness decreases a image file brightness.
func DecreaseBrightness(dir string, fns []string) ([]string, error) {
	cnt := len(fns)
	list := make([]string, cnt)
	per := int(100 / cnt)
	for i, v := range fns {
		srcimg, err := decpng(filepath.Join(dir, v))
		if err != nil {
			return nil, err
		}
		p := (i + 1) * per
		fn := fmt.Sprintf("brightness_%d.png", p)
		dstimg := filter(srcimg, gift.Brightness(float32(-p)))
		if err := encpng(filepath.Join(dir, fn), dstimg); err != nil {
			return nil, err
		}
		list[i] = fn
	}
	return list, nil
}

// Rotate180 upsides dows a image file.
func Rotate180(dir string, fns []string) ([]string, error) {
	cnt := len(fns)
	list := make([]string, cnt)
	per := int(100 / cnt)
	for i, v := range fns {
		srcimg, err := decpng(filepath.Join(dir, v))
		if err != nil {
			return nil, err
		}
		p := (i + 1) * per
		fn := fmt.Sprintf("rotate_%d.png", p)
		dstimg := filter(srcimg, gift.Rotate180())
		if err := encpng(filepath.Join(dir, fn), dstimg); err != nil {
			return nil, err
		}
		list[i] = fn
	}
	return list, nil
}

func filter(srcimg image.Image, filter gift.Filter) image.Image {
	g := gift.New(filter)
	dstimg := image.NewNRGBA(g.Bounds(srcimg.Bounds()))
	g.Draw(dstimg, srcimg)
	return dstimg
}

func decpng(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}

func encpng(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
