package main

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var lista = list.New()
var i = 0

func conteo() {
	for {
		for e := lista.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value, ": ", i)

		}
		fmt.Println("-------")
		i++
		time.Sleep(time.Second * 2)
	}
}

func serverExit() {
	s, error := net.Listen("tcp", ":9998")
	if error != nil {
		fmt.Println(error)
		fmt.Println("listen")
		return
	}
	//go conteo()
	for {
		c, error := s.Accept()
		if error != nil {
			fmt.Println(error)
			fmt.Println("accept")
			continue
		}
		go handleExitingClient(c)
	}
}

func handleExitingClient(c net.Conn) {
	var proc int
	error := gob.NewDecoder(c).Decode(&proc)
	if error != nil {
		fmt.Println(error)
	}
	lista.PushBack(proc)
}

func server() {
	s, error := net.Listen("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		fmt.Println("listen")
		return
	}
	go conteo()
	for {
		c, error := s.Accept()
		if error != nil {
			fmt.Println(error)
			fmt.Println("accept")
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	error := gob.NewEncoder(c).Encode(i)
	if error != nil {
		fmt.Println(error)
		fmt.Println("encode I")
		return
	}
	e := lista.Front()
	error = gob.NewEncoder(c).Encode(e.Value)
	if error != nil {
		fmt.Println(error)
		fmt.Println("encode E")
		return
	}
	lista.Remove(e)
}

func main() {
	for j := 1; j < 6; j++ {
		lista.PushBack(j)
	}
	go server()
	go serverExit()
	var pausa string
	fmt.Scanln(&pausa)
}
