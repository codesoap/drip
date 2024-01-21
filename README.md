drip takes data via standard input and outputs it, one line at a time,
via standard output. After each line has been printed, drip waits for a
while before printing the next line.

This is useful if you want to limit the rate at which lines are fed to a
program, e.g. to limit resource usage.

# Usage
```console
$ drip -h
Usage:
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
```

# Installation
You can download precompiled binaries from the [releases
page](https://github.com/codesoap/drip/releases) or install it with
`go install github.com/codesoap/drip@latest`.
