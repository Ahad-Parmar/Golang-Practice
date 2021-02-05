package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)



func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func main() {
	que := make(map[string][]string, 10)

	Arguments := os.Args
	if len(Arguments) == 1 {
		fmt.Println("PORT number")
		return
	}

	PORT := ":" + Arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	var queue []string
	var username string
	for {
		newdata, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(newdata)) == "STOP!!" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("Received message -> ", string(newdata))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))

		if strings.TrimSpace(string(newdata)) == "new" {
			fmt.Println("New user")

			
			newdata, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			username = string(newdata)
			que[username] = queue

		} else {
			queue = append(queue, string(newdata))
			
			que[username] = queue
		}

		if strings.TrimSpace(string(newdata)) == "BREAK" {
			newdata, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Command:", string(newdata))
			queue = nil
			
		}

		if strings.TrimSpace(string(newdata)) == "PRINT" {
		
			for key, value := range que {
				value = value[:len(value)-1]
				fmt.Printf("%s value is %v\n", key, value)

			}
			
		}

		if strings.TrimSpace(string(newdata)) == "FIRST" {
			
			newdata, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			username = string(newdata)

			for key, value := range que {
				if key == username {
					fmt.Println("The first message sent by ", key, " is ", value[:1])
				}

			}
			
		}

		}
	}
}
