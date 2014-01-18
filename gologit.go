package main

import (
	"os"
	"io"
	"bufio"
	"flag"
	"log"
	"log/syslog"
)

// thanks shurcooL - https://gist.github.com/shurcooL/5157525
func SendLinesFromReader(r io.Reader, s *syslog.Writer) {
	br := bufio.NewReader(r)
	for line, err := br.ReadString('\n'); err == nil; line, err = br.ReadString('\n') {
		s.Info(line[:len(line)-1]) // Trim last newline
	}
}

func main() {

	var readFromStdin bool

	destPtr := flag.String("dest", "", "destination host -- \"host:port\"")
	msgPtr := flag.String("msg", "", "message -- \"string\" or \"-\" for stdin")
	tagPtr := flag.String("tag", "", "tag -- \"string\"")

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

	if ! readFromStdin  {
		err = s.Info(*msgPtr)
	} else {
		reader := bufio.NewReader(os.Stdin)
		SendLinesFromReader(reader, s)
	}

	if err != nil {
		log.Fatal(err)
	}
}
