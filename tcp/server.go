package tcp

import (
	"P2P-File-Sharing/common"
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buffer [512]byte
	conn.Read(buffer[:])
	fmt.Println("Someone requested for" + string(buffer[:]))

	_, err := conn.Write([]byte("I'll send you that!"))
	checkError(err)
}

// Server ...
func Server(myNode common.Node) {

	service := myNode.IP + ":" + myNode.TCPPort
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}

}