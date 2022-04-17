package asset

import (
	"encoding/json"
	"errors"
	"image"
	_ "image/png"
	"io/fs"
	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/unitoftime/packer"
)

type Load struct {
	fsys fs.FS
}

// NewLoad returns a new loaded filesystem
func NewLoad(fsys fs.FS) *Load { return &Load{fsys: fsys} }

// Open is fs.Open
func (l *Load) Open(path string) (fs.File, error) {
	return l.fsys.Open(path)
}

func (l *Load) Image(path string) (image.Image, error) {
	f, err := l.fsys.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (l *Load) Sprite(path string) (*pixel.Sprite, error) {
	img, err := l.Image(path)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pic, pic.Bounds()), nil
}

func (l *Load) JSON(path string, data interface{}) error {
	f, err := l.fsys.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	jData, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(jData, &data)
}

func (l *Load) SpriteSheet(path string) (*SpriteSheet, error) {
	serializedSS := packer.SerializedSpritesheet{}
	if err := l.JSON(path, &serializedSS); err != nil {
		return nil, err
	}
	img, err := l.Image(serializedSS.ImageName)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)
	b := pic.Bounds()
	lookup := make(map[string]*pixel.Sprite)
	for name, sf := range serializedSS.Frames {
		rect := pixel.R(
			sf.Frame.X,
			b.H()-sf.Frame.Y,
			sf.Frame.X+sf.Frame.W,
			b.H()-(sf.Frame.Y+sf.Frame.H)).Norm()

		lookup[name] = pixel.NewSprite(pic, rect)
	}
	return NewSpriteSheet(pic, lookup), nil
}

// SpriteSheet contains a picture of all the packed sprites and a lookup map of
// all the sprites in the picture.
type SpriteSheet struct {
	pic    pixel.Picture
	lookup map[string]*pixel.Sprite
}

// NewSpriteSheet takes in a picture and lookup map and gives back a
// SpriteSheet.
func NewSpriteSheet(p pixel.Picture, lookup map[string]*pixel.Sprite) *SpriteSheet {
	return &SpriteSheet{pic: p, lookup: lookup}
}

// Get gets a sprite by it's name if it is loaded or returns an error
func (s *SpriteSheet) Get(name string) (*pixel.Sprite, error) {
	spr, found := s.lookup[name]
	if !found {
		return nil, errors.New("Invalid sprite name")
	}
	return spr, nil
}

// Picture gets the SpriteSheet's initialized picture
func (s *SpriteSheet) Picture() pixel.Picture { return s.pic }
