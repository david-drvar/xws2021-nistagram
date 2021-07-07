package images

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type Encoder func(io.Writer, image.Image) error

const (
	EncoderSaving        = "saving"
	EncoderLoading       = "loading"
	EncoderMediumQuality = 70
	EncoderMaxQuality    = 100
)

// GetJPEGEncoder modes are saving or loading.
// Saving will reduce quality but happen only once
// whereas loading with always load it with max quality that it has.
func GetJPEGEncoder(mode string) Encoder {
	switch mode {
	case EncoderSaving:
		return jpegEncoder(EncoderMediumQuality)
	case EncoderLoading:
		return jpegEncoder(EncoderMaxQuality)
	default:
		return jpegEncoder(EncoderMaxQuality)
	}
}

func GetPNGEncoder() Encoder {
	return pngEncoder()
}

// JPEGEncoder returns an encoder to JPEG given the argument 'quality'
func jpegEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}

// PNGEncoder returns an encoder to PNG
func pngEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}
