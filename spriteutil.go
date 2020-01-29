package spriteutil

import (
	"fmt"
	"image/gif"
	_ "image/png"
	"io"

	"github.com/faiface/pixel"
)

// LoadGif returns an AnimatedSprite
func LoadGif(file io.Reader) (*AnimatedSprite, error) {
	g, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}
	if len(g.Image) == 0 {
		return nil, fmt.Errorf("no frames in gif")
	}
	frames := make([]pixel.Picture, len(g.Image))
	sprites := make([]*pixel.Sprite, len(g.Image))
	for i := range g.Image {
		frames[i] = pixel.PictureDataFromImage(g.Image[i])
		sprites[i] = pixel.NewSprite(frames[i], pixel.R(0, 0, float64(g.Config.Width), float64(g.Config.Height)))
	}
	return &AnimatedSprite{frames: frames, delays: g.Delay, sprites: sprites}, nil
}

// AnimatedSprite holds gif frames and manages timing. don't forget to call `a.Update(dt)`
type AnimatedSprite struct {
	anims   map[string][]pixel.Rect
	frames  []pixel.Picture
	sprites []*pixel.Sprite
	delays  []int
	dt      float64
	frame   int
}

// Draw a gif frame onto a target
func (a *AnimatedSprite) Draw(t pixel.Target, m pixel.Matrix) {
	a.sprites[a.frame].Draw(t, m)
}

// Update gif frame based on time that has passed since last call to Update
// Get dt in your mainloop like this: `dt = time.Since(last).Seconds()`
func (a *AnimatedSprite) Update(dt float64) {
	a.dt += dt
	if 100*a.dt > float64(a.delays[a.frame]) {
		a.frame = (a.frame + 1) % len(a.frames)
		a.dt = 0
	}

}
