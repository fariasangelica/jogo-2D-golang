package game

import (
	"jogo-2d-golang/assets"

	"github.com/hajimehoshi/ebiten/v2"
)



type Player struct {
	image *ebiten.Image
	position Vector
	game    *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) /2

    position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,

	}

	return &Player{
		image: image,
		game:  game,
		position: position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {

	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	p.laserLoadingTimer.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição x e y que a imagem será desenhada na tela
	op.GeoM.Translate(p.position.X, p.position.Y)
	// Desenha a imagem na tela
	screen.DrawImage(p.image, op)
	
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewReact(p.position.X, 
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dx()))
}