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
	hours := "h"
	minutes := "m"
	seconds := "s"
	milliseconds := "ms"
	var sleep int = 0 // Time to sleep in milliseconds
	var time []string
	var denominator []string
	var previouscharacter string
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i > 0 {
				for _, character := range arg {
					if strings.Contains(string(character), seconds) {
						if strings.Contains(previouscharacter, minutes) == true {
							fmt.Println("Milliseconds were passed")
							if len(denominator) > 0 {
								denominator = denominator[:len(denominator)-1]
								denominator = append(denominator, milliseconds)
							} else {
								denominator = append(denominator, milliseconds)
							}
						}
					}
					if strings.Contains(string(character), hours) || strings.Contains(string(character), minutes) {
						denominator = append(denominator, string(character))
					}
					if timeRegex.MatchString(string(character)) == true { // Match string again numbers and add it to time array
						if timeRegex.MatchString(previouscharacter) {
							time = time[:len(time)-1]
							time = append(time, string(previouscharacter+string(character)))
						} else {
							time = append(time, string(character))
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
	if len(time) > 0 && len(denominator) > 0 {
		for i, x := range denominator {
			if i > len(denominator) {
				break
			}
			switch string(x) {
			case "h":
				// fmt.Printf("Sleeping for %s hours.\n", time[i:count])
				// fmt.Printf("Sleeping for %s hours.\n", getValueFromArray(time, i))
				temp, _ = strconv.Atoi(getValueFromArray(time, i))
				sleep = (sleep) + ((temp * 60) * 60)
			case "m":
				// fmt.Printf("Sleeping for %s minutes.\n", time[i:count])
				// fmt.Printf("Sleeping for %s minutes.\n", getValueFromArray(time, i))
				temp, _ = strconv.Atoi(getValueFromArray(time, i))
				sleep = (sleep) + (temp * 60)
			case "s":
				fmt.Printf("Sleeping for %s seconds.\n", time[i:count])
				temp, _ = strconv.Atoi(getValueFromArray(time, i))
				sleep = (sleep) + (temp)
			case "ms":
				fmt.Printf("Sleeping for %s milliseconds.\n", time[i:count])
				temp, _ = strconv.Atoi(getValueFromArray(time, i))
				sleep = (sleep * 1000) + (temp) // Previously sleep was equal to seconds so convert it to milliseconds and then add milliseconds
			}
			count++
		}
	}
	fmt.Printf("Sleeping for %d seconds", sleep) // This isn't dynamic because it doesn't need to be, as it's just here for testing, and to work out the math of the above case statement
}
