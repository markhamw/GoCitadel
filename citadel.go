package main

type citentity interface {
	hitroll(int) bool
	isalive() bool
	GetName() string
	GetHealth() string
	getdamageroll() int
}

type playerchar struct {
	Name          string
	battlecry     string //add this
	Health        int
	CurrHealthMax int
	Area          string //chapter/floor
	location      uint8  //primary key for maplocs
	Atk           int    //scales with lvl. Starts at 2. multiplier of Rend method func
	Weapon        string
	Weapondmg     int
	Level         int
	Str           int
	Dex           int
	Con           int
	Int           int
	totxp         int
	kills         int
	inv           []*item
}

type item struct {
	name     string
	descrip  string
	caneat   bool
	candrink bool
	canequip bool
	cost     int
}

type enemy struct {
	Name          string
	state         string
	Typeofen      string
	Typeofatk     string
	Atk           int
	Battlecry     string
	Deathcry      string
	Health        int
	CurrHealthMax int
	worthxp       int
	img           []string
	items         []*item
}

type maploc struct {
	roomid      uint8 //used with playerchar mapping
	grid        string
	area        string
	descrip     string
	lookdescrip []string
	title       string
	mobs        []*enemy
	items       []*item
	exits       []int // value of -1 = no exit , predefined list of end point rooms for each room, adjacency map. eg -1,-1,3,1,-1,-1
}

var p1 = new(playerchar)

func main() {
	promptofshame()
	titleintro()
	storyintro()
	initworld()
	tutintro(p1)
	navigator(p1)

}
