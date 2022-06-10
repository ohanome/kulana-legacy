package format

import (
	"fmt"
	"strings"
)

const Reset = "\033[0m"
const Bold = "\033[1m"
const Underline = "\033[4m"
const Inverse = "\033[7m"
const BoldOff = "\033[21m"
const UnderlineOff = "\033[24m"
const InverseOff = "\033[27m"

const Black = "\033[30m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Purple = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"

const BGBlack = "\033[40m"
const BGRed = "\033[41m"
const BGGreen = "\033[42m"
const BGYellow = "\033[43m"
const BGBlue = "\033[44m"
const BGPurple = "\033[45m"
const BGCyan = "\033[46m"
const BGWhite = "\033[47m"

func PrintErr(err error) {
	fmt.Println(Red + Bold + "[ERROR] " + Reset + Red + err.Error() + Reset)
}

func Rainbow(str string) {
	spl := strings.Split(str, "")
	cs := []string{
		Red,
		Yellow,
		Green,
		Cyan,
		Blue,
		Purple,
	}

	s := ""
	ci := 0
	for _, sp := range spl {
		fmt.Printf("sp: '%s'\n", sp)
		if sp != " " {
			if ci == len(cs) {
				ci = 0
			}
			s += cs[ci] + sp
			ci++
		} else {
			s += sp
		}
	}
	s += Reset

	fmt.Println(s)
}
