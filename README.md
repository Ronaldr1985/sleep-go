# Countdown Timer

A simple countdown timer written in Go, it was built to be an alternative to the sleep command on Linux.

## Usage

First things first compile the application: 

    go build

 And then run it passing it the time you want to sleep for. For example:

So to sleep for 5 minutes and 30 seconds

    ./countdown-timer 5m30s 

Please note this application is still very much in **testing**, and is still being developed.

## Bugs

If you don't pass milliseconds to the program first then it unforunately won't work and you will get undesired results. **Fixed on 09/01/2021**