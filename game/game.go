package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	lasers []*Laser
	meteors  []*Meteor
	meteorSpawnTimer *Timer
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
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}


