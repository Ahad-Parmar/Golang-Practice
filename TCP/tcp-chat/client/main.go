package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	Arguments := os.Args
	if len(Arguments) == 1 {
		fmt.Println("Providing host:Port")
		return
	}

	CONNECT := Arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP!!" {
			fmt.Println("Stopping TCP server")
			return
		}
	}
}
