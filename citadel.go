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
}

type enemy struct {
	Name          string
	Typeofen      string
	Typeofatk     string
	Atk           int
	Battlecry     string
	Deathcry      string
	Health        int
	CurrHealthMax int
	worthxp       int
}

type maploc struct {
	roomid  uint8 //foreign key for playerchar mapping
	grid    string
	area    string
	descrip string
	title   string
	mobs    []*enemy
	exits   []int // value of -1 = cant go , predefined list of end point rooms for each room, adjacency map. eg -1,-1,3,1,-1,-1
}

var P1 playerchar

func main() {
	titleintro()
	storyintro()
	P1 := new(playerchar)          //creates player struct
	initworld()                    //sets up Enemies and maplocs
	initplayer(P1)                 //set initial player values
	P1.SetName(initplayername(P1)) //set name and confirm choice
	P1.SetHealth()
	getcharpause() //PAUSE
	tutintro(P1)
	navigator(P1) //??? calls drawplayerbar and generates room data

	//xp doesnt tally
	//level up text stays on screen
	//keep play-testing
}
