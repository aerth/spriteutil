package spriteutil

import (
	_ "image/png"
	"io"
	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
)

func LoadTTF(rdr io.Reader, size float64, origin pixel.Vec) *text.Text {
	b, err := ioutil.ReadAll(rdr)
	if err != nil {
		panic(err)
	}
	font, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(origin, atlas)
	return txt
}
