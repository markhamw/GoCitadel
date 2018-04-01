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
func (f *playerchar) updatehealthbar() {
	if f.Health <= 0 {
		playerdeath()
	}
	fmt.Printf(HEALTHBAR)
	fmt.Print(DELETETOENDOFLINE)
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Printf("[")
	fmt.Printf(strings.Repeat(" ", int(percent.PercentOf(f.CurrHealthMax, f.CurrHealthMax))/10))
	fmt.Printf("]")
	fmt.Printf(HEALTHBARFILL)
	fmt.Printf(strings.Repeat("▓", int(percent.PercentOf(f.Health, f.CurrHealthMax))/10))
	fmt.Print(RETURNTOSAVEDCURSOR)
}
func (e *enemy) updatehealthbar() {
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	if e.Health <= 0 {
		e.reportdeath()
	}
	fmt.Printf(ENEMYHEALTHBAR)
	fmt.Print(DELETETOENDOFLINE)
	conred()
	fmt.Printf("[")
	fmt.Printf(strings.Repeat(" ", int(percent.PercentOf(e.CurrHealthMax, e.CurrHealthMax))/10))
	fmt.Printf("]")
	fmt.Printf(ENEMYHEALTHBARFILL)
	fmt.Printf(strings.Repeat("▓", int(percent.PercentOf(e.Health, e.CurrHealthMax))/10))
	fmt.Print(RETURNTOSAVEDCURSOR)
}
func (e *enemy) makehealthbar() {
	if e.Health <= 0 {
		e.reportdeath()
	}
	fmt.Print(ENEMYHEALTHBAR)
	conred()
	fmt.Printf("[")
	fmt.Printf(strings.Repeat(" ", int(percent.PercentOf(e.CurrHealthMax, e.CurrHealthMax))/10))
	fmt.Printf("]")
	fmt.Print(ENEMYHEALTHBARFILL)
	fmt.Printf(strings.Repeat("▓", int(percent.PercentOf(e.Health, e.CurrHealthMax))/10))
}

func (f *playerchar) brpds(e *enemy) {
	if f.isalive() && e.isalive() {
		color.Set(color.FgGreen, color.Bold, color.BgBlack)

		x := f.getdamageroll()
		e.Updatehp(e.Health - x)

		fmt.Printf("\n\t%v«%v» %v %v with %v for %v\n", entprompt(f), f.GetHealth(), f.attackstring(), e.GetName(), f.Weapon, strconv.Itoa(x))
		time.Sleep(time.Millisecond * 2000)
	} else {
		playerdeath()
	}

} //battle report details string

func (f *playerchar) brmiss() { //battle report MISS
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Printf("\n\t%v«%v» %v\n", entprompt(f), f.GetHealth(), f.miss())
	time.Sleep(time.Millisecond * 2000)
}

func (e *enemy) enemybrpds(p *playerchar) {
	if e.isalive() && p.isalive() {
		conred()
		x := e.getdamageroll()
		p.Updatehp(p.Health - x)
		fmt.Printf("\n\t%v«%v» %v %v for %v\n", entprompt(e), e.GetHealth(), e.attackstring(), p.GetName(), strconv.Itoa(x))
		time.Sleep(time.Millisecond * 2000)
	} else {
		e.reportdeath()
	}

}

func (e *enemy) enemybrpdsmiss() {
	conred()
	fmt.Printf("\n\t%v«%v» %v %v\n", entprompt(e), e.GetHealth(), e.GetName(), e.miss())
	time.Sleep(time.Millisecond * 2000)
}

func battreport(p *playerchar, en *enemy) {
	for en.Health > 0 {
		if p.hitroll(10) {
			p.brpds(en)
		} else {
			p.brmiss()
		}

		if en.hitroll(7) { //difficulty modifier?? //increase to increase liklihood of hitting
			en.enemybrpds(p)
		} else {
			en.enemybrpdsmiss()
		}
	}

}

func (e *enemy) miss() string {
	x := []string{
		0: "tries to bite and misses",
		1: "tries to hit!, and misses.",
		2: "squeeks and grits its teeth.",
		3: "gets dodged on.",
		4: "leaves a dropping",
	}
	return x[rand.Intn(5)]
}

func (f *playerchar) miss() string {
	x := []string{
		0: "misses.",
		1: "tries to hit!, and misses.",
		2: "grunts and swings wide.",
		3: "gets dodged on.",
		4: "almost hit!",
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
		0: "grazes",
		1: "hits",
		2: "damages",
		3: "deflects",
		4: "wollops",
	}
	return x[rand.Intn(5)]
}

func (f *playerchar) attack(mob *enemy) {
	clear()
	fmt.Println(ZEROHOME)
	drawplayerbar(f)
	for mob.Health > 0 {
		if f.isalive() == false {
			playerdeath()
		} else {
			battreport(f, mob)
		}

	}

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

func (f *playerchar) checklvl() {
	if f.totxp == 1200 {
		f.levup()
	} else if f.totxp == 3000 {
		f.levup()
	}

}
func (f *playerchar) addxp(xp int) {
	(*f).totxp += xp
}

func (f *playerchar) levup() {
	f.Level++
	f.Str++
	f.Con++
	f.Atk++
	f.Int++
	f.SetHealth()
	f.Updatehp(f.CurrHealthMax)
	color.Set(color.FgBlue, color.Bold, color.BgBlack)
	fmt.Printf("%v gained a Level!!!\n", strings.ToUpper(f.Name))
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Printf("Str +1\nCon +1\nAtk+1\nInt+1\n")
	fmt.Printf("New XP total: %v\n", f.totxp)
	time.Sleep(time.Millisecond * 2300)

}
func (f *playerchar) printstats() {
	clear()
	drawplayerbar(f)
	color.Set(color.FgMagenta, color.Bold, color.BgBlack)
	fmt.Print(entprompt(f))
	fmt.Printf("\nBase Attack Power: %v (Using %v)\n\n", f.Atk, f.Weapon)
	fmt.Printf("░ STR:%v ░ DEX:%v ░ CON:%v ░ INT:%v ░\n\n", f.Str, f.Dex, f.Con, f.Int)
	fmt.Printf("Total XP: %v\n\n", f.totxp)
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
func (e *enemy) reportdeath() {
	fmt.Printf("\t%v dies and cries out: \"%v\"", e.Name, e.Deathcry)
	color.Set(color.FgMagenta, color.Bold, color.BgBlack)
	fmt.Printf("\n\n\t%v is slain. %v XP awarded.", e.Name, e.worthxp)
	time.Sleep(time.Millisecond * 2000)
	p1.addxp(e.worthxp)

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

	fmt.Print("\033[25;27H")
	fmt.Print(SAVECURSOR)
	var playername string
	fmt.Scan(&playername)
	fmt.Print(RETURNTOSAVEDCURSOR)
	fmt.Printf("Use %v? (PRESS ANY KEY TO CONTINUE, or n to try again)", playername)
	answer, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	answerstring := string(answer)

	switch answerstring {

	case "n", "N":
		fmt.Print("\033[25;27H")
		fmt.Print(DELETETOENDOFLINE) //repos cursor to after yn
		fmt.Print("\033[25;27H")
		fmt.Scan(&playername)
	case "y", "Y":
		return playername
	default:
		return playername
	}

	return playername
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
	cellar1.descrip = "You're standing in the western room of the cellar\nbasement. Boxes and crates are stacked the area. Rats have gnawed the\nboxes to bits. You're able to move eastward to the central basement.\n"

	cellar2.descrip = "This area the cellar is a large underground area\nfor storage and staging supplies. Water pipes and\nair ducts are leading in all directions. Darkness\nsurrounds everything and you can hear the furnace.\nYou can go East and go deeper into the basement or\nWest towards the storage area of the basement. A\nseparate path leads down into the sub-basement.\n"

	cellar3.descrip = "You're in the eastern part of the basement. This\npart of the basement is desecrated and has been\ndecaying for decades. You can hear the furnace to\nthe west. You can go west back towards the central\narea of the basement.\n"

	cellar4.descrip = "The sub-basement is littered with rat corpses and\npiles of rat treasures and rotting detritus. The\nfurnance is loud from above and keeps the rats warm\nin the winter. A ladder leads back up to the Cellar.\n"

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
	rat1.Deathcry = "Eeeeeeeeeppppppsss!!"
	rat2.Deathcry = "krEeeeeeeeeppppppsss!!"
	rat3.Deathcry = "crEeeeeeeeettt!!"
	bossrat.Deathcry = "crEeeeeeeeettt!!"
	rat1.Health, rat2.Health, rat3.Health, bossrat.Health = 12, 11, 14, 19
	rat1.CurrHealthMax, rat2.CurrHealthMax, rat3.CurrHealthMax, bossrat.CurrHealthMax = 12, 11, 14, 19
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
	}
	ratnames2 := []string{
		0: "jaki",
		1: "nold",
		2: "tikt",
		3: "jhilz",
		4: "kansta",
	}

	(*e).Name = ratnames1[random(0, 5)] + ratnames2[random(0, 5)]
}

func drawtitle() {
	color.Set(color.FgBlue)
	for _, text := range titletext {
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

	//inline func to pause random 170ms
	randtime := func() time.Duration {
		rand.Seed(time.Now().UTC().UnixNano())
		//sets type for proper return value
		randms := time.Duration(rand.Intn(10))
		return randms
	}
	for key := range storyintrotext1 {

		for _, text := range storyintrotext1[key] {

			fmt.Print(string(text))
			time.Sleep(time.Millisecond * randtime())
		}

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
func drawplayertitleframe(player *playerchar) {
	skulldrawdown()
	titlebar1 := strings.Repeat("═", 140)
	fmt.Print(PLAYERINFOHOME) //removed Println
	fmt.Print(DELETETOENDOFLINE)
	//color.Set(color.FgBlue)
	//fmt.Println(titlebar1)
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print("\tCharacter:")
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Print(player.Name+"(lvl"+strconv.Itoa(int(player.Level))+")", "\t")
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print("\tMax HP:")
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Print(player.CurrHealthMax, "\t")
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print("\tWielding:")
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Print(player.Weapon, "(", player.Weapondmg, "dmg)", "\t")
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Print("\tLocation:")
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
	fmt.Print(player.Area, "\n")

	//Draw bottom blue bar
	color.Set(color.FgBlue)
	fmt.Println(titlebar1)
	color.Set(color.FgWhite, color.Bold, color.BgBlack)

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

const ( //test comment

	ANIMATEDELAY        = 100           //as
	ANIMATEDELAYSMALL   = 50            //
	MINHEALTH           = 0             //
	BASEDAM             = 6             //
	ZEROHOME            = "\033[0;0H"   //
	HOME                = "\033[12;0H"  //to line 12
	HIDECURSOR          = "\033[?25l"   //
	PLAYERINFOHOME      = "\033[11;0H"  //
	NAVAREA             = "\033[09;16H" //
	NAVLOC              = "\033[04;0H"
	COMPASSDRAW1        = "\033[4;26H" //
	COMPASSDRAW2        = "\033[5;24H" //
	COMPASSDRAW3        = "\033[6;26H" //
	NAVDESCRIP          = "\033[11;0H" //
	NAVTRAVEL           = "\033[18;0H" //
	ENEMIESLISTED       = "\033[19;0H" //
	HEALTHBAR           = "\033[4;40H" //
	HEALTHBARFILL       = "\033[4;41H" //
	ENEMYHEALTHBAR      = "\033[6;40H" //
	ENEMYHEALTHBARFILL  = "\033[6;41H" //
	NAVTRAVELOUTPUT     = "\033[17;0H" //
	ATTACKSTART         = "\033[6;0H"  //
	COMBATSTART         = "\033[23;0H" //
	DELETETOENDOFLINE   = "\033[K"     //
	SAVECURSOR          = "\033[s"     //
	RETURNTOSAVEDCURSOR = "\033[u"     //
	RESETCOLORS         = "\033[0m"    //
)

func rndtime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randms := time.Duration(rand.Intn(10))
	return randms
}
