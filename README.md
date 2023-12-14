# gooser

OSINT username search tool.

Uses [Whats My Name](https://github.com/WebBreacher/WhatsMyName) json file to call each site and see what the status code is. Depending on status code type it will result in a "Hit" or a "Nope". 

gooser is a play on go-user (gouser). Think about it. 

Made this for fun, but I think it could be a great beginner project for people who are interested in OSINT/Cybersecurity and Golang.

Run it with
```
go run *.go -username={username}
```

Extension ideas:
- Clean up the functions by creating another file to hold them.
- Add more flags so users can narrow down by category
- Add flag for most common socials so it doesn't have to go through all 500+ sites
- Make it concurrent (Can you make it faster? Good chance to mess around with go routines, channels, mutex, waitgroup etc)