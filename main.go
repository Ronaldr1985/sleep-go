package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
	os.Exit(1)
}

func writeHelp() {
	fmt.Printf("A simple countdown program written in Go.\n")
	os.Exit(0)
}

func countdown(minutes, seconds int64) {
	fmt.Printf("Seconds before adding minutes, %d \n", seconds)
	fmt.Printf("Minutes: %d \n", minutes)
	seconds += minutes * 60
	fmt.Printf("Seconds after addition: %d \n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("Time's up\n")
}

func main() {
	timeRegex, _ := regexp.Compile("[0-9]")
	checkhours := "h"
	checkminutes := "m"
	checkseconds := "s"
	checkmilliseconds := "ms"
	var count int = 0
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i > 0 {
				var time []string
				var denominator []string
				var denominatorplace []int
				for i, x := range arg { // Loop through argument as individual strings
					if timeRegex.MatchString(string(x)) == true { // Match string again numbers and add it to time array
						x := string(x)
						time = append(time, x)
					} else { // if it doesn't match regex check if it's a millisecond, second, minute, hour
						if strings.Contains(string(x), checkhours) == true || strings.Contains(string(x), checkminutes) == true || strings.Contains(string(x), checkseconds) || strings.Contains(string(x), checkmilliseconds) {
							denominator = append(denominator, string(x))
							denominatorplace = append(denominatorplace, i)
						}
					}
				}
				fmt.Printf("Which means sleeping for %s%s\n", time[0:denominatorplace[0]], denominator[0:1])
				fmt.Printf("Denominator is in place %d\n", denominatorplace[0:])
				// Do case statement here to check whether sleeping for hours, minutes, seconds or milliseconds
				fmt.Println(time)
				for i, x := range denominator { // Loop through denominators
					switch string(x) {
					case "m":
						fmt.Printf("Sleeping for %s minutes", arg[denominatorplace[count]:denominatorplace[i]])
					case "s":
						// fmt.Println("\nCount variable is equal to:", count)
						// fmt.Println("denominatorplace[i] variable is equal to:", denominatorplace[0])
						// fmt.Println("Time array is equal to:", time)
						// fmt.Println("Variable i is equal to:", i)
						// var seconds string
						// seconds = strings.Join(time[count:denominatorplace[0]], " ")
						fmt.Printf("Sleeping for %s seconds", arg[(denominatorplace[count]+1):denominatorplace[i]])
					}
				}
			}
		}
	} else {
		writeHelp()
	}
}
