package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

const (
	animateddelay          = 100
	animateddelaySMALL     = 50
	MINHEALTH              = 0
	BASEDAM                = 6
	ZEROHOME               = "\033[0;0H"
	playerbaruiboxname     = "\033[02;02H"
	playerbaruiboxnamedone = "\033[04;01H"

	playerbaruiboxloc1    = "\033[01;39H"
	playerbaruiboxloc2    = "\033[02;39H"
	playerbaruiboxloc3    = "\033[03;39H"
	playerbaruiboxloctext = "\033[02;40H"
	HOME                  = "\033[12;0H"
	HIDECURSOR            = "\033[?25l"
	PLAYERINFOHOME        = "\033[11;0H"
	NAVAREA               = "\033[09;16H"
	NAVLOC                = "\033[04;0H"
	COMPASSDRAW1          = "\033[4;26H"
	COMPASSDRAW2          = "\033[5;24H"
	COMPASSDRAW3          = "\033[6;26H" //
	NAVDESCRIP            = "\033[11;0H" //
	NAVTRAVEL             = "\033[24;0H" //
	ENEMIESLISTED         = "\033[19;0H" //
	HEALTHBAR             = "\033[4;40H" //
	HEALTHBARFILL         = "\033[4;41H" //
	NAVTRAVELOUTPUT       = "\033[26;0H" //
	ATTACKSTART           = "\033[6;0H"  //
	DELETETOENDOFLINE     = "\033[K"     //
	SAVECURSOR            = "\033[s"     //
	MOVECURSORUPONE       = "\033[3A"
	RETURNTOSAVEDCURSOR   = "\033[u"  //
	RESETCOLORS           = "\033[0m" //
)

func promptofshame() {
	clear()
	fmt.Println(ZEROHOME)
	conyellow()
	fmt.Println("RECOMMENDED:\nSet FONT SIZE to 18+")
	fmt.Println("Expand this window a good bit. Go on.")
	fmt.Println("now press a key to continue")
	getcharpause()

}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print(HIDECURSOR)

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
func conyellow() {
	color.Set(color.FgYellow, color.Bold, color.BgBlack)
}
func congreen() {
	color.Set(color.FgGreen, color.Bold, color.BgBlack)
}
func conblue() {
	color.Set(color.FgBlue, color.Bold, color.BgBlack)
}
func concyan() {
	color.Set(color.FgCyan, color.Bold, color.BgBlack)
}
