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
var rat1 = new(Enemy)
var rat2 = new(Enemy)
var rat3 = new(Enemy)
var bossrat = new(Enemy)

func (f *Playerchar) SetName(name string) {
	(*f).Name = name
}

func (f *Playerchar) hitroll(max int) bool {
	var results bool
	randplayermiss := rand.Intn(max)
	if randplayermiss <= 2 {
		results = false
	}
	if randplayermiss > 2 {
		results = true
	}
	return results
}
func (e *Enemy) hitroll(max int) bool {
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

func (f *Playerchar) makehealthbar() {
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
func (f *Playerchar) attack(mob *Enemy) {
	clear()
	fmt.Println(ZEROHOME)
	drawplayerbar(f)

	if mob.Health <= 0 {
		navigator(f) //quit out
	}

	for mob.Health > 0 {

		//f.makehealthbar() //check for <=0 HP and draw percentage bar

		if f.hitroll(10) == true {
			fmt.Printf("\t%v's attack hits!! %v damage to %v\n", f.Name, f.Rend(mob), mob.Name)
			mob.Health = mob.Health - f.Rend(mob)
			time.Sleep(time.Millisecond * 1000)
		} else {
			fmt.Printf("\t%v misses\n", f.Name)
			time.Sleep(time.Millisecond * 2000)
		}

		if mob.hitroll(4) == true {
			fmt.Printf("\t%v hits. The %v does %v damage\n", mob.Name, mob.Typeofatk, mob.ratbite(f))
			f.Health = f.Health - int(mob.ratbite(f))
			time.Sleep(time.Millisecond * 1000)
		} else {
			fmt.Printf("\t%v tries to bite and misses.\n", mob.Name)
			time.Sleep(time.Millisecond * 2000)
		}

	}
	color.Set(color.FgRed, color.Bold, color.BgWhite)
	fmt.Printf("\n\n\t%v is slain. %v XP awarded.", mob.Name, mob.worthxp)
	color.Set(color.FgRed, color.Bold, color.BgBlack)
	time.Sleep(time.Millisecond * 2000)
	f.addxp(mob.worthxp)
	f.checklvl()
}

func (f *Playerchar) battleinit() {
	switch f.location {
	case 1:
		f.attack(rat1)
	case 2:
		f.attack(rat2)
	case 3:
		f.attack(rat3)
	case 4:
		f.attack(bossrat)
	}
}

func (f *Playerchar) checklvl() {
	if f.totxp == 1200 {
		f.levup()
	} else if f.totxp == 3000 {
		f.levup()
	}

}
func (f *Playerchar) addxp(xp int) {
	(*f).totxp += xp
}

func (f *Playerchar) levup() {
	f.Level++
	f.Str++
	f.Con++
	f.Atk++
	f.Int++
	f.SetHealth()
	f.Updatehp(f.CurrHealthMax)
	fmt.Println(PLAYERMOVEMENT)
	color.Set(color.FgBlue, color.Bold, color.BgBlack)
	fmt.Printf("%v gained a Level!!!\n", strings.ToUpper(f.Name))
	color.Set(color.FgWhite, color.Bold, color.BgBlack)
	fmt.Printf("Str +1\nCon +1\nAtk+1\nInt+1\n")
	fmt.Printf("New XP total: %v\n", f.totxp)
	time.Sleep(time.Millisecond * 2300)

}

func (f *Playerchar) SetHealth() {
	(*f).Health = int(f.Con * f.Level * 2)
	(*f).CurrHealthMax = f.Health
}

func (f *Playerchar) Updatehp(newtot int) {
	(*f).Health = newtot
}

func (f *Playerchar) Rend(e1 *Enemy) int {
	//return a damage value starting value ~ 4
	return f.Atk + f.Str/10*f.Weapondmg
}

func (f *Playerchar) GetName() string {
	return f.Name
}

func (f *Playerchar) GetHealth() int {
	return f.Health
}

func playerdeath() {
	fmt.Println(ZEROHOME)
	clear()
	fmt.Println("You died alone. Feels Bad Man.")
	os.Exit(0)
}
func initplayername(player *Playerchar) string {

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
func initplayer(player *Playerchar) { //set initial player values
	player.Weapon, player.Weapondmg = "Flail", 2
	player.Str, player.Dex, player.Int, player.Con, player.Level, player.Atk = 10, 10, 10, 10, 1, 2
	player.Area = "Castle Iota"
	player.location = 2
	player.SetHealth()

}

var compass = "  N\nW-╬-E\n  S"

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

	cellar1.title = "Western Cellar Room"
	cellar2.title = "Main Cellar Room"
	cellar3.title = "Eastern Cellar Room"
	cellar4.title = "Sub-Basement"
	cellar1.descrip = "You're standing in the western room of the cellar basement.\nBoxes and crates litter the area. Rat droppings are visible. You're able to move eastward to the central basement.\n"
	cellar2.descrip = "This area the cellar is a large underground area for storage and staging supplies. Water pipes and air ducts are leading in all directions.\nDarkness surrounds everything and you can hear the furnace.\nYou can go East and go deeper into the basement or West towards the storage area of the basement.\nA separate path leads down into the sub-basement.\n"
	cellar3.descrip = "You're in the eastern part of the basement. This part of the basement is desecrated and has been decaying for decades.\nYou can hear the furnace to the west. You can go west back towards the central area of the basement.\n"
	cellar4.descrip = "The sub-basement is littered with rat droppings.\nThe furnance is loud from above and keeps the rats warm in the winter. You can move up the ladder."

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
	rat1.Health, rat2.Health, rat3.Health, bossrat.Health = 8, 9, 8, 16
	cellar1.mobs = []*Enemy{rat1}
	cellar2.mobs = []*Enemy{rat2}
	cellar3.mobs = []*Enemy{rat3}
	cellar4.mobs = []*Enemy{bossrat}
}
func (en *Enemy) ratbite(target *Playerchar) int {
	return en.Atk + int(drawrand4()) //returns uint8 to subtract from playerhp
}

func (en *Enemy) GetHealth() int {
	return en.Health
}
func (en *Enemy) Ratnamegen() {

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

	(*en).Name = ratnames1[random(0, 5)] + ratnames2[random(0, 5)]
}

func drawtitle() {
	color.Set(color.FgBlue)
	for _, text := range titletext {
		fmt.Println(text)
	}
	//titlebar1 := strings.Repeat("═", 140)
	fmt.Println("\n\n")
	color.Set(color.FgRed)
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

func drawplayerbar(player *Playerchar) {
	fmt.Print(ZEROHOME)
	titlebar1 := strings.Repeat("═", 140)
	color.Set(color.FgBlue)
	fmt.Println(titlebar1)
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
func drawplayertitleframe(player *Playerchar) {
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

func OKP() rune {
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
	ANIMATEDELAY        = 100
	ANIMATEDELAYSMALL   = 50
	MINHEALTH           = 0
	BASEDAM             = 6
	ZEROHOME            = "\033[0;0H"
	HOME                = "\033[12;0H" //to line 12
	HIDECURSOR          = "\033[?25l"
	PLAYERINFOHOME      = "\033[11;0H"
	NAVAREA             = "\033[09;0H"
	NAVLOC              = "\033[04;0H"
	NAVDESCRIP          = "\033[11;0H"
	NAVTRAVEL           = "\033[18;0H"
	ENEMIESLISTED       = "\033[19;0H"
	HEALTHBAR           = "\033[4;40H"
	HEALTHBARFILL       = "\033[4;41H"
	NAVTRAVELOUTPUT     = "\033[17;0H"
	PLAYERPROMPT        = "\033[21;0H"
	PLAYERMOVEMENT      = "\033[22;0H"
	ATTACKSTART         = "\033[6;0H"
	COMBATSTART         = "\033[23;0H"
	DELETETOENDOFLINE   = "\033[K"
	SAVECURSOR          = "\033[s"
	RETURNTOSAVEDCURSOR = "\033[u"
	RESETCOLORS         = "\033[0m"
)

func rndtime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randms := time.Duration(rand.Intn(10))
	return randms
}
