package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
	os.Exit(1)
}

func getValueFromArray(arr []string, position int) string {
	return arr[position]
}

func sleepMilliseconds(milliseconds int64) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func sleepSeconds(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func writeHelp() {
	fmt.Printf("A simple countdown program written in Go.\n")
	os.Exit(0)
}

func main() {
	timeRegex, _ := regexp.Compile("[0-9]")
	hours := "h"
	minutes := "m"
	seconds := "s"
	milliseconds := "ms"
	var sleep int = 0 // Time to sleep in milliseconds
	var amount []string
	var denominator []string
	var previouscharacter string
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i > 0 {
				for _, character := range arg {
					if strings.Contains(string(character), seconds) {
						if strings.Contains(previouscharacter, minutes) == true {
							// fmt.Println("Milliseconds were passed")
							if len(denominator) > 0 {
								denominator = denominator[:len(denominator)-1]
								denominator = append(denominator, milliseconds)
							} else {
								denominator = append(denominator, milliseconds)
							}
						} else {
							denominator = append(denominator, string(character))
						}
					}
					if strings.Contains(string(character), hours) || strings.Contains(string(character), minutes) {
						denominator = append(denominator, string(character))
					}
					if timeRegex.MatchString(string(character)) == true { // Match string again numbers and add it to amountarray
						if timeRegex.MatchString(previouscharacter) {
							amount = amount[:len(amount)-1]
							amount = append(amount, string(previouscharacter+string(character)))
						} else {
							amount = append(amount, string(character))
						}
					}
					previouscharacter = string(character)
				}
			}
		}
	} else {
		writeHelp()
	}
	var count int = 1
	var temp int
	var msTrue bool
	if len(amount) > 0 && len(denominator) > 0 {
		for i, x := range denominator {
			if i > len(denominator) {
				break
			}
			switch string(x) {
			case "h":
				temp, _ = strconv.Atoi(getValueFromArray(amount, i))
				sleep = (sleep) + ((temp * 60) * 60)
			case "m":
				temp, _ = strconv.Atoi(getValueFromArray(amount, i))
				sleep = (sleep) + (temp * 60)
			case "s":
				temp, _ = strconv.Atoi(getValueFromArray(amount, i))
				sleep = (sleep) + (temp)
			case "ms": // The way that milliseconds works if it isn't the last one it won't probably convert everything to milliesconds
				temp, _ = strconv.Atoi(getValueFromArray(amount, i))
				sleep = (sleep * 1000) + (temp) // Previously sleep was equal to seconds so convert it to milliseconds and then add milliseconds
				msTrue = true
			}
			count++
		}
	}
	if msTrue == true {
		fmt.Printf("Sleeping for %d milliseconds", sleep)
		sleepMilliseconds(int64(sleep))
	} else {
		fmt.Printf("Sleeping for %d seconds", sleep)
		sleepSeconds(int64(sleep))
	}
}
