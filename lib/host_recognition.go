package lib

import (
	"log"
	"net"
	"recvnet0/util"
	"strconv"
	"strings"
	"time"
)

func RecognizeHost(remote string, port int) {
	conn, err := net.DialTimeout("tcp", remote+":"+strconv.Itoa(port), time.Second*10)
	log.Println("Connected")
	util.CheckErr(err)

	buff := make([]byte, 4096)

	n, err := conn.Read(buff)
	util.CheckErr(err)
	response := string(buff[:n])
	log.Println(response)

	defer conn.Close()
}

func parse_response(response string) []string {
	parts := strings.Split(response, " ")
	return parts
}
