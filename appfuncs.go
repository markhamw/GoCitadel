package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/dariubs/percent"
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

func (f *playerchar) SetName(name string) {
	(*f).Name = name
}

func (f *playerchar) getdamageroll() int {
	return f.Atk + f.Str/10*f.Weapondmg + int(drawrand4())
}
func (e *enemy) getdamageroll() int {
	return e.Atk + int(drawrand4())
}
func (f *playerchar) lookmap(loc *maploc) string {
	return loc.lookdescrip
}
func (f *playerchar) lookroom(x uint8) {
	clear()
	drawplayerbar(f)
	fmt.Printf("\n\nYou look around and search the area...\n\n")
	switch x {
	case 1:
		f.lookenemy(rat1)
		congreen()
		fmt.Println(f.lookmap(cellar1))
	case 2:
		f.lookenemy(rat2)
		congreen()
		fmt.Println(f.lookmap(cellar2))
	case 3:
		f.lookenemy(rat3)
		congreen()
		fmt.Println(f.lookmap(cellar3))
	case 4:
		f.lookenemy(bossrat)
		congreen()
		fmt.Println(f.lookmap(cellar4))
	}
	conwhite()
	fmt.Println("▒▓▌PRESS ANY KEY TO CONTINUE▐▓▒")
	getcharpause()
}

func (f *playerchar) lookenemy(e *enemy) {
	x := e.isalive()
	if x {
		conmagenta()
		fmt.Printf("%v is here.\n\n\n", e.Name)
		for _, x := range e.img {

			fmt.Println(x)
		}
		fmt.Print("\n\n")
	} else if !x {
		conwhite()
		fmt.Printf("\n\nThe remains of %v are scattered. %v's corpse is resting here.\n\n", e.Name, e.Name)
		fmt.Print("")
	}

}

func (f *playerchar) hitroll(max int) bool {
	var results bool
	x := rand.Intn(max)
	if x <= 2 {
		results = false
	}
	if x > 2 {
		results = true
	}
	return results
}

func isstanding(character citentity) bool {
	if character.isalive() {
		return true
	}
	return false
}

func (e *enemy) isalive() bool {
	var x bool
	if e.Health > 0 {
		x = true
	} else if e.Health <= 0 {
		x = false
	}
	return x
}

func (f *playerchar) isalive() bool {
	var x bool
	if f.Health > 0 {
		x = true
	} else if f.Health <= 0 {
		x = false
	}
	return x
}

func (e *enemy) hitroll(max int) bool {
	var results bool
	randenemymiss := rand.Intn(max)
	if randenemymiss <= 2 {
		results = false
	}
	if randenemymiss > 2 {
		results = true
	}
	return results
}

func (f *playerchar) makehealthbar() {
	if f.Health <= 0 {
		playerdeath()
	}
	fmt.Printf(HEALTHBAR)
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Printf("[")
	fmt.Printf(strings.Repeat(" ", int(percent.PercentOf(f.CurrHealthMax, f.CurrHealthMax))/10))
	fmt.Printf("]")
	fmt.Printf(HEALTHBARFILL)
	fmt.Printf(strings.Repeat("▓", int(percent.PercentOf(f.Health, f.CurrHealthMax))/10))
}

func (f *playerchar) brpds(e *enemy) {
	color.Set(color.FgGreen, color.Bold, color.BgBlack)

	if (f.isalive() && e.isalive()) == true {
		x := f.getdamageroll()
		e.Updatehp(e.Health - x)
		fmt.Printf("\n%v«%v» %v %v with %v for %v\n", entprompt(f), f.GetHealth(), f.attackstring(), e.GetName(), f.Weapon, strconv.Itoa(x))
		time.Sleep(time.Millisecond * 2000)

	} else if f.isalive() == false {
		playerdeath()
	} else if e.isalive() == false {
		e.reportdeath(f)
	}

} //battle report details string

func (f *playerchar) brmiss(e *enemy) {

	both := (e.isalive() && f.isalive())
	pl1dead := !(f.isalive())
	en1dead := !(e.isalive())

	if both {
		color.Set(color.FgGreen, color.Bold, color.BgBlack)
		fmt.Printf("\n%v«%v» %v %v\n", entprompt(f), f.GetHealth(), f.miss(), e.GetName())
		time.Sleep(time.Millisecond * 2000)
	} else if pl1dead {
		playerdeath()
	} else if en1dead {
		e.reportdeath(f)
	}
}

func (e *enemy) enemybrpds(p *playerchar) {
	conwhite()

	if (e.isalive() && p.isalive()) == true {

		x := e.getdamageroll()
		p.Updatehp(p.Health - x)
		fmt.Printf("\n%v«%v» %v %v for %v\n", entprompt(e), e.GetHealth(), e.attackstring(), p.GetName(), strconv.Itoa(x))
		time.Sleep(time.Millisecond * 2000)

	} else if p.isalive() == false {
		playerdeath()
	} else if e.isalive() == false {
		e.reportdeath(p)

	}

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
	case 'l':
		player.lookroom(player.location)

	}
	time.Sleep(time.Millisecond * 400)
	navigator(player)
}

func (l maploc) printdescrip() {
	fmt.Print(NAVDESCRIP)
	fmt.Printf(" ¬│░░░░▒▒▒▒▒▒▓▓▓▓▓▌DESCRIPTION▐▓▓▓▓▓▒▒▒▒▒▒░░░░│⌐ \n")
	fmt.Print(l.descrip)
}

func navigator(player *playerchar) {
	clear()
	drawplayerbar(player)
	player.makehealthbar()
	color.Set(color.FgYellow, color.Bold, color.BgBlack)

	switch player.location {
	case 1:
		generatemaplocdata(cellar1, player)
		checkbattle(player, rat1)
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
func (e *enemy) enemybrpdsmiss(p *playerchar) {
	if (e.isalive() && p.isalive()) == true {

		color.Set(color.FgWhite, color.Bold, color.BgBlack)
		fmt.Printf("\n%v«%v» %v %v\n", entprompt(e), e.GetHealth(), e.miss(), p.GetName())
		time.Sleep(time.Millisecond * 2000)

	} else if p.isalive() == false {
		playerdeath()
	} else if e.isalive() == false {
		e.reportdeath(p)
	}
}

func battreport(p *playerchar, en *enemy) {
	conwhite()
	fmt.Printf("\n%v engages in combat with %v\n\n", p.Name, en.Name)
	conred()
	fmt.Printf("%v charges into battle screaming \"%v\"\n\n", en.Name, en.Battlecry)
	for en.Health > 0 {
		if p.hitroll(10) {
			p.brpds(en)
		} else {
			p.brmiss(en)
		}

		if en.hitroll(7) { //difficulty modifier?? //increase to increase liklihood of hitting
			en.enemybrpds(p)
		} else {
			en.enemybrpdsmiss(p)
		}
	}

}

func (e *enemy) miss() string {
	x := []string{
		0: "misses",
		1: "almost gnaws",
		2: "squeeks and grits its teeth angrily at",
		3: "gets dodged on by",
		4: "runs around aimlessly near",
	}
	return x[rand.Intn(5)]
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

	time.Sleep(time.Millisecond * 100)
}

func (f *playerchar) miss() string {
	x := []string{
		0: "misses",
		1: "barely misses",
		2: "fumbles and misses",
		3: "gets dodged on by",
		4: "almost hit",
	}
	return x[rand.Intn(5)]
}

func (e *enemy) attackstring() string {
	x := []string{
		0: "nips",
		1: "slashes",
		2: "bites",
		3: "chomps",
		4: "pisses off",
	}
	return x[rand.Intn(5)]
}
func (f *playerchar) attackstring() string {
	x := []string{
		0: "crushes",
		1: "hits",
		2: "damages",
		3: "smashes",
		4: "wollops",
	}
	return x[rand.Intn(5)]
}

func (f *playerchar) attack(mob *enemy) {
	clear()
	fmt.Println(ZEROHOME)
	drawplayerbar(f)
	for mob.Health > 0 {

		battreport(f, mob)

	}
	fmt.Println(".")
}

func generatemaplocdata(loc *maploc, p *playerchar) {
	loc.printmap()
	fmt.Print(NAVTRAVEL)

	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Print("Use n,s,e,w ░ u(up),d(down) ░ b=battle ░\ni=check iventory ░ q=char info ░ l=look")
	p.checklvl()
}

func (l *maploc) printmap() {
	fmt.Print(NAVLOC)
	conred()
	fmt.Printf("Detailed Map:\n%v\n", l.grid)
	conwhite()
	fmt.Print(NAVAREA)
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
	fmt.Printf("╠═%v═╣\n\n", l.title)
	color.Set(color.FgCyan, color.Bold, color.BgBlack)

	l.printdescrip()
	printcompass()
}

func (f *playerchar) getmaploc() string {
	var x string
	switch f.location {
	case 1:
		x = cellar1.title
	case 2:
		x = cellar2.title
	case 3:
		x = cellar3.title
	case 4:
		x = cellar4.title
	}
	return x
}

func (f *playerchar) battleinit() {

	switch f.location {
	case 1:
		if rat1.isalive() {
			f.attack(rat1)
		} else {
			navigator(f)
		}

	case 2:
		if rat2.isalive() {
			f.attack(rat2)
		} else {
			navigator(f)
		}
	case 3:
		if rat3.isalive() {
			f.attack(rat3)
		} else {
			navigator(f)
		}
	case 4:
		if bossrat.isalive() {
			f.attack(bossrat)
		} else {
			navigator(f)
		}
	}
}

var lvl1 = true
var lvl2 = true

func (f *playerchar) checklvl() {
	if f.totxp == 1200 {

		if lvl1 {

			f.levup()
			lvl1 = false
		} else if !lvl1 {
			return
		}

	}
	if f.totxp == 3000 {
		if lvl2 {
			f.levup()
			lvl2 = false
		}

	}

}

func (f *playerchar) addxp(xp int) {
	(*f).totxp += xp
	(*f).kills++
}

func (f *playerchar) levup() {
	f.Level++
	f.Str++
	f.Con++
	f.Atk++
	f.Int++
	f.SetHealth()
	f.Updatehp(f.CurrHealthMax)
	conblue()
	fmt.Printf("\n%v gained a Level!!!\n", strings.ToUpper(f.Name))
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Printf("Str +1\nCon +1\nAtk+1\nInt+1\n")
	fmt.Printf("New XP total: %v\n", f.gettotalxp())
	time.Sleep(time.Millisecond * 2300)

}
func (f *playerchar) gettotalxp() string {
	return strconv.Itoa(f.totxp)
}
func (f *playerchar) printstats() {
	clear()
	drawplayerbar(f)
	color.Set(color.FgMagenta, color.Bold, color.BgBlack)
	fmt.Printf("\nBase Atk: %v (Using %v)\n\n", f.Atk, f.Weapon)
	fmt.Printf("░ STR:%v ░ DEX:%v ░ CON:%v ░ INT:%v ░\n\n", f.Str, f.Dex, f.Con, f.Int)
	fmt.Printf("Total XP: %v\n", f.gettotalxp())
	fmt.Printf("Total Kills: %v\n\n", f.kills)
	fmt.Printf("Current location: %v\n\n", f.getmaploc())
	fmt.Println("PRESS ANY KEY TO CONTINUE..")
	getcharpause()
}

func (f *playerchar) SetHealth() {
	(*f).Health = int(f.Con * f.Level * 2)
	(*f).CurrHealthMax = f.Health
}

func (f *playerchar) Updatehp(newtot int) {
	(*f).Health = newtot
}

func (f *playerchar) Rend(e *enemy) int {
	//return a damage value starting value ~ 4
	return f.Atk + f.Str/10*f.Weapondmg
}

func (f *playerchar) GetName() string {
	return f.Name
}
func (e *enemy) reportdeath(p *playerchar) {
	conred()
	fmt.Printf("%v dies and cries out: \"%v\"", e.Name, e.Deathcry)
	color.Set(color.FgMagenta, color.Bold, color.BgBlack)
	fmt.Printf("\n\n%v is slain. %v XP awarded.", e.Name, e.worthxp)
	time.Sleep(time.Millisecond * 2000)
	p.addxp(e.getxpworth())
	time.Sleep(time.Millisecond * 500)
	p.checklvl()
}

func (e *enemy) getxpworth() int {
	return e.worthxp
}
func (e *enemy) GetName() string {
	return e.Name
}
func (f *playerchar) GetHealth() string {
	x := strconv.Itoa(f.Health) + "HP"
	return x
}

func playerdeath() {
	fmt.Println(ZEROHOME)
	clear()
	fmt.Println("You died alone. Feels Bad Man.")
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
	conwhite()
	fmt.Println(castle)
	conwhite()
	readblock(castletext)

}

func initplayer(player *playerchar) { //set initial player values
	player.Weapon, player.Weapondmg = "Flail", 2
	player.Str, player.Dex, player.Int, player.Con, player.Level, player.Atk = 10, 10, 10, 10, 1, 2
	player.Area = "Castle Iota"
	player.location = 2
	player.SetHealth()

}

var compass = []string{0: "N", 1: "W-╬-E", 2: "S"}

func initworld() { //creates all maplocs and enemys
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

	cellar1.lookdescrip = "Vermin have clearly attempted to re-purpose\nthis area as their own. Useless containers line the\nwalls and a few have bite marks. Important containers\nare encased in Elder Wax for safe storage.\n"
	cellar2.lookdescrip = "Fire light is escaping from furnace vents.\nThe room is large and the ceiling is lost to darkness.\nThe sub-basement is open and is usually kept closed.\nYou recall that the sub-basement might be dangerous."
	cellar3.lookdescrip = "Cellar walls are crumbling and masons have\ntried to repair the failing infrastructure. It appears\nvermin have piled bits of bone and food in this part\nof the cellar."
	cellar4.lookdescrip = "The sub-basement is a stately room for rulers\nof rats and vermin.  A RUS-sized hole leads to the other\npathways that are too small to fit through. Bits\nof hair and bone and rotting rat treasures are present."

	//ENEMY STRUCTS
	rand.Seed(21)
	time.Sleep(time.Millisecond * 300)
	rat1.Ratnamegen()
	time.Sleep(time.Millisecond * 120)
	rand.Seed(time.Now().UTC().UnixNano())
	rat2.Ratnamegen()
	time.Sleep(time.Millisecond * 320)
	rand.Seed(8822)
	rat3.Ratnamegen()
	time.Sleep(time.Millisecond * 110)
	rand.Seed(time.Now().UnixNano())
	bossrat.Ratnamegen()
	rat1.worthxp, rat2.worthxp, rat3.worthxp, bossrat.worthxp = 400, 400, 400, 800
	rat1.Typeofen, rat2.Typeofen, rat3.Typeofen, bossrat.Typeofen = "Giant Rat", "Giant Rat", "Giant Rat", "King Rat"
	rat1.Typeofatk, rat2.Typeofatk, rat3.Typeofatk, bossrat.Typeofatk = "bite", "bite", "bite", "gnaw"
	rat1.Atk, rat2.Atk, rat3.Atk, bossrat.Atk = 1, 1, 1, 3
	rat1.Battlecry = "Reeettsssttt!!"
	rat2.Battlecry = "ReeeEEEEEttt!!"
	rat3.Battlecry = "KreeeEEEEssstttt!!"
	bossrat.Battlecry = "KiKreeeEEEEss Eeeep!!"
	rat1.Deathcry = "Eeeeeeeeeppppppsss!! hgffff!"
	rat2.Deathcry = "krEeeeeeeeeppppppsss!! eeeeeaapp!!"
	rat3.Deathcry = "crEeeeeeeeettt!! shhhttttt! shtt!"
	bossrat.Deathcry = "AAAGGGGGGGGhhhhhhhhcrEeeeeeeeettt!! SSSSSttttt!!!!!!!! *gurgle*"
	rat1.Health, rat2.Health, rat3.Health, bossrat.Health = 12, 11, 14, 19
	rat1.CurrHealthMax, rat2.CurrHealthMax, rat3.CurrHealthMax, bossrat.CurrHealthMax = 12, 11, 14, 19
	rat1.img = ratimage1
	rat2.img = ratimage2
	rat3.img = ratimage3
	bossrat.img = ratimage4
	cellar1.mobs = []*enemy{rat1}
	cellar2.mobs = []*enemy{rat2}
	cellar3.mobs = []*enemy{rat3}
	cellar4.mobs = []*enemy{bossrat}
}
func (e *enemy) Rend(target citentity) int {
	return e.Atk + int(drawrand4()) //returns uint8 to subtract from playerhp
}

func (e *enemy) GetHealth() string {
	x := strconv.Itoa(e.Health) + "HP"
	return x
}
func (e *enemy) Updatehp(x int) {
	(*e).Health = x
}

func (e *enemy) Ratnamegen() {

	ratnames1 := []string{
		0: "Bo",
		1: "Mo",
		2: "Do",
		3: "Ka",
		4: "Ret",
		5: "El",
		6: "Ra",
		7: "Rump",
	}
	ratnames2 := []string{
		0: "jaki",
		1: "nold",
		2: "tikt",
		3: "jhilz",
		4: "kansta",
		5: "fildo",
		6: "bhendu",
		7: "nado",
	}

	(*e).Name = ratnames1[random(0, 7)] + ratnames2[random(0, 7)]
}

func drawtitle() {
	color.Set(color.FgBlue)
	for _, text := range titletext2 {
		fmt.Println(text)
	}
	//titlebar1 := strings.Repeat("═", 140)
	fmt.Printf("\n\n\n")
	conred()
}

func readblock(t []string) {
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	for x := range t {
		for _, text := range t[x] {
			fmt.Print(string(text))
			time.Sleep(time.Millisecond * rndtime())
		}

	}

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
	fmt.Print(ZEROHOME)
	titlebar1 := "\n░░░░░░░▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓█████▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒░░░░░░░"
	color.Set(color.FgBlue)
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print("PLAYER:")
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Print(player.Name+"(lvl"+strconv.Itoa(int(player.Level))+")", "\t")
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print("\tMax HP:")
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Print(player.CurrHealthMax)
	color.Set(color.FgBlue)
	fmt.Print(titlebar1)

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
func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print(HIDECURSOR)

}

func drawrand4() int {
	rand.Seed(time.Now().UnixNano())
	boof := rand.Intn(4)
	return boof
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

const (
	ANIMATEDELAY      = 100           //as
	ANIMATEDELAYSMALL = 50            //
	MINHEALTH         = 0             //
	BASEDAM           = 6             //
	ZEROHOME          = "\033[0;0H"   //
	HOME              = "\033[12;0H"  //to line 12
	HIDECURSOR        = "\033[?25l"   //
	PLAYERINFOHOME    = "\033[11;0H"  //
	NAVAREA           = "\033[09;16H" //
	NAVLOC            = "\033[04;0H"
	COMPASSDRAW1      = "\033[4;26H" //
	COMPASSDRAW2      = "\033[5;24H" //
	COMPASSDRAW3      = "\033[6;26H" //
	NAVDESCRIP        = "\033[11;0H" //
	NAVTRAVEL         = "\033[24;0H" //
	ENEMIESLISTED     = "\033[19;0H" //
	HEALTHBAR         = "\033[4;40H" //
	HEALTHBARFILL     = "\033[4;41H" //
	NAVTRAVELOUTPUT   = "\033[26;0H" //
	ATTACKSTART       = "\033[6;0H"  //
	//COMBATSTART         = "\033[23;0H" //
	DELETETOENDOFLINE   = "\033[K"  //
	SAVECURSOR          = "\033[s"  //
	RETURNTOSAVEDCURSOR = "\033[u"  //
	RESETCOLORS         = "\033[0m" //
)

func rndtime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randms := time.Duration(rand.Intn(10))
	return randms
}
func conred() {
	color.Set(color.FgRed, color.Bold, color.BgBlack)
}
func conwhite() {
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
}
func conmagenta() {
	color.Set(color.FgMagenta, color.Bold, color.BgBlack)
}

func congreen() {
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
}
func conblue() {
	color.Set(color.FgBlue, color.Bold, color.BgBlack)
}
func entprompt(thing citentity) string {
	promptval := strings.ToUpper("░▒▓▌" + thing.GetName() + "▐▓▒░")
	return promptval
}
func printcompass() {
	conwhite()
	fmt.Print(COMPASSDRAW1)
	fmt.Print(compass[0])
	fmt.Print(COMPASSDRAW2)
	fmt.Print(compass[1])
	fmt.Print(COMPASSDRAW3)
	fmt.Print(compass[2])
}
