package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// métodos da lib para o funcionamento update, draw e layout

// Responsável por atualizar a lógica do jogo
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Desenhar objetos na tela
func (g *Game) Draw(screen *ebiten.Image) {
	
	g.player.Draw(screen)

}

// Tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (f * Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
