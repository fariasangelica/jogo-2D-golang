package game

import (
	"fmt"
	"image/color"
	"jogo-2d-golang/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player *Player
	lasers []*Laser
	meteors  []*Meteor
	meteorSpawnTimer *Timer
	score    int
	
}

func NewGame() *Game {
	g := &Game{
		
		meteorSpawnTimer: NewTimer(24),
		
		
	}
	player := NewPlayer(g)
	g.player = player
	return g
	}

// métodos da lib para o funcionamento update, draw e layout

// Responsável por atualizar a lógica do jogo
func (g *Game) Update() error {
	g.player.Update()

	for _, l := range(g.lasers) {
		l.Upadate()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset() 

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range(g.meteors){
		m.Upadate()
	}

	for _, m := range g.meteors {
		if(m.Collider().Intersects(g.player.Collider())){
			fmt.Println("Você perdeu!")
			g.Reset()

		}
	}

	for i, m := range(g.meteors) {

		for j, l := range(g.lasers) {

			if m.Collider().Intersects(l.Collider()){
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1

			}
		
		}
	}

	return nil
}


// Tamanho da tela
// Desenhar objetos na tela

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range(g.lasers) {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos : %d", g.score), assets.FontUi, 20, 100, color.White)
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers =nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}


