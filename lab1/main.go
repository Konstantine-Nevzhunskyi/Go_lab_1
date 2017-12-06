package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/pkg/profile"
)

var result []string

func pick(mask, str, foundStr string) {
	if len(str) == 0 || len(mask) == 0 {
		if len(mask) == 0 || (len(mask) == 1 && mask[0] == '*') {
			result = append(result, foundStr) //add foundStr to RESULT
			//fmt.Println(foundStr)
		}
	} else {
		newStr := strings.Replace(str, string(str[0]), "", 1)
		newMask := strings.Replace(mask, string(mask[0]), "", 1)
		if mask[0] == '*' {
			pick(mask, newStr, foundStr+string(str[0])) //нашли часть звездочки
			pick(newMask, str, foundStr)                //нашли конец звездочки
		} else {
			if mask[0] == str[0] {
				pick(newMask, newStr, foundStr+string(str[0]))
			}
		}
	}
}

func start(mask, str string) {
	l := len(str)
	for i := 0; i < l; i++ {
		pick(mask, str, "")
		str = strings.Replace(str, string(str[0]), "", 1)
	}
}

func main() {

	defer profile.Start().Stop()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your string (will be read from the file if empty): ")
	scanner.Scan()
	str := scanner.Text()

	if str == "" {
		bs, err := ioutil.ReadFile("standartStr.txt")
		if err != nil {
			return
		}
		str = string(bs)
	}

	fmt.Print("Enter your mask: ")
	scanner.Scan()
	mask := scanner.Text()

	startTime := time.Now()
	start(mask, str)
	endTime := time.Now()

	exeTime := endTime.Sub(startTime)

	fmt.Printf("Done. %d results.", len(result))
	fmt.Printf("Time: %fs\n", exeTime.Seconds())

	//*
	fmt.Print("\n'y' to print: ")
	scanner.Scan()
	str = scanner.Text()
	if str == "y" {
		fmt.Print(result)
	} //*/
}
