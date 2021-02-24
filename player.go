package citadel

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/dariubs/percent"
	"github.com/fatih/color"
)

func (f *playerchar) SetName(name string) {
	(*f).Name = name
}

func (f *playerchar) getbattlerap() {

	x := []string{
		0:  "You shall not pass.",
		1:  "Take your stinking paws off me, you damned dirty ape!",
		2:  "Soylent Green is people!",
		3:  "Time is a flat circle.",
		4:  "Hell is just a word. The reality is much worse.",
		5:  "Liberate tu ta me ex inferis!",
		6:  "You look tired.",
		7:  "You have failed miserably. Lesson is, never try.",
		8:  "Have at thee!",
		9:  "Do not take me for some conjuror of cheap tricks.",
		10: "I hate you.",
		11: "I love you.",
		12: "Bad Timing innit?",
		13: "You're below average at best.",
		14: "Slow down there partner.",
		15: "Turn in your badge and gun.",
		16: "It hurts you more than it hurts me.",
		17: "Say hello to my little friend.",
		18: "(RatsNameHere) is bad at this.",
		19: "I heard something today, it sounded like this: Enemys in the game called Citadel really suck.",
	}
	string := x[rand.Intn(20)]
	y := rand.Intn(12)

	if y > 9 {
		readline(string)
	}
}

func (f *playerchar) getdamageroll() int {
	return f.Atk + f.Str/10*f.Weapondmg + int(drawrand4())
}

func (f *playerchar) lookmap(loc *maploc) {
	readblock(loc.lookdescrip)
}

func (f *playerchar) getitems(items []*item) {
	for _, item := range items {
		f.inv = append(f.inv, item)
	}

}
func (f *playerchar) addxp(xp int) {
	(*f).totxp += xp
	(*f).kills++
}
func (f *playerchar) listinv(items []*item) {
	clear()
	congreen()
	fmt.Println("Carrying:")
	concyan()
	for _, y := range items {
		fmt.Println(y.name)
	}

	fmt.Println("\n\n▒▓▌PRESS ANY KEY TO CONTINUE▐▓▒")
	getcharpause()
}

func (f *playerchar) lookroom(x uint8) {
	clear()
	drawplayerbar(f)
	conblue()
	fmt.Printf("\n\nYou look around and search the area...\n\n")

	switch x {
	case 1:
		cellar1.listitems()
		f.lookenemy(rat1)
		congreen()
		f.lookmap(cellar1)
		itemmenu(f, cellar1.items)
	case 2:
		cellar2.listitems()
		f.lookenemy(rat2)
		congreen()
		f.lookmap(cellar2)
		itemmenu(f, cellar2.items)
	case 3:
		cellar3.listitems()
		f.lookenemy(rat3)
		congreen()
		f.lookmap(cellar3)
		itemmenu(f, cellar3.items)
	case 4:
		cellar4.listitems()
		f.lookenemy(bossrat)
		congreen()
		f.lookmap(cellar4)
		itemmenu(f, cellar4.items)
	}

}

func (f *playerchar) lookenemy(e *enemy) {
	x := e.examine()
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

func (f *playerchar) isalive() bool {
	x := true
	if f.Health <= 0 {
		playerdeath()
		x = false
	}
	return x
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
	x := f.getdamageroll()
	e.Updatehp(e.Health - x)
	f.getbattlerap()
	congreen()
	fmt.Printf("%v«%v» %v %v with %v for ", entprompt(f, x), f.GetHealth(), f.attackstring(), e.GetName(), f.Weapon)
	conred()
	fmt.Printf("%v\n", strconv.Itoa(x))
	congreen()
	time.Sleep(time.Millisecond * 2000)

}

func (f *playerchar) brmiss(e *enemy) {
	conyellow()
	fmt.Printf("%v«%v» %v %v\n", entprompt(f, 1), f.GetHealth(), f.miss(), e.GetName())
	time.Sleep(time.Millisecond * 2000)

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
	conyellow()
	fmt.Println("\nYou recover from battle and ready yourself")
	time.Sleep(time.Millisecond * 200)
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
	congreen()
	fmt.Printf("New XP total: %v\n", f.gettotalxp())
	conmagenta()
	fmt.Printf("New max health: %v\n", f.GetMaxHealth())
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
func (f *playerchar) GetHealth() string {
	x := strconv.Itoa(f.Health)
	return x
}

func (f *playerchar) GetMaxHealth() string {
	x := strconv.Itoa(f.CurrHealthMax)
	return x
}

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
