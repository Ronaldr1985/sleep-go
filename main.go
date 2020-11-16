package main

import (
	"fmt"
	"os"
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
	// timeRegex, _ := regexp.Compile("[0-9]")
	hours := "h"
	minutes := "m"
	seconds := "s"
	milliseconds := "ms"
	// var time []string
	var denominator []string
	// var denominatorplace []int
	// var count int = 0
	var previouscharacter string
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i > 0 {
				var count int = 0
				for _, character := range arg {
					if strings.Contains(string(character), seconds) {
						if strings.Contains(previouscharacter, minutes) == true {
							fmt.Println("Milliseconds were passed")
							count++
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
					previouscharacter = string(character)
				}
				fmt.Println(denominator)
			}
		}
	} else {
		writeHelp()
	}
}
