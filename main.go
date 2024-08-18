package main

import (
	"flag"
	"recvnet0/lib"
)

func main() {
	// what I am trying to do by this package is not building the tool itself
	// but getting used to programming with the go programming language.

	//step1: get host ip
	var ip string

	flag.StringVar(&ip, "i", "192.168.1.1", "ip address for remote host.")
	flag.Parse()

	//step2: ping host
	//lib.RecognizeHost("titan.picoctf.net", 51975) this one returns SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.11;
	//lib.RecognizeHost("google.com", 443) this one returns nothing so many tests will be needed to check if the server is an http server.
	lib.DialUDPRemote("ede", 101)
}
