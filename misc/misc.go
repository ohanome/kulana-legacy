package misc

import (
	"math"
	"regexp"
	"runtime"
	"time"
)

const Version = "1.0.0"

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Welcome() {
	//now := time.Now().Year()
	//dateParsed, err := time.Parse("2006", "2022")
	//date := dateParsed.Year()
	//str := fmt.Sprintf("%d", date)
	//if now != date {
	//	str += fmt.Sprintf(" - %d", now)
	//}
	//
	//if err != nil {
	//	l.Emergency(err.Error())
	//}
	//l.Info(fmt.Sprintf("==========================================================="))
	//l.Info(fmt.Sprintf("*kulana %s - (c) %s *ohano, All rights reserved", Version, str))
	//l.Info(fmt.Sprintf("Check out the repository: https://github.com/ohanome/kulana"))
	//l.Info(fmt.Sprintf("==========================================================="))
}

func MicroTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func GetNLChar() string {
	nl := "\n"

	if //goland:noinspection ALL
	runtime.GOOS == "windows" {
		nl = "\r\n"
	}

	return nl
}

func RegexMatch(pattern string, str string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}
