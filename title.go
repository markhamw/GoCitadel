package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func drawskull() {
	color.Set(color.FgRed, color.Bold, color.BgBlack)
	for _, text := range skull {
		fmt.Println("\t\t\t\t", text)
	}

}
func skullrefresh() {
	for _, text := range skull {
		fmt.Println("\t\t\t\t", text)
	}
}
func skulldrawdown() {
	color.Set(color.FgRed, color.Bold, color.BgBlack)
	titlebar1 := strings.Repeat("_", len(titletext[1])/2) //Prints Red Bar below erase effect f
	fmt.Println(HOME)                                     //Moves cursor to start of 11th line
	fmt.Println(DELETETOENDOFLINE)
	time.Sleep(time.Millisecond * ANIMATEDELAY)
	pulldownanimation := func(x int) {
		for count := 0; count < x; count++ {
			fmt.Print("\033[1A") //move cursor up 1 line
			fmt.Print("\033[K")  //Delete everything from the cursor to the end of the line.
			fmt.Println(strings.Repeat(" ", 180))
			fmt.Println(titlebar1) //Prints Red Bar below erase effect f
			time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
		}
	}
	pulldownanimation(30) //28 times/lines

	fmt.Println(HOME)
}
func skullanimateeyesleft() {
	clear()
	drawtitle()
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'   \u001b[37;1m█-█\u001b[31;1m██   `98v8P'  █\u001b[37;1m█-█\u001b[31;1m██  `XXP' `9XXXXXXXXXXXP'"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
}
func skullanimateeyesright() {
	clear()
	drawtitle()
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'   ██\u001b[37;1m█-█\u001b[31;1m   `98v8P'  ███\u001b[37;1m█-█\u001b[31;1m  `XXP' `9XXXXXXXXXXXP'"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
}
func titleintro() {
	clear()
	fmt.Println(HIDECURSOR)

	drawtitle()
	drawskull()
	skullanimateeyesleft()
	skullanimateeyesright()
	skullanimateeyesleft()
	skullanimatefinale()
	skulldrawdown()
	time.Sleep(time.Millisecond * 100)

}
func skullanimatefinale() {
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.`▓'.X XXXX"
	skull[17] = "                              XP^X'`▓   ▓'`X^XX"
	skull[18] = "                              X. ▓  `   '  ▓ )X"
	skull[19] = "                              `▓  `       '  ▓'"
	skull[20] = "                                `             '"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.`\u001b[33;1m█\u001b[31;1m'.X XXXX"
	skull[17] = "                              XP^X'`▓ ▓ ▓'`X^XX"
	skull[18] = "                              X. ▓  `   '  ▓ )X"
	skull[19] = "                              `▓  `       '  ▓'"
	skull[20] = "                                `             '"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'   █\u001b[37;1m█-█\u001b[31;1m█   `98v8P'  █\u001b[37;1m█-█\u001b[31;1m█  `XXP' `9XXXXXXXXXXXP'"
	skull[16] = "                              XXXX X.` '.X XXXX"
	skull[17] = "                              XP^X'`\u001b[33;1m▓▓▓▓▓\u001b[31;1m'`X^XX"
	skull[18] = "                              X.\u001b[33;1m█▓▓▓\u001b[31;1m▓ ▓ ▓\u001b[33;1m▓▓▓█\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓` `   '  '▓ ▓'"
	skull[20] = "                                `             '"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█▓▓▓█\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m███▓█████▓███\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓`▓`▓ ▓'▓▓'▓ ▓'"
	skull[20] = "                                `             '"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██▓▓█████▓▓█\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓`\u001b[33;1m▓█████▓\u001b[31;1m'▓ ▓'"
	skull[20] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[21] = "                                  XX XX XX"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓\u001b[33;1m█  ███  █\u001b[31;1m'▓ ▓'"
	skull[20] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[21] = "                                  XX XX XX"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[21] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[22] = "                                  XX XX XX"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[21] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[22] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[23] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[24] = "                                  XX XX XX"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'  █\u001b[37;1m██-██\u001b[31;1m█  `98v8P' █\u001b[37;1m██-██\u001b[31;1m█ `XXP' `9XXXXXXXXXXXP'"
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[21] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[22] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[23] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[24] = "                                  XX XX XX"
	skullrefresh()
	time.Sleep(time.Millisecond * ANIMATEDELAYSMALL)
	clear()
	drawtitle()
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'  █\u001b[37;1m██-██\u001b[31;1m█  `98v8P' █\u001b[37;1m██-██\u001b[31;1m█ `XXP' `9XXXXXXXXXXXP'"
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[21] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[22] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[23] = "                                XX`▓ \u001b[33;1m███\u001b[31;1m ▓`XX'"
	skull[24] = "                                  XX▓`▓`▓X"
	skull[25] = "                                  XX XX XX"
	skullrefresh()

	time.Sleep(time.Millisecond * ANIMATEDELAY)
}

var skull = [26]string{
	0:  "         .                                                      .",
	1:  "        .n                   .                 .                  n.",
	2:  "  .   .dP                  dP                   9b                 9b.    .",
	3:  "  4    qXb         .       dX                     Xb       .        dXp     t",
	4:  "dX.    9Xb      .dXb    __                         __    dXb.     dXP     .Xb",
	5:  "9XXb._       _.dXXXXb dXXXXbo.                 .odXXXXb dXXXXb._       _.dXXP",
	6:  " 9XXXXXXXXXXXXXXXXXXXVXXXXXXXXOo.           .oOXXXXXXXXVXXXXXXXXXXXXXXXXXXXP",
	7:  "  `9XXXXXXXXXXXXXXXXXXXXX'~   ~`OOO8b   d8OOO'~   ~`XXXXXXXXXXXXXXXXXXXXXP'",
	8:  "    `9XXXXXXXXXXXP' `9XX'   ██\u001b[37;1m█-█\u001b[31;1m   `98v8P'  ███\u001b[37;1m█-█\u001b[31;1m `XXP' `9XXXXXXXXXXXP'",
	9:  "        ~~~~~~~       9X.          .db|db.          .XP       ~~~~~~~",
	10: "                        )b.  .dbo.dP'`v'`9b.odb.  .dX(",
	11: "                      ,dXXXXXXXXXXXb     dXXXXXXXXXXXb.",
	12: "                     dXXXXXXXXXXXP'   .   `9XXXXXXXXXXXb",
	13: "                    dXXXXXXXXXXXXb   d|b   dXXXXXXXXXXXXb",
	14: "                    9XXb'   `XXXXXb.dX|Xb.dXXXXX'   `dXXP",
	15: "                     `'      9XXXXXX(   )XXXXXXP      `'",
	16: "                              XXXX X.`v'.X XXXX",
	17: "                              XP^X'`b   d'`X^XX",
	18: "                              X. 9  `   '  P )X",
	19: "                              `b  `       '  d'",
	20: "                                `             '",
}

func tutintro(player *Playerchar) {
	drawplayertitleframe(player)
	fmt.Println(castle)
	readblock(castletext)
}

var ratimage1 = []string{
	0: ",,==.",
	1: "//    `",
	2: "||      ,--~~~~-._ _(\\--,_",
	3: "\\._,-~   \\      '    *  `o'",
	4: " `---~\\( _/,___( /_/`---~~/",
	5: "	   ``==-    `==-,",
}
var ratimage2 = []string{
	0: "         ____()()",
	1: "        /      @@",
	2: "`~~~~~\\_;m__m._>o",
}
var ratimage3 = []string{
	0: "(q\\_/p)",
	1: " /. .\\.-\"\"\"\"-.      ___",
	2: "=\\_t_/=    /  `\\   (",
	3: " )\\ ))__ _\\    |___)",
	4: "nn-nn`  `nn---'",
}
var castle = `
                       _    __    __    __    __    _                          
                      | |__|__|__|__|__|__|__|__|__|_|                         
 __    __    __       |_|___|___|___|___|___|___|___||       __    __    __   
|__|  |__|  |__|      |___|___|___|___|___|___|___|__|      |__|  |__|  |__|   
|__|__|__|__|__|       \____________________________/       |__|__|__|__|__|   
|_|___|___|___||        |_|___|___|___|___|___|___||        |_|___|___|___||
|___|___|___|__|        |___|___|___|___|___|___|__|        |___|___|___|__|   
 \_|__|__|___|/          \________________________/          \_|__|__|__|_/    
  \__|____|__/            |___|___|___|___|___|__|            \__|__|__|_/     
   |||_|_|_||             |_|___|C A S T L E|__|_|             |_|_|_|_||      
   ||_|_|||_|__    __    _| _  __ |I O T A|__  _ |_    __    __||_|_|_|_|      
   |_|_|_|_||__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|__|_|_|_|_||      
   ||_|||_|||___|___|___|___|___|___|___|___|___|___|___|___|__||_|_|_|_|      
   |_|_|_|_||_|___|___|___|___|___|___|___|___|___|___|___|___||_|_|_|_||      
   ||_|_|_|_|___|___|___|___|___|___|___|___|___|___|___|___|__||_|_|_|_|      
   |_|||_|_||_|___|___|___|___|___|___|___|___|___|___|___|___||_|_|_|_||      
   ||_|_|_|_|___|___|___|___|___|_/| | | \__|___|___|___|___|__||_|_|_|_|      
   |_|_|_|_||_|___|___|___|___|__/ | | | |\___|___|___|___|___||_|_|_|_||      
   ||_|_|_|||___|___|___|___|___|| | | | | |____|___|___|___|__||_|_|_|_|      
   |_|_|_|_||_|___|___|___|___|_|| | | | | |__|___|___|___|___||_|_|_|_||      
  /___|___|__\__|___|___|___|___|| | | | | |____|___|___|___|_/_|___|__|_\     
 |_|_|_|_|_|_||___|___|___|___|_|| | | | | |__|___|___|___|__|_|__|__|__|_|    
 ||_|_|_|_|_|_|_|___|___|___|___||_|_|_|_|_|____|___|___|____|___|__|__|__|`

var castletext = []string{
	0: "\nYour're an adventurer living in Castle Iota, currently working as a mercenary for the emporer.\n",
	1: "Your last assignment was to clear any rats from the cellar & you are approaching the cellar now.\n",
	2: "Out of control cellar vermin are causing trouble with kitchen staff, again. Master Gurlows expects raticide.\n",
	3: "You ready your flail and prepare to enter the Cellar.\nPRESS ANY KEY TO CONTINUE",
}

var storyintrotext1 = []string{
	0: "\tSlowly..you regain consciousness.\n\n",
	1: "\t............\n\n",
	2: "\tYou're waking up in your meager rented quarters in Castle Iota within the Citadel walls.\n",
	3: "\n",
	4: "\tLast nights celebration has left you weak and hungover! You warily climb to your feet.\n",
	5: "\n",
	6: "\tYou drank too much, again. As you begin to stand on two feet you have trouble remembering your own name.\n",
	7: "\n",
	8: "\tWhat is it again?...\n\n",
	9: "\tEnter your name: ",
}
var titletext = []string{
	0:  "  ▄████████  ▄█      ███        ▄████████ ████████▄     ▄████████  ▄█                ███      ▄██████▄   ▄█     █▄     ▄████████    ▄████████ ",
	1:  " ███    ███ ███  ▀█████████▄   ███    ███ ███   ▀███   ███    ███ ███            ▀█████████▄ ███    ███ ███     ███   ███    ███   ███    ███ ",
	2:  " ███    █▀  ███▌    ▀███▀▀██   ███ ▄  ███ ███    ███   ███    █▀  ███        ██     ▀███▀▀██ ███  ▄ ███ ███     ███   ███    █▀    ███  ▄ ███ ",
	3:  " ███        ███▌     ███   ▀   ███    ███ ███    ███  ▄███▄▄▄     ███                ███   ▀ ███    ███ ███     ███  ▄███▄▄▄      ▄███▄▄▄▄██▀ ",
	4:  " ███░░░░░░░░████░░░░░███░░░░░████████████░███░░░░███░█████▀▀▀░░░░░███░░░░░░░░░░░░░░░░███░░░░░███░░░░███░███░░░░░███░████████░░░░░░█████▀▀▀▀▀   ",
	5:  " ███    █▄  ███      ███       ███    ███ ███    ███   ███    █▄  ███        ██      ███     ███    ███ ███     ███   ███    █▄  ▀███████████ ",
	6:  " ███    ███ ███      ███       ███    ███ ███   ▄███   ███    ███ ███▌    ▄          ███     ███    ███ ███ ▄█▄ ███   ███    ███   ███    ███ ",
	7:  " ███    █▀  █▀      ▄████▀     ███    █▀  ████████▀    ██████████ █████▄▄██         ▄████▀    ▀██████▀   ▀███▀███▀    ██████████   ███    ███ ",
	8:  " ████████              ▄                                           ▀                                                                ███    ███ ",
	9:  " ██████",
	10: "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-",
}
