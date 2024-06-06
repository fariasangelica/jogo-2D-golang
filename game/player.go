package game

import "github.com/hajimehoshi/ebiten/v2"



type Player struct {
	image *ebiten.Image
	position Vector
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) /2

    position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,

	}

	return &Player{
		image: image,
		position: position,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição x e y que a imagem será desenhada na tela
	op.GeoM.Translate(p.position.X, p.position.Y)
	// Desenha a imagem na tela
	screen.DrawImage(p.image, op)
	
}