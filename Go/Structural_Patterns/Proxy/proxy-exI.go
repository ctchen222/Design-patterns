package proxy

import (
	"fmt"
	"sync"
)

type Image interface {
	Display()
}

type RealImage struct {
	ImageFile []string
}

func (r *RealImage) Display() {
	for _, image := range r.ImageFile {
		fmt.Printf("Loading image from disk %s\n", image)
	}
}

func NewRealImage(imageFile []string) *RealImage {
	return &RealImage{
		ImageFile: imageFile,
	}
}

type ProxyImage struct {
	ImageFile []string
	realImage *RealImage
	once      sync.Once
}

func (p *ProxyImage) Display() {
	p.once.Do(func() {
		p.realImage = NewRealImage(p.ImageFile)
	})
	p.realImage.Display()
}
