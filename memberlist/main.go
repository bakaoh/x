package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	go create()
	go join()
	time.Sleep(100 * time.Second)
}

func create() {
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	n, err := list.Join([]string{"localhost"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	fmt.Println("create", "n", n)
	for {
		for _, member := range list.Members() {
			fmt.Printf("Create Member: %s %s:%d\n", member.Name, member.Addr, member.Port)
		}

		time.Sleep(2 * time.Second)
	}
}

func join() {
	conf := memberlist.DefaultLocalConfig()
	conf.BindPort = 0
	conf.Name = "Slave"
	list, err := memberlist.Create(conf)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	n, err := list.Join([]string{"localhost:7946"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	fmt.Println("join", "n", n)
	for {
		for _, member := range list.Members() {
			fmt.Printf("Join Member: %s %s:%d\n", member.Name, member.Addr, member.Port)
		}

		time.Sleep(2 * time.Second)
	}
}
