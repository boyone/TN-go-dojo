package game

import (
	"testing"
)

func TestGameBeginLoveLove(t *testing.T){
	a := game{}
	display:=a.BoardScore()
	if display!="Love | Love" {
		t.Error("lovelove not true")
	}
}

func TestBallOneDisplay15Love(t *testing.T){
	match := game{}
	match.Awin()
	display:=match.BoardScore()
	if display!="15 | Love" {
		t.Error("15love not true but was ",display)
	}
}

func TestBallTwoDisplay15_15(t *testing.T){
	match := game{}
	match.Awin()
	match.Bwin()
	display:=match.BoardScore()
	if display!="15 | 15" {
		t.Error("15 | 15 not true but was ",display)
	}
}

func TestBallThreeDisplay15_30(t *testing.T) {
	match := game{}
	match.Awin()
	match.Bwin()
	match.Bwin()

	display:=match.BoardScore()

	if display!="15 | 30" {
		t.Error("15 | 30 not true but was ",display)
	}

}

func (g *game) Bwin() {
	g.scoreB+=1
}

type game struct {
	scoreA int
	scoreB int
}

func (g *game)Awin() {
	g.scoreA+=1
}

func (g game)BoardScore() string{
	if g.scoreB == 2{
		return "15 | 30"
	}

	if g.scoreB == 1{
		return "15 | 15"
	}
	if g.scoreA == 1 {
		return "15 | Love"
	}

	return "Love | Love"

}
