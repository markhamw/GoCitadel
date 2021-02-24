package citadel

import (
	"fmt"
	"time"
)

func drawskull() {
	conred()
	fmt.Println(HOME)
	for _, text := range skull {
		fmt.Println(text)
	}
	fmt.Println(HOME)
}

func skulldrawdown(x ...string) {

	conred()

	fmt.Println(HOME) //Moves cursor to start of 11th line
	fmt.Println(DELETETOENDOFLINE)
	time.Sleep(time.Millisecond * animateddelay)

	pulldownanimation := func(x int) {
		conwhite()
		for count := 0; count < x; count++ {
			fmt.Print(MOVECURSORUPONE)
			fmt.Println(DELETETOENDOFLINE)
			fmt.Printf(ratimage2[0] + "\n")
			fmt.Printf(ratimage2[1] + "\n")
			fmt.Printf(ratimage2[2] + "\n")
			time.Sleep(time.Millisecond * animateddelaySMALL)
		}
	}
	pulldownanimation(30) //28 times/lines

	fmt.Println(HOME)

}
func skullanimateeyesleft() {

	skull[8] = "    `9XXXXXXXXXXXP' `9XX'   \u001b[37;1m█-█\u001b[31;1m██   `98v8P'  █\u001b[37;1m█-█\u001b[31;1m██  `XXP' `9XXXXXXXXXXXP'"
	drawskull()
	time.Sleep(time.Millisecond * 300)
}
func skullanimateeyesright() {

	skull[8] = "    `9XXXXXXXXXXXP' `9XX'   ██\u001b[37;1m█-█\u001b[31;1m   `98v8P'  ███\u001b[37;1m█-█\u001b[31;1m  `XXP' `9XXXXXXXXXXXP'"
	drawskull()
	time.Sleep(time.Millisecond * 250)
}
func titleintro() {
	clear()
	drawtitle()
	time.Sleep(time.Millisecond * 2000)
	drawskull()
	time.Sleep(time.Millisecond * 1000)
	skullanimateeyesleft()
	skullanimateeyesright()
	skullanimateeyesleft()
	skullanimatefinale()
	skulldrawdown()
	time.Sleep(time.Millisecond * 100)

}
func skullanimatefinale() {

	skull[16] = "                              XXXX X.`▓'.X XXXX"
	skull[17] = "                              XP^X'`▓   ▓'`X^XX"
	skull[18] = "                              X. ▓  `   '  ▓ )X"
	skull[19] = "                              `▓  `       '  ▓'"
	skull[20] = "                                `             '"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[16] = "                              XXXX X.`\u001b[33;1m█\u001b[31;1m'.X XXXX"
	skull[17] = "                              XP^X'`▓ ▓ ▓'`X^XX"
	skull[18] = "                              X. ▓  `   '  ▓ )X"
	skull[19] = "                              `▓  `       '  ▓'"
	skull[20] = "                                `             '"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'   █\u001b[37;1m█-█\u001b[31;1m█   `98v8P'  █\u001b[37;1m█-█\u001b[31;1m█   `XXP' `9XXXXXXXXXXXP'"
	skull[16] = "                              XXXX X.` '.X XXXX"
	skull[17] = "                              XP^X'`\u001b[33;1m▓▓▓▓▓\u001b[31;1m'`X^XX"
	skull[18] = "                              X.\u001b[33;1m█▓▓▓\u001b[31;1m▓ ▓ ▓\u001b[33;1m▓▓▓█\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓` `   '  '▓ ▓'"
	skull[20] = "                                `             '"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█▓▓▓█\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m███▓█████▓███\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓`▓`▓ ▓'▓▓'▓ ▓'"
	skull[20] = "                                `             '"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██▓▓█████▓▓█\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓`\u001b[33;1m▓█████▓\u001b[31;1m'▓ ▓'"
	skull[20] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[21] = "                                  XX XX XX"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              `▓ ▓\u001b[33;1m█  ███  █\u001b[31;1m'▓ ▓'"
	skull[20] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[21] = "                                  XX XX XX"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[21] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[22] = "                                  XX XX XX"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[21] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[22] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[23] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[24] = "                                  XX XX XX"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
	skull[8] = "    `9XXXXXXXXXXXP' `9XX'  █\u001b[37;1m██-██\u001b[31;1m█  `98v8P' █\u001b[37;1m██-██\u001b[31;1m█  `XXP' `9XXXXXXXXXXXP'"
	skull[16] = "                              XXXX X.║║║.X XXXX"
	skull[17] = "                              XP^X║║\u001b[33;1m█████\u001b[31;1m║║X^XX"
	skull[18] = "                              X.\u001b[33;1m██  █████  ██\u001b[31;1m)X"
	skull[19] = "                              X.\u001b[33;1m███ █████ ███\u001b[31;1m)X"
	skull[20] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[21] = "                              X.\u001b[33;1m█████████████\u001b[31;1m)X"
	skull[22] = "                              `▓ ▓\u001b[33;1m██ ███ ██\u001b[31;1m'▓ ▓'"
	skull[23] = "                                XX`▓ ▓``▓ ▓`XX'"
	skull[24] = "                                  XX XX XX"
	drawskull()
	time.Sleep(time.Millisecond * animateddelaySMALL)
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
	drawskull()

	time.Sleep(time.Millisecond * animateddelay)
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
	8:  "    `9XXXXXXXXXXXP' `9XX'   ██\u001b[37;1m█-█\u001b[31;1m   `98v8P'  ███\u001b[37;1m█-█\u001b[31;1m  `XXP' `9XXXXXXXXXXXP'",
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

var ratimage4 = []string{
	0: ",,==.",
	1: "//   `",
	2: "||      ,--~~~~-._ _(\\--,_",
	3: "\\\\._,-~/  \\      '     *  `o'",
	4: " `---~\\( _/,___( /_/`---~~/",
	5: "       ``==-    `==-,",
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
	3: "  )\\ ))__ _\\    |___)",
	4: " nn-nn`  `nn---'",
}
var ratimage1 = []string{
	0: "(\\,;,/)",
	1: " (o o)\\//,",
	2: "  \\ /     \\,",
	3: "  `+'(  (   \\   )",
	4: "     //  \\  |_./",
	5: "   '~' '~----'",
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
	0: "You're a mercenary adventurer...living in the castle...working for the Empress,\n",
	1: "and apprentice to Exarch Gurlows. Your commission is to neutralize feral finks in the\n",
	2: "cellar & you're approaching the cellar now. Rumors dictate a king rat holds the\n",
	3: "sub-basement. Aggressive ones are causing trouble with kitchen staff. Gurlows\n",
	4: "expects raticide. You arrive at the Cellar, weapon ready, hungover and wary.\n",
	5: "                      ▒▓▌PRESS ANY KEY TO CONTINUE▐▓▒",
}

var storyintrotext1 = []string{
	0: "Slowly.....you begin to regain consciousness.........",
	1: "Visions of the Beast fade as your nightmare subsides....",
	2: "You wake up in rented quarters inside the Citadel.",
	3: ".................",
	4: "You struggle to your feet.",
	5: "The Beast be damned. You're too hungover to remember your name.",
	6: "What is it again!?...",
}

/* var titletext = []string{
	0:  ` ▄████████  ▄█      ███        ▄████████ ████████▄     ▄████████  ▄█                ███      ▄██████▄   ▄█     █▄     ▄████████    ▄████████ `,
	1:  ` ███    ███ ███  ▀█████████▄   ███    ███ ███   ▀███   ███    ███ ███            ▀█████████▄ ███    ███ ███     ███   ███    ███   ███    ███ `,
	2:  ` ███    █▀  ███▌    ▀███▀▀██   ███ ▄  ███ ███    ███   ███    █▀  ███        ██     ▀███▀▀██ ███  ▄ ███ ███     ███   ███    █▀    ███  ▄ ███ `,
	3:  ` ███        ███▌     ███   ▀   ███    ███ ███    ███  ▄███▄▄▄     ███                ███   ▀ ███    ███ ███     ███  ▄███▄▄▄      ▄███▄▄▄▄██▀ `,
	4:  ` ███░░░░░░░░████░░░░░███░░░░░████████████░███░░░░███░█████▀▀▀░░░░░███░░░░░░░░░░░░░░░░███░░░░░███░░░░███░███░░░░░███░████████░░░░░░█████▀▀▀▀▀   `,
	5:  ` ███    █▄  ███      ███       ███    ███ ███    ███   ███    █▄  ███        ██      ███     ███    ███ ███     ███   ███    █▄  ▀███████████ `,
	6:  ` ███    ███ ███      ███       ███    ███ ███   ▄███   ███    ███ ███▌    ▄          ███     ███    ███ ███ ▄█▄ ███   ███    ███   ███    ███ `,
	7:  ` ███    █▀  █▀      ▄████▀     ███    █▀  ████████▀    ██████████ █████▄▄██         ▄████▀    ▀██████▀   ▀███▀███▀    ██████████   ███    ███ `,
	8:  ` ████████              ▄                                           ▀                                                                ███    ███ `,
	9:  ` ██████`,
	10: `░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██████████████████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░`,
} */
var titletext2 = []string{
	0:  ` ▄████████  ▄█      ███        ▄████████ ████████▄     ▄████████  ▄█`,
	1:  ` ███    ███ ███  ▀█████████▄   ███    ███ ███   ▀███   ███    ███ ███`,
	2:  ` ███    █▀  ███▌    ▀███▀▀██   ███ ▄  ███ ███    ███   ███    █▀  ███`,
	3:  ` ███        ███▌     ███   ▀   ███    ███ ███    ███  ▄███▄▄▄     ███`,
	4:  ` ███░░░░░░░░████░░░░░███░░░░░████████████░███░░░░███░█████▀▀▀░░░░░███`,
	5:  ` ███    █▄  ███      ███       ███    ███ ███    ███   ███    █▄  ███`,
	6:  ` ███    ███ ███      ███       ███    ███ ███   ▄███   ███    ███ ███▌    ▄`,
	7:  ` ███    █▀  █▀      ▄████▀     ███    █▀  ████████▀    ██████████ █████▄▄██`,
	8:  ` ████████              ▄                                           ▀       `,
	9:  ` ██████`,
	10: `░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓█████████████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒░░░░░░░░`,
}

var castle2 = `                                 o
                             .-""|
                             |-""|
                                 |   ._--+.
                                .|-""      '.
                               +:'           '.
                               | '.        _.-'|
                               |   +    .-"   J
            _.+        .....'.'|    '.-"      |
       _.-""   '.   ..'88888888|     +       J''..
    +:"          '.'88888888888;-+.  |    _+.|8888:
    | \         _.-+888888888_." _.F F +:'   '.8888'....
     L \   _.-""   |8888_.-"  _." J J J  '.    +88888888:
     |  '+"        |_.-"  _.-"    | | |    +    '.888888'._''.
   .'8L  L         J  _.-"        | | |     '.    '.88_.-"    '. 
  :888|  |         J-"            F F F       '.  _.-"          '.
 :88888L  L     _+  L            J J J          '|.               '; 
:888888J  |  +-"  \ L          _.+.|.+.          F '.          _.-" F
:8888888|  L L\    \|      _.-"    '   '.       J    '.     .-"    |
:8888888.L | | \    ', _.-"              '.     |      "..-"      J'.
:888888: |  L L '.    \     _.-+.          '.   :+-.     |        F88'.
:888888:  L | |   \    ;.-""     '.          :-"    ":+ J        |88888:
:888888:  |  L L   +:""            '.    _.-"     .-" | |       J:888888:
:888888:   L | |   J \               '.-'     _.-'   J J        F :888888:
 :88888:    \ L L   L \             _.-+  _.-'       | |       |   :888888:
 :888888:    \| |   |  '.       _.-"   |-"          J J       J     :888888:
 :888888'.    +'\   J    \  _.-"       F    ,-T"\  | |     .-'      :888888:
  :888888 '.     \   L    +"          J    /  | J  J J  .-'        .'888888:
   :8888888 :     \  |    |           |    F  '.|.-'+|-'         .' 8888888:
    :8888888 :     \ J    |           F   J     '...           .' 888888888:
     :8888888 :     \ L   |          J    |      \88'.''.''''.' 8888888888:
      :8888888 :     \|   |          |  .-'\      \8888888888888888888888:
       :8888888 '.    J   |          F-'  .'\      \8888888888888888888.'
        :88888888 :    L  |         J     : 8\      \8888888888888888.'
         :88888888 :   |  |        .+  ...' 88\      \8888888888.''.'
          :88888888 :  J  |     .-'  .'    8888\      \'.'''.'.' 
           :88888888 :  \ |  .-'   .' 888888888.\    _-'     
           :888888888 :  \|-'     .' 888888888.' \_-"
            '.88888888'..         : 8888888.' 
              :88888888  ''''.''.' 88888888:         
              :8888888888888888888888888888:
               :88888888888888888888888888:
                :888888888888888888888888:
                 ''.8888888888888...'.'''
                     '''''......''`
