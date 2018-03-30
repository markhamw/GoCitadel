package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

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

func conred() {
	color.Set(color.FgRed, color.Bold, color.BgBlack)
}

func entprompt(thing citentity) string {
	promptval := strings.ToUpper("░░░▒▒▓ " + thing.GetName() + " ▓▒▒░░░ ")
	return promptval
}
func printcompass() {
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print(COMPASSDRAW)
	fmt.Print(compass)
}
func (l *maploc) printmap() {
	fmt.Print(NAVLOC)
	conred()
	fmt.Printf("Detailed Map:\n%v\n", l.grid)
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print(NAVAREA)
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Printf("╠═%v═╣\n\n", l.title)

	color.Set(color.FgCyan, color.Bold, color.BgBlack)
	fmt.Print(NAVDESCRIP)
	fmt.Printf(" ¬│░░░░▒▒▒▒▒▒▓▓▓▓▓▌DESCRIPTION▐▓▓▓▓▓▒▒▒▒▒▒░░░░│⌐ \n")
	l.printdescrip()
	printcompass()
}

func (l maploc) printdescrip() {
	var counter int
	for _, x := range l.descrip {
		counter++
		fmt.Print(string(x))
		if counter == 53 { //Number of rows for column
			fmt.Print("\n")
			counter = 0
		}
	}

}

func generatemaplocdata(loc *maploc) {

	loc.printmap()
	fmt.Print(NAVTRAVEL)
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Printf("Use n,s,e,w   u(up),d(down)  b=battle i=view iventory\nq=View %v Stats l=look ", p1.Name)

}

func checkbattle(player *playerchar, mob *enemy) {
	fmt.Print(ENEMIESLISTED)
	conred()

	if mob.Health <= 0 {
		color.Set(color.FgWhite, color.BgBlack)
		fmt.Printf("\nThe corpse of %v the %v is here.\n", mob.Name, mob.Typeofen)
	} else if mob.Health > 0 {
		fmt.Printf("\n%v the %v is here.\n", mob.Name, mob.Typeofen)
	}

	time.Sleep(time.Millisecond * 200)
}

func playercantgo(player *playerchar, c rune) {
	entprompt(player)

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
func checkmove(player *playerchar, exits []int) {
	//color.Set(color.FgHiMagenta, color.Bold, color.BgBlack)
	//fmt.Printf(entprompt(player))
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	switch getcharpause() {
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
	case 'q':
		player.printstats()
	}
	time.Sleep(time.Millisecond * 400)
	navigator(player)
}

func navigator(player *playerchar) {
	clear()
	drawplayerbar(player)
	player.makehealthbar()
	color.Set(color.FgYellow, color.Bold, color.BgBlack)

	switch player.location {
	case 1:
		generatemaplocdata(cellar1)
		checkbattle(player, rat1)
		checkmove(player, cellar1.exits) //passes to checkmove: pointer to playerchar struct and slice of string from maploc
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

var p1 playerchar

func main() {

	titleintro()
	storyintro()
	//clear()                        //temp - remove after titleintro and storyintro are back in
	p1 := new(playerchar)          //creates player struct
	initworld()                    //sets up Enemies and maplocs
	initplayer(p1)                 //set initial player values
	p1.SetName(initplayername(p1)) //set name and confirm choice
	p1.SetHealth()
	//tutintro(P1)
	getcharpause() //PAUSE

	//todo: create functions for main loop and combat loops and drawing each room
	navigator(p1) //??? calls drawplayerbar and generates room data

}
