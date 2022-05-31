package misc

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"time"
)

func MicroTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Die(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func Ternary(condition bool, trueValue any, falseValue any) any {
	if condition {
		return trueValue
	}

	return falseValue
}

func GetNLChar() string {
	nl := "\n"

	if //goland:noinspection ALL
	runtime.GOOS == "windows" {
		nl = "\r\n"
	}

	return nl
}
