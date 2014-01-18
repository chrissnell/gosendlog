package main

import (
		"flag"
		"log"
		"log/syslog"
)

func main()	{

	destPtr := flag.String("dest", "", "Destination host <host:port>")
	msgPtr := flag.String("msg", "", "Message <string>")

	flag.Parse()

	if *destPtr == "" {
		log.Fatal("Must pass a destination host. Use -h for help.")
	}

	if *msgPtr == "" {
		log.Fatal("Must pass a message to log.  Use -h for help.")
	}

	s, err :=  syslog.Dial("tcp", *destPtr, syslog.LOG_INFO|syslog.LOG_LOCAL6, "snelly_test")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Info(*msgPtr)
	if err != nil {
		log.Fatal(err)
	}
}
