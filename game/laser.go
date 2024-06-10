package game

import (
	"jogo-2d-golang/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()

	halfW := float64(bounds.Dx()) / 2 //Metade da largura da imagem do laser
	halfH := float64(bounds.Dy()) /2 // Metade da altura

	position.X -= halfW
	position.Y -= halfH


	return &Laser{
		image: image,
		position: position,
	}
}

func (l *Laser) Upadate() {
	speed := 7.0

	l.position.Y += -speed

}
	
func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
    
	//Posição x e y que a imagem será desenhada a tela
	op.GeoM.Translate(l.position.X, l.position.Y)
 
	//Desenha imagem na tela
	screen.DrawImage(l.image, op)
}