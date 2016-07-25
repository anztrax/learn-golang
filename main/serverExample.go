package main

import (
	"encoding/gob" //go binary format ?
	"fmt"
	"net" //network
)

func server() {
	ln, err := net.Listen("tcp", ":9099")
	if err != nil {
		fmt.Println(err)
		return
	}

	//looping forever
	for {
		//accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//recieving the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg) //decode need pointer
	if err != nil {
		fmt.Println("err when decode ", err)
	} else {
		fmt.Println("Recieved", msg)
	}

	c.Close()
}

//this is stimulating client
func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9099")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
