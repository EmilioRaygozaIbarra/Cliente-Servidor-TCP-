package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var proc int

func mostrarProceso(ch chan bool) {
	var i int
	c, error := net.Dial("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
	}
	error = gob.NewDecoder(c).Decode(&i)
	if error != nil {
		fmt.Println(error)
	}
	error = gob.NewDecoder(c).Decode(&proc)
	if error != nil {
		fmt.Println(error)
	}
	c.Close()
	for {
		select {
		default:
			fmt.Println(proc, ": ", i)
			time.Sleep(time.Second * 2)
			i++
		case <-ch:
			break
		}

	}
}

func salir(ch chan bool) {
	ch <- true
	c, error := net.Dial("tcp", ":9998")
	if error != nil {
		fmt.Println(error)
	}
	error = gob.NewEncoder(c).Encode(proc)
	if error != nil {
		fmt.Println(error)
	}
	c.Close()
}

func main() {
	ch := make(chan bool)
	go mostrarProceso(ch)

	var pausa string
	fmt.Scanln(&pausa)
	salir(ch)
}
