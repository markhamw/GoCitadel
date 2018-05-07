$bar = @('░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓██████████▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░')
$citadel = @("▄████████  ▄█      ███        ▄████████ ████████▄     ▄████████  ▄█`n";
"███    ███ ███  ▀█████████▄   ███    ███ ███   ▀███   ███    ███ ███       `n";
"███    █▀  ███▌    ▀███▀▀██   ███    ███ ███    ███   ███    █▀  ███       `n";
"███        ███▌     ███   ▀   ███    ███ ███    ███  ▄███▄▄▄     ███       `n";
"███        ███▌     ███     ▀███████████ ███    ███ ▀▀███▀▀▀     ███       `n";
"███    █▄  ███      ███       ███    ███ ███    ███   ███    █▄  ███       `n";
"███    ███ ███      ███       ███    ███ ███   ▄███   ███    ███ ███▌    ▄ `n";
"████████▀  █▀      ▄████▀     ███    █▀  ████████▀    ██████████ █████▄▄██ `n";)
$skull = ('                      :::!~!!!!!:.',
'                  .xUHWH!! !!?M88WHX:.',
'                .X*#M@$!!  !X!M$$$$$$WWx:.',
'               :!!!!!!?H! :!$!$$$$$$$$$$8X:',
'              !!~  ~:~!! :~!$!#$$$$$$$$$$8X:',
'             :!~::!H!<   ~.U$X!?R$$$$$$$$MM!',
'             ~!~!!!!~~ .:XW$$$U!!?$$$$$$RMM!',
'               !:~~~ .:!M"T#$$$$WX??#MRRMMM!',
'               ~?WuxiW*`   `"#$$$$8!!!!??!!!',
'             :X- M$$$$       `"T#$T~!8$WUXU~',
'            :%`  ~#$$$m:        ~!~ ?$$$$$$',
'          :!`.-   ~T$$$$8xx.  .xWW- ~""##*"',
'.....   -~~:<` !    ~?T#$$@@W@*?$$      /`',
'W$@@M!!! .!~~ !!     .:XUW$W!~ `"~:    :',
'#"~~`.:x%`!!  !H:   !WM$$$$Ti.: .!WUn+!`',
':::~:!!`:X~ .: ?H.!u "$$$B$$$!W:U!T$$M~',
'.~~   :X@!.-~   ?@WTWo("*$$$W$TH$! `',
'Wi.~!X$?!-~    : ?$$$B$Wu("**$RM!',
'$R@i.~~ !     :   ~$$$$$B$$en:``',
'?MXT@Wx.~    :     ~"##*$$$$M~');

$Host.UI.RawUI.CursorPosition = @{ X = $x; Y = $y }

function Set-Cursorpos([int]$x, [int] $y) {
    $Host.UI.RawUI.CursorPosition = New-Object System.Management.Automation.Host.Coordinates $x , $y
    
} 

function Set-World(){
  [console]::WindowHeight = 40
  [console]::WindowWidth = 76
  [console]::BufferHeight = 40
  [console]::BufferWidth = 76
}

function Set-Console-Blue(){
[console]::ForegroundColor = "Blue"
}
function Set-Console-Red(){
  [console]::ForegroundColor = "Red"
  }
function Set-Console-Green(){
    [console]::ForegroundColor = "Green"
    }

function Set-Console-Gray(){
    [console]::ForegroundColor = "Gray"
    }
function Invoke-Title(){
Clear-Host
Set-Cursorpos 0 1

Set-Console-Green
Write-Host $citadel

Set-Console-Blue
Write-Host $bar ,"`n"

Set-Console-Gray
foreach($line in $skull){
  Write-Host $line[-$boof..-1]
  $boof++
 
}

}




#main
Set-World
Invoke-Title