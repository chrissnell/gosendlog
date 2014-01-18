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

func sendLineToSyslog(message string, logger *syslog.Writer) {
	logger.Info(message)
}

func main() {

	var readFromStdin bool

	destPtr := flag.String("dest", "", "Destination host <host:port>")
	msgPtr := flag.String("msg", "", "Message <string>")
	tagPtr := flag.String("tag", "", "Tag <string>")

	flag.Parse()

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

	s, err := syslog.Dial("tcp", *destPtr, syslog.LOG_INFO|syslog.LOG_LOCAL6, *tagPtr)
	if err != nil {
		log.Fatal(err)
	}

	if !readFromStdin {
		err = s.Info(*msgPtr)
	} else {
		reader := bufio.NewReader(os.Stdin)
		ProcessLinesFromReader(reader, func(str string) { sendLineToSyslog(str, s) })
	}

	if err != nil {
		log.Fatal(err)
	}
}
