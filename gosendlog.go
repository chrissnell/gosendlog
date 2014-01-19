package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
)

// Syslog priorities from /usr/include/sys/syslog.h
var priorityStrings = map[string]syslog.Priority{
	"emerg":   0,
	"alert":   1,
	"crit":    2,
	"err":     3,
	"warning": 4,
	"notice":  5,
	"info":    6,
	"debug":   7,
}

var facilityStrings = map[string]syslog.Priority{
	"kern":     0 << 3,
	"user":     1 << 3,
	"mail":     2 << 3,
	"daemon":   3 << 3,
	"auth":     4 << 3,
	"syslog":   5 << 3,
	"lpr":      6 << 3,
	"news":     7 << 3,
	"uucp":     8 << 3,
	"cron":     9 << 3,
	"authpriv": 10 << 3,
	"ftp":      11 << 3,
	"local0":   16 << 3,
	"local1":   17 << 3,
	"local2":   18 << 3,
	"local3":   19 << 3,
	"local4":   20 << 3,
	"local5":   21 << 3,
	"local6":   22 << 3,
	"local7":   23 << 3,
}

func sendLineToSyslog(message []byte, logger *syslog.Writer) {
	logger.Write(message)
}

func main() {

	var readFromStdin bool

	destPtr := flag.String("dest", "", "Destination host <host:port>")
	msgPtr := flag.String("msg", "", "Message <string>")
	tagPtr := flag.String("tag", "", "Tag <string>")
	prioPtr := flag.String("priority", "info", "Priority (default: info)")
	facilPtr := flag.String("facility", "local0", "Facility (default: local0)")
	protoPtr := flag.String("proto", "udp", "Protocol <udp/tcp>")

	flag.Parse()

	mappedPriority := priorityStrings[*prioPtr]
	mappedFacility := facilityStrings[*facilPtr]

	if *destPtr == "" {
		log.Fatal("Must pass a destination host. Use -h for help.")
	}

	if *msgPtr == "-" || *msgPtr == "" {
		readFromStdin = true
	} else if *msgPtr == "" {
		log.Fatal("Must pass a message to log.  Use -h for help.")
	}

	if *tagPtr == "" {
		log.Fatal("Must pass a tag.  Use -h for help.")
	}

	s, err := syslog.Dial(*protoPtr, *destPtr, mappedPriority|mappedFacility, *tagPtr)

	if err != nil {
		log.Fatal(err)
	}

	if !readFromStdin {
		msg := []byte(*msgPtr)
		sendLineToSyslog(msg, s)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			sendLineToSyslog([]byte(scanner.Text()), s)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
