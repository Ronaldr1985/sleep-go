package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func after(value string, match string) string {
	pos := strings.LastIndex(value, match)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(match)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func write_help(exit_value int) {
	if exit_value == 0 {
		fmt.Printf("Usage: sleep-go NUMBER[SUFFIX]..\n\tor: sleep-go OPTION\nPause for NUMBER of seconds.  SUFFIX may be 's' for seconds (the default), 'm' for minutes, 'h' for hours or 'd' for days.  Given two or more arguments pause for the sum of all the values.\n\n\t--help     display this help and exit\n\t--version  output version information and exit\n\n")
	} else {
		fmt.Fprintf(os.Stderr, "sleep-go: invalid option -- '%s'\nTry 'sleep-go --help' for more information\n", string(os.Args[1][1]))
	}
	os.Exit(exit_value)
}

func write_version() {
	fmt.Printf("%s 0.1\nLicense BSD-2-Clause\n\nWritten by Ronald 1985.\n", after(os.Args[0], "/"))
}

func main() {
	if len(os.Args) > 1 {
		var err error
		var current_number, next_character string
		var seconds, tmp_number float64
		if os.Args[1][0] == '-' {
			if len(os.Args[1]) == 1 {
				fmt.Fprintf(os.Stderr, "sleep-go: invalid option -- '%s'\nTry 'sleep-go --help' for more information\n", string(os.Args[1][0]))
				os.Exit(1)
			} else if unicode.IsLetter(rune(os.Args[1][1])) {
				write_help(1)
			} else if os.Args[1][1] == '-' {
				if strings.Compare(os.Args[1][2:], "help") == 0 {
					write_help(0)
				} else if strings.Compare(os.Args[1][2:], "version") == 0 {
					write_version()
				} else {
					write_help(1)
				}
			}
			os.Exit(1)
		}
		for i, argument := range os.Args {
			if i > 0 {
				for i, ch := range argument {
					if unicode.IsNumber(ch) {
						current_number += string(ch)
					} else if unicode.IsLetter(ch) && i+1 == 1 {
						fmt.Println("I'm here")
						write_help(1)
					}
					switch ch {
					case 'd':
						tmp_number, err = strconv.ParseFloat(current_number, 8)
						if err != nil {
							log.Println("error: ", err)
						}
						seconds += (tmp_number * 24 * 60 * 60)
						current_number = ""
					case 'h':
						tmp_number, err = strconv.ParseFloat(current_number, 8)
						if err != nil {
							log.Fatal("error ", err)
						}
						seconds += (tmp_number * 60 * 60)
						current_number = ""
					case 'm':
						tmp_number, err = strconv.ParseFloat(current_number, 8)
						if err != nil {
							log.Fatal("error ", err)
						}
						seconds += tmp_number * 60
						current_number = ""
					case 's':
						tmp_number, err = strconv.ParseFloat(current_number, 8)
						if err != nil {
							log.Fatal("error ", err)
						}
						seconds += tmp_number
						current_number = ""
					case '.':
						current_number += "."
					default:
						if _, err := strconv.Atoi(next_character); err != nil && i == len(argument) {
							tmp_number, err = strconv.ParseFloat(current_number, 8)
							if err != nil {
								log.Fatal("error ", err)
							}
							seconds += tmp_number
							current_number = ""
						} else if unicode.IsLetter(ch) {
							write_help(1)
						}
					}
				}
			}
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
