package citadel

import (
	"fmt"

	"github.com/fatih/color"
)

func (l maploc) listitems() {

	conwhite()
	for _, y := range l.items {
		fmt.Printf("%v is on the floor\n", y.name)
	}
	if len(l.items) == 0 {
		fmt.Println("0 items in the area...")
	}
}

func (l maploc) printdescrip() {
	fmt.Print(NAVDESCRIP) //move cursor
	fmt.Printf(" ¬│░░░░▒▒▒▒▒▒▓▓▓▓▓▌DESCRIPTION▐▓▓▓▓▓▒▒▒▒▒▒░░░░│⌐ \n")
	fmt.Print(l.descrip)
}

func (l *maploc) printmap(playerloc uint8) {
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
