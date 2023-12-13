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

Examples:
yes 'Hello, world!' | drip 500ms
cat /path/to/chat/msgs | drip 10s 1m30s | send_chat_message -to paul
```
