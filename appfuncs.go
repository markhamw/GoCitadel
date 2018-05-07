package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var cellar1 = new(maploc)
var cellar2 = new(maploc)
var cellar3 = new(maploc)
var cellar4 = new(maploc)
var rat1 = new(enemy)
var rat2 = new(enemy)
var rat3 = new(enemy)
var bossrat = new(enemy)
var flask = new(item)
var rdrops = new(item)

func itemmenu(f *playerchar, items []*item) {
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Print("\n\nUse T(take) ░ D(drop) ░ B=Go Back \n\n")
	conwhite()
	sel := getcharpause()

	switch sel {
	case 't', 'T':
		f.getitems(items)

	}

}

func isstanding(character citentity) bool {
	if character.isalive() {
		return true
	}
	return false
}

func playercantgo(player *playerchar, c rune) {

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
		fmt.Printf("Figure out something else, %v. Can't proceed.", strings.ToTitle(player.Name))
	case 'u':
		fmt.Printf("You wot mate? Can't go up here, %v.", strings.ToTitle(player.Name))
	}

	time.Sleep(time.Millisecond * 650)
}

func checkmove(player *playerchar, exits []int) {

	congreen()
	switch getcharpause() {
	case 'n', 'N':
		switch exits[0] {
		case -1:
			playercantgo(player, 'n')
			break
		default:
			fmt.Printf("%v heads north", strings.ToTitle((player.Name)))
			player.location = uint8(exits[0])

		}

	case 's', 'S':
		switch exits[1] {
		case -1:
			playercantgo(player, 's')
			break
		default:
			fmt.Printf("%v heads south", strings.ToTitle((player.Name)))
			player.location = uint8(exits[1])
		}

	case 'e', 'E':

		switch exits[2] {
		case -1:
			playercantgo(player, 'e')
			break
		default:
			fmt.Printf("%v heads east", strings.ToTitle((player.Name)))
			player.location = uint8(exits[2])
		}

	case 'w', 'W':
		switch exits[3] {
		case -1:
			playercantgo(player, 'w')
			break
		default:
			fmt.Printf("%v heads west", strings.ToTitle((player.Name)))
			player.location = uint8(exits[3])
		}
	case 'u', 'U':
		switch exits[4] {
		case -1:
			playercantgo(player, 'u')
			break
		default:
			fmt.Printf("%v goes up", strings.ToTitle((player.Name)))
			player.location = uint8(exits[4])
		}
	case 'd', 'D':
		switch exits[5] {
		case -1:
			playercantgo(player, 'd')
			break
		default:
			fmt.Printf("%v heads downward", strings.ToTitle((player.Name)))
			player.location = uint8(exits[5])
		}
	case 'b', 'B':
		player.battleinit()
	case 'q', 'Q':
		player.printstats()
	case 'l', 'L':
		player.lookroom(player.location)
	case 'i', 'I':
		player.listinv(player.inv)
	}
	time.Sleep(time.Millisecond * 400)
	navigator(player)
}

func navigator(player *playerchar) {
	clear()
	drawplayerbar(player) //make this a method
	player.makehealthbar()

	//find another way to get player location and draw screens
	//switch case can be improved on

	//generatemaplocdata(player.getroom())
	switch player.location {
	case 1:

		generatemaplocdata(player.getroom()) //prints maploc grid and description, NSEW move options

		checkbattle(player, rat1)        //prints enemies in room
		checkmove(player, cellar1.exits) //passes to checkmove: pointer to playerchar struct and slice of string from maploc
	case 2:
		generatemaplocdata(cellar2, player)
		checkbattle(player, rat2)
		checkmove(player, cellar2.exits)
	case 3:
		generatemaplocdata(cellar3, player)
		checkbattle(player, rat3)
		checkmove(player, cellar3.exits)
	case 4:
		generatemaplocdata(cellar4, player)
		checkbattle(player, bossrat)
		checkmove(player, cellar4.exits)
	}

	time.Sleep(time.Millisecond * 2000)

}

func battreport(p *playerchar, en *enemy) {
	conwhite()
	fmt.Printf("\n%v engages in combat with %v\n\n", p.Name, en.Name)

	conred()
	enemycharges := []string{
		0: "charges at you screaming",
		1: "runs toward you aggressively screeching",
		2: "approaches at impressive speed saying",
		3: "engages you in battle while chanting",
		4: "is pissed and runs at you belting out",
		5: "trundles at you militarily, droning",
	}

	fmt.Printf("%v %v \"%v\"\n\n", en.Name, enemycharges[rand.Intn(6)], en.Battlecry)
	for en.Health > 0 {

		p.isalive()
		if p.hitroll(10) {
			p.isalive()
			p.brpds(en)
			en.isalive()
		} else {
			p.isalive()
			p.brmiss(en)
		}

		if en.hitroll(7) { //difficulty modifier?? //increase to increase liklihood of hitting
			en.isalive()
			en.enemybrpds(p)
			p.isalive()
		} else {
			en.isalive()
			en.enemybrpdsmiss(p)
		}
	}

}

func checkbattle(player *playerchar, mob *enemy) {
	fmt.Print(ENEMIESLISTED) //mopve cursor
	conred()

	if mob.Health <= 0 {
		color.Set(color.FgWhite, color.BgBlack)
		fmt.Printf("\nThe corpse of %v the %v is here.\n", mob.Name, mob.Typeofen)
	} else if mob.Health > 0 {

		fmt.Printf("\n%v the %v is here.\n", mob.Name, mob.Typeofen)
		if mob == bossrat {
			fmt.Printf("His vileness's matted coat is impressive and dull fangs\nare bright orange from munching Elder Wax\n")
		}

	}

	time.Sleep(time.Millisecond * 100)
}

func generatemaplocdata(loc *maploc) {
	loc.printmap(p1.getroom())
	fmt.Print(NAVTRAVEL)

	conyellow()
	fmt.Print("Use N,S,E,W ░ U(up),D(down) ░ B=battle ░\nI=items ░ Q=stats ░ L=look ░")
	p1.checklvl()
}

var lvl1 = true
var lvl2 = true

func playerdeath() {
	fmt.Println(ZEROHOME)
	clear()
	fmt.Println("You died alone. Feels Bad Man.")
	fmt.Printf("Remaining HP: %v ", p1.Health)
	os.Exit(0)
}

func initplayername(player *playerchar) string {
	var playername string
	name := func() rune {
		fmt.Print(HOME)
		fmt.Println(DELETETOENDOFLINE)
		fmt.Println(DELETETOENDOFLINE)
		fmt.Print(HOME)
		fmt.Print("Remember your first name (type it, press enter):")
		fmt.Print(SAVECURSOR)
		fmt.Scan(&playername)
		fmt.Print(RETURNTOSAVEDCURSOR)
		fmt.Printf("\nUse %v? (PRESS ANY KEY TO CONTINUE, or n to try again)", playername)
		answer, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		return answer
	}

	answerstring := string(name())

	switch answerstring {

	case "n", "N":
		name()
	default:
		return playername
	}

	return playername
}

func tutintro(player *playerchar) {
	defer getcharpause()
	clear()
	drawplayerbar(player)
	congreen()
	fmt.Println(castle2)
	conwhite()
	readblock(castletext)

}

func initplayer(player *playerchar) { //set initial player values
	player.Weapon, player.Weapondmg = "Flail", 2
	player.Str, player.Dex, player.Int, player.Con, player.Level, player.Atk = 10, 10, 10, 10, 1, 2
	player.Area = "Castle Iota"
	player.location = 2
	player.SetHealth()
	player.inv = append(player.inv, flask)

}

var compass = []string{0: "N", 1: "W-╬-E", 2: "S"}

func initworld() {
	p1.SetName(initplayername(p1))
	p1.SetHealth()
	p1.Weapon, p1.Weapondmg = "Flail", 2
	p1.Str, p1.Dex, p1.Int, p1.Con, p1.Level, p1.Atk = 10, 10, 10, 10, 1, 2
	p1.Area = "Castle Iota"
	p1.location = 2
	p1.SetHealth()
	p1.inv = append(p1.inv, flask)
	/* type item struct {
		name     string
		descrip  string
		caneat   bool
		candrink bool
		canequip bool
		cost     int
	} */

	flask.name = "Flask"
	flask.descrip = "A tin flask from which to drink. Meant for spirits, liquor or alchemical solutions."
	flask.caneat = false
	flask.candrink = true
	flask.canequip = false
	flask.cost = 0

	rdrops.name = "Excrement of Rattus Norvegicus"
	rdrops.descrip = "It's rat poo. Ehck. DELICIOUS to rats. Hilarious to adventurers. You shouldn't be carrying it."

	cellar1.area = "Cellar"
	cellar2.area = "Cellar"
	cellar3.area = "Cellar"
	cellar4.area = "Cellar"

	cellar1.roomid = 1
	cellar2.roomid = 2
	cellar3.roomid = 3
	cellar4.roomid = 4

	cellar1.grid = "[X]--[ ]--[ ]"

	cellar2.grid = "[ ]--[X]--[ ]"

	cellar3.grid = "[ ]--[ ]--[X]"

	cellar4.grid = "[ ]--[d]--[ ]"

	cellar1.exits = []int{-1, -1, 2, -1, -1, -1}
	cellar2.exits = []int{-1, -1, 3, 1, -1, 4}
	cellar3.exits = []int{-1, -1, -1, 2, -1, -1}
	cellar4.exits = []int{-1, -1, -1, -1, 2, -1}

	cellar1.title = `Western Cellar`
	cellar2.title = ` Main Cellar  `
	cellar3.title = `Eastern Cellar`
	cellar4.title = ` Sub-Basement `
	cellar1.descrip = "You're standing in the western part of the cellar\nbasement. Boxes and crates are stacked in the area.\nRats have gnawed the boxes to bits. You're able\nto move eastward to the central basement. \n"
	cellar2.descrip = "This large area of the cellar is for staging and\nstorage of supplies. Water aqueducts crowd the ceiling\nand air ducts are leading in all directions. Darkness\nclings to everything and you can hear the furnace.\nYou can go East and go deeper into the basement or\nWest towards the storage area of the basement. A\nseparate path leads down into the sub-basement.\n"
	cellar3.descrip = "The eastern area of the basement. This part of the\nbasement should be condemned.  It has been\ndecaying for decades. You can hear the furnace to\nthe west. You can go west back towards the central\narea of the basement. It smells like a Rattus nest.\n"
	cellar4.descrip = "The sub-basement is littered with rat corpses and\npiles of rat treasures and rotting detritus. The\nfurnance is loud from above and keeps the rats warm\nin the winter. A ladder leads back up to the Cellar.\n"

	cellar1.lookdescrip = []string{
		0: "Vermin have clearly attempted to re-purpose\n",
		1: "this area as their own. Useless containers\n",
		2: "line the walls and a few have bite marks.\n",
		3: "A few important useless containers are encased\n",
		4: "in Elder Wax for safe storage.\n",
	}

	cellar2.lookdescrip = []string{
		0: "Fire light is escaping from furnace vents.\n",
		1: "The room is large and the ceiling is lost\n",
		2: "to darkness. The sub-basement is open and\n",
		3: "is usually kept closed. You recall that the\n",
		4: "sub-basement might be dangerous. Rat nobility.\n",
	}

	cellar3.lookdescrip = []string{
		0: "Cellar walls are crumbling and masons have\n",
		1: "tried to repair the failing infrastructure.\n",
		2: "It appears vermin have piled bits of bone and\n",
		3: "food in this part of the cellar.\n",
	}

	cellar4.lookdescrip = []string{
		0: "The sub-basement is a stately room for rulers\n",
		1: "of rats and vermin.  A RUS-sized hole leads to\n",
		2: "the other pathways that are too small to fit\n",
		3: "through. Bits of hair and bone and rotting rat\n",
		4: "treasures are present.\n",
	}

	cellar1.items = append(cellar1.items, rdrops)
	cellar2.items = append(cellar2.items, rdrops, rdrops)
	cellar3.items = append(cellar3.items, rdrops, rdrops)
	cellar4.items = append(cellar4.items, rdrops, rdrops, rdrops)
	//ENEMY STRUCTS

	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Millisecond * 300)
	rat1.Ratnamegen()
	time.Sleep(time.Millisecond * 120)
	rand.Seed(time.Now().UTC().UnixNano())
	rat2.Ratnamegen()
	time.Sleep(time.Millisecond * 320)
	rand.Seed(time.Now().UTC().UnixNano())
	rat3.Ratnamegen()
	time.Sleep(time.Millisecond * 110)
	rand.Seed(time.Now().UTC().UnixNano())
	bossrat.Ratnamegen()
	rat1.worthxp, rat2.worthxp, rat3.worthxp, bossrat.worthxp = 400, 400, 400, 800
	rat1.Typeofen, rat2.Typeofen, rat3.Typeofen, bossrat.Typeofen = "Giant Rat", "Giant Rat", "Giant Rat", "King Rat"
	rat1.Typeofatk, rat2.Typeofatk, rat3.Typeofatk, bossrat.Typeofatk = "bite", "bite", "bite", "gnaw"
	rat1.Atk, rat2.Atk, rat3.Atk, bossrat.Atk = 1, 1, 1, 3
	rat1.Battlecry = "Reeettsssttt!! pfftt!!"
	rat2.Battlecry = "ReeeEEEEEttt!!"
	rat3.Battlecry = "KreeeEEEEssstttt!! pfftppfftt"
	bossrat.Battlecry = "KiKreeeEEEEss Eeeep!!"
	rat1.Deathcry = "Eeeeeeeeeppppppsss!! hgffff! pffftt!"
	rat2.Deathcry = "krEeeeeeeeeppppppsss!! eeeeeaapp!!"
	rat3.Deathcry = "crEeeeeeeeettt!! shhhttttt! shtt! pfft!!"
	bossrat.Deathcry = "AAAGGGGhhhcrEeeeeeeeettt!! SSSSShhhtt!!!!! *gurgle*"
	rat1.Health, rat2.Health, rat3.Health, bossrat.Health = 12, 11, 14, 19
	rat1.CurrHealthMax, rat2.CurrHealthMax, rat3.CurrHealthMax, bossrat.CurrHealthMax = 12, 11, 14, 19
	rat1.img = ratimage1
	rat2.img = ratimage2
	rat3.img = ratimage3
	bossrat.img = ratimage4
	cellar1.mobs = []*enemy{rat1, rat1}
	cellar2.mobs = []*enemy{rat2}
	cellar3.mobs = []*enemy{rat3}
	cellar4.mobs = []*enemy{bossrat}
}

func drawtitle() {
	color.Set(color.FgBlue)
	for _, text := range titletext2 {
		fmt.Println(text)
	}

	fmt.Printf("\n\n\n")
	conred()
}

func readblock(t []string) {
	concyan()
	for x := range t {
		for _, text := range t[x] {
			fmt.Print(string(text))
			time.Sleep(time.Millisecond * rndtime())
		}

	}

}
func readline(t string) {
	conwhite()
	fmt.Printf("%v says \" ", p1.Name)
	conyellow()
	for _, x := range t {
		fmt.Print(string(x))
		time.Sleep(time.Millisecond * 20)
	}
	conwhite()
	fmt.Printf("\"\n")
}

func storyintro() {
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	for key := range storyintrotext1 {
		fmt.Print("▐▓")
		for _, text := range storyintrotext1[key] {
			fmt.Print(string(text))
			time.Sleep(time.Millisecond * 10)
		}
		fmt.Print("▓▌")
		time.Sleep(time.Millisecond * 1410)
		fmt.Println(HOME)
		fmt.Print(DELETETOENDOFLINE)
		fmt.Println(HOME)
	}

}

func drawplayerbar(player *playerchar) {
	conblue()
	fmt.Print(ZEROHOME)
	x := len(player.Name)
	fmt.Print("╔" + strings.Repeat("═", x) + "╗\n")
	fmt.Print("║" + strings.Repeat(" ", x) + "║\n")
	fmt.Print("╚" + strings.Repeat("═", x) + "╝\n")
	fmt.Print(playerbaruiboxname)
	conwhite()
	fmt.Print(player.GetName())
	conblue()
	fmt.Print(playerbaruiboxloc1)
	fmt.Print("╔" + strings.Repeat("═", 12) + "╗")
	fmt.Print(playerbaruiboxloc2)
	fmt.Print("║" + strings.Repeat(" ", 12) + "║")
	fmt.Print(playerbaruiboxloc3)
	fmt.Print("╚" + strings.Repeat("═", 12) + "╝")
	fmt.Print(playerbaruiboxloctext)
	conwhite()
	fmt.Printf("Health:%v/%v", player.GetHealth(), player.GetMaxHealth())
	fmt.Print(playerbaruiboxnamedone)

}

func yn() string {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	return string(char)
}

func getcharpause() rune {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	return char
}

//clear entire console window and init CMD environemtn

func drawrand4() int {
	rand.Seed(time.Now().UnixNano())
	boof := rand.Intn(4)
	return boof
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func rndtime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randms := time.Duration(rand.Intn(5))
	return randms
}

func entprompt(thing citentity, dam int) string {

	switch dam {
	case 1, 2, 3, 4, 5:
		promptval := strings.ToUpper("░" + thing.GetName() + "░")
		return promptval
	case 6, 7, 8, 9:
		promptval := strings.ToUpper("░▒│" + thing.GetName() + "│▒░")
		return promptval
	case 10, 11, 12, 13:
		promptval := strings.ToUpper("░▒▓" + thing.GetName() + "▓▒░")
		return promptval
	default:
		promptval := strings.ToUpper("░▒▓▌" + thing.GetName() + "▐▓▒░")
		return promptval

	}

}

func printcompass() {
	conmagenta()
	fmt.Print(COMPASSDRAW1)
	fmt.Print(compass[0])
	fmt.Print(COMPASSDRAW2)
	fmt.Print(compass[1])
	fmt.Print(COMPASSDRAW3)
	fmt.Print(compass[2])
}
