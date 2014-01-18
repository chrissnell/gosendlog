package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"log/syslog"
	"os"
)

func ProcessLinesFromReader(r io.Reader, processFunc func(string)) {
	br := bufio.NewReader(r)
	for line, err := br.ReadString('\n'); err == nil; line, err = br.ReadString('\n') {
		processFunc(line[:len(line)-1]) // Trim last newline
	}
}

func sendLineToSyslog(message []byte, logger *syslog.Writer) {
	logger.Write(message)
	// logger.Info(message)
}

func main() {

	var readFromStdin bool

	var prio map[string]syslog.Priority
	prio = make(map[string]syslog.Priority)

	prio["emerg"] = 0
	prio["alert"] = 1
	prio["crit"] = 2
	prio["err"] = 3
	prio["warning"] = 4
	prio["notice"] = 5
	prio["info"] = 6
	prio["debug"] = 7

	var facil map[string]syslog.Priority
	facil = make(map[string]syslog.Priority)

	facil["kern"] = 0 << 3
	facil["user"] = 1 << 3
	facil["mail"] = 2 << 3
	facil["daemon"] = 3 << 3
	facil["auth"] = 4 << 3
	facil["syslog"] = 5 << 3
	facil["lpr"] = 6 << 3
	facil["news"] = 7 << 3
	facil["uucp"] = 8 << 3
	facil["cron"] = 9 << 3
	facil["authpriv"] = 10 << 3
	facil["ftp"] = 11 << 3
	facil["local0"] = 16 << 3
	facil["local1"] = 17 << 3
	facil["local2"] = 18 << 3
	facil["local3"] = 19 << 3
	facil["local4"] = 20 << 3
	facil["local5"] = 21 << 3
	facil["local6"] = 22 << 3
	facil["local7"] = 23 << 3

	destPtr := flag.String("dest", "", "Destination host <host:port>")
	msgPtr := flag.String("msg", "", "Message <string>")
	tagPtr := flag.String("tag", "", "Tag <string>")
	prioPtr := flag.String("priority", "info", "Priority (default: info)")
	facilPtr := flag.String("facility", "local0", "Facility (default: local0)")
	protoPtr := flag.String("proto", "udp", "Protocol <udp/tcp>")

	flag.Parse()

	mappedPriority := prio[*prioPtr]
	mappedFacility := facil[*facilPtr]

	if *destPtr == "" {
		log.Fatal("Must pass a destination host. Use -h for help.")
	}

	if *msgPtr == "-" {
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
		err = s.Info(*msgPtr)
	} else {
		reader := bufio.NewReader(os.Stdin)
		ProcessLinesFromReader(reader, func(str string) { sendLineToSyslog([]byte(str), s) })
	}

	if err != nil {
		log.Fatal(err)
	}
}
