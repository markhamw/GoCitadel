package citadel

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func (e *enemy) getdamageroll() int {
	return e.Atk + int(drawrand4())
}

func (e *enemy) isalive() bool {
	x := true

	if e.Health <= 0 {
		e.reportdeath(p1)
		navigator(p1)
	} else if x == true {
		x = true
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

func (e *enemy) enemybrpds(p *playerchar) {
	concyan()
	x := e.getdamageroll()
	p.Updatehp(p.Health - x)
	fmt.Printf("%v«%v» %v %v for ", entprompt(e, x), e.GetHealth(), e.attackstring(), p.GetName())
	conmagenta()
	fmt.Printf("%v\n", strconv.Itoa(x))
	time.Sleep(time.Millisecond * 2000)

}

func (e *enemy) enemybrpdsmiss(p *playerchar) {
	concyan()
	fmt.Printf("%v«%v» %v %v\n", entprompt(e, 1), e.GetHealth(), e.miss(), p.GetName())
	time.Sleep(time.Millisecond * 2000)

}

func (e *enemy) miss() string {
	x := []string{
		0: "lunges past",
		1: "almost gnaws",
		2: "squeeks and grits its teeth angrily at",
		3: "gets dodged on by",
		4: "runs around aimlessly near",
	}
	return x[rand.Intn(5)]
}

func (e *enemy) attackstring() string {
	x := []string{
		0: "rends the ankles of",
		1: "chews on",
		2: "bites",
		3: "chomps into the leg of",
		4: "slyly pisses off",
	}
	return x[rand.Intn(5)]
}
func (e *enemy) reportdeath(p *playerchar) {
	conred()
	fmt.Printf("\n\n%v dies and cries out: \"%v\"", e.Name, e.Deathcry)
	color.Set(color.FgMagenta, color.Bold, color.BgBlack)
	fmt.Printf("\n\n%v is slain. %v XP awarded.\n", e.Name, e.worthxp)
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
func (e *enemy) Rend(target citentity) int {
	return e.Atk + int(drawrand4()) //returns uint8 to subtract from playerhp
}

func (e *enemy) GetHealth() string {
	x := strconv.Itoa(e.Health)
	return x
}
func (e *enemy) Updatehp(x int) {
	(*e).Health = x
}

func (e *enemy) examine() bool {
	var x bool
	if e.Health > 0 {
		x = true
	}
	return x
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

	(*e).Name = ratnames1[random(0, 8)] + ratnames2[random(0, 8)]
}
