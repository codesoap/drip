package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage:
drip
drip DELAY
drip MIN_DELAY MAX_DELAY

drip takes lines of text via standard input and outputs them at a
limited rate. If no argument is given, a delay of 1s will be used
between lines. If one argument is given, a constant delay will be used.
If two arguments are given, a random delay between MIN_DELAY and MAX_DELAY
will be used.

Delays must be numbers, suffixed with a unit: us, ms, s, m or h.

If the delay is below 25ms, multiple lines will be printed at once due
to technical limitations. Delay ranges are not allowed here.

Examples:
yes 'Hello, world!' | drip 500ms
cat /path/to/chat/msgs | drip 10s 1m30s | send_chat_message -to paul
`)
	os.Exit(1)
}

var dur1 time.Duration
var dur2 time.Duration

func init() {
	if len(os.Args) > 3 {
		usage()
	} else if len(os.Args) == 1 {
		dur1 = time.Second
	} else {
		var err error
		dur1, err = time.ParseDuration(os.Args[1])
		if err != nil || dur1 == 0 {
			usage()
		}
		if len(os.Args) == 3 {
			dur2, err = time.ParseDuration(os.Args[2])
			if err != nil {
				usage()
			}
			if dur2 <= dur1 {
				fmt.Fprintf(os.Stderr, "Error: MAX_DELAY is smaller than or equal to MIN_DELAY.\n")
				os.Exit(1)
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Workaround for the fact that time.Ticker does not work well for
	// small durations.
	var linesPerTick int64 = 1
	if dur1 < 25 * time.Millisecond {
		if dur2 > 0 {
			fmt.Fprintln(os.Stderr, "Error: Delay ranges are not allowed, if the lowest delay is below 25ms.")
			os.Exit(1)
		}
		factor := 1 + 25 * time.Millisecond / dur1
		dur1 *= factor
		linesPerTick *= int64(factor)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error: could not read standard input:", err)
			os.Exit(1)
		}
		return
	}
	// Print first line immediately.
	fmt.Println(scanner.Text())

	var ticker *time.Ticker
	if dur2 > 0 {
		ticker = time.NewTicker(randomDelay())
	} else {
		ticker = time.NewTicker(dur1)
	}
	defer ticker.Stop()

	var i int64
	for i = 0; scanner.Scan(); i++ {
		if i % linesPerTick == 0 {
			<-ticker.C
		}
		fmt.Println(scanner.Text())
		if dur2 > 0 {
			ticker.Reset(randomDelay())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: could not read standard input:", err)
		os.Exit(1)
	}
}

func randomDelay() time.Duration {
	return time.Duration(rand.Int63n(int64(dur2-dur1))) + dur1
}
