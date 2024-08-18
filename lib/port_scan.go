package lib

import (
	"log"
	"net"
	"recvnet0/util"
	"strconv"
	"time"
)

func TestTCPRemote(remote string, ports []int) {
	// tcp based protocols are :
	// HTTP, FTP, SMTP, POP3, IMAP, TELNET, SSH
	open_ports := make([]int, 0)
	close_ports := make([]int, 0)
	for port := range ports {
		conn, err := net.DialTimeout("tcp", remote+strconv.Itoa(port), time.Second*15)
		if err == nil {
			open_ports = append(open_ports, port)
		} else {
			close_ports = append(close_ports, port)
		}
		defer conn.Close()
	}
	// the flow goes like this the connection from the server comes and then get's tested for TCP then UDP
	TestTCPRemote(remote, close_ports)
}

func TestUDPRemote(remote string, ports []int) {

	for port := range ports {
		DialUDPRemote(remote, port)
	}

}
func DialUDPRemote(remote string, port int) {
	remoteAddr, err := net.ResolveUDPAddr("udp", "1.1.1.1:53")
	util.CheckErr(err)

	localAddr, err := net.ResolveUDPAddr("udp", ":0") // :0 lets the system assign an available port
	util.CheckErr(err)

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	util.CheckErr(err)
	log.Println("connected")
	buff := make([]byte, 4096)

	n, err := conn.Read(buff)

	util.CheckErr(err)

	response := string(buff[:n])
	log.Println(response)

	defer conn.Close()
}
