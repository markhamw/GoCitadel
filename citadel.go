package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Playerchar struct {
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
}

type Enemy struct {
	Name      string
	Typeofen  string
	Typeofatk string
	Atk       int
	Battlecry string
	Deathcry  string
	Health    int
	worthxp   int
}

type maploc struct {
	roomid  uint8 //foreign key for playerchar mapping
	grid    string
	area    string
	descrip string
	title   string
	mobs    []*Enemy
	exits   []int // value of -1 = cant go , predefined list of end point rooms for each room, adjacency map. eg -1,-1,3,1,-1,-1
}

func playerprompt(player *Playerchar) string {
	fmt.Print("\n")
	promptval := strings.ToUpper(player.Name + "═")
	prompthp := strconv.Itoa(int(player.Health))
	return promptval + string(prompthp) + "HP═>"
}

func generatemaplocdata(loc *maploc) {
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Print(NAVAREA)
	fmt.Printf("Area: ╠═%v═╣\n\n", loc.title)
	fmt.Print(NAVLOC)
	color.Set(color.FgRed, color.Bold, color.BgBlack)
	fmt.Printf("Detailed Map:\n%v\n", loc.grid)
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Println(compass)
	color.Set(color.FgCyan, color.Bold, color.BgBlack)
	fmt.Print(NAVDESCRIP)
	fmt.Println("~~░░▒ DESCRIPTION ▒░░~~")
	fmt.Println(loc.descrip)
	fmt.Print(NAVTRAVEL)
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Println("Use n,s,e,w   u(travel upwards),d(descend)  b=battle i=inventory q=stats l=look ")

}

func checkbattle(player *Playerchar, mob *Enemy) {
	fmt.Print(ENEMIESLISTED)
	color.Set(color.FgRed, color.Bold, color.BgBlack) //movecursor

	if mob.Health <= 0 {
		color.Set(color.FgWhite, color.BgBlack)
		fmt.Printf("\nThe corpse of %v the %v is here.\n", mob.Name, mob.Typeofen)
	} else if mob.Health > 0 {
		fmt.Printf("\n%v the %v is here.", mob.Name, mob.Typeofen)
	}

	time.Sleep(time.Millisecond * 200)
}

func playercantgo(player *Playerchar, c rune) {
	fmt.Print(PLAYERPROMPT)
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	switch c {
	case 'n':
		fmt.Printf("%v cannot move north any farther.", strings.ToTitle(player.Name))
	case 's':
		fmt.Printf("%v is entirely unable to travel further south.", strings.ToTitle(player.Name))
	case 'e':
		fmt.Printf("%v can't go further east.", strings.ToTitle(player.Name))
	case 'w':
		fmt.Printf("%v can't head further west.", strings.ToTitle(player.Name))
	case 'd':
		fmt.Printf("You want to burrow into the ground? Figure out something else, %v.", strings.ToTitle(player.Name))
	case 'u':
		fmt.Printf("You wot mate? Can't go up here, %v.", strings.ToTitle(player.Name))
	}

	time.Sleep(time.Millisecond * 650)
}
func checkmove(player *Playerchar, exits []int) {
	color.Set(color.FgHiMagenta, color.Bold, color.BgBlack)
	fmt.Printf(playerprompt(player))
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	switch OKP() {
	case 'n':
		switch exits[0] {
		case -1:
			playercantgo(player, 'n')
			break
		default:
			fmt.Printf("%v heads north", strings.ToTitle((player.Name)))
			player.location = uint8(exits[0])

		}

	case 's':
		switch exits[1] {
		case -1:
			playercantgo(player, 's')
			break
		default:
			fmt.Printf("%v heads south", strings.ToTitle((player.Name)))
			player.location = uint8(exits[1])
		}

	case 'e':

		switch exits[2] {
		case -1:
			playercantgo(player, 'e')
			break
		default:
			fmt.Printf("%v heads east", strings.ToTitle((player.Name)))
			player.location = uint8(exits[2])
		}

	case 'w':
		switch exits[3] {
		case -1:
			playercantgo(player, 'w')
			break
		default:
			fmt.Printf("%v heads west", strings.ToTitle((player.Name)))
			player.location = uint8(exits[3])
		}
	case 'u':
		switch exits[4] {
		case -1:
			playercantgo(player, 'u')
			break
		default:
			fmt.Printf("%v goes up", strings.ToTitle((player.Name)))
			player.location = uint8(exits[4])
		}
	case 'd':
		switch exits[5] {
		case -1:
			playercantgo(player, 'd')
			break
		default:
			fmt.Printf("%v heads downward", strings.ToTitle((player.Name)))
			player.location = uint8(exits[5])
		}
	case 'b':
		player.battleinit()

	}
	time.Sleep(time.Millisecond * 400)
	navigator(player)
}

func navigator(player *Playerchar) {
	clear()
	drawplayerbar(player)
	player.makehealthbar()
	color.Set(color.FgYellow, color.Bold, color.BgBlack)

	switch player.location {
	case 1:
		generatemaplocdata(cellar1)
		checkbattle(player, rat1)
		checkmove(player, cellar1.exits) //passes to checkmove: pointer to Playerchar struct and slice of string from maploc
	case 2:
		generatemaplocdata(cellar2)
		checkbattle(player, rat2)
		checkmove(player, cellar2.exits)
	case 3:
		generatemaplocdata(cellar3)
		checkbattle(player, rat3)
		checkmove(player, cellar3.exits)
	case 4:
		generatemaplocdata(cellar4)
		checkbattle(player, bossrat)
		checkmove(player, cellar4.exits)
	}

	time.Sleep(time.Millisecond * 2000)

}

func main() {

	titleintro()
	storyintro()
	//              //temp - remove after titleintro and storyintro are back in
	P1 := new(Playerchar)          //creates player struct
	initworld()                    //sets up Enemies and maplocs
	initplayer(P1)                 //set initial player values
	P1.SetName(initplayername(P1)) //set name and confirm choice
	P1.SetHealth()
	tutintro(P1)
	OKP() //PAUSE

	//todo: create functions for main loop and combat loops and drawing each room
	navigator(P1) //??? calls drawplayerbar and generates room data

}
