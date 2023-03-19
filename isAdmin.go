package main

import (
	"C"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
)

//export isAdmin
func isAdmin() bool {
	cmd := exec.Command("net", "session")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	file := filepath.Join("isAdmin.dll")
	dll, err := syscall.LoadDLL(file)
	if err != nil {
		log.Println(err)
		return
	}
	proc, err := dll.FindProc("isAdmin")
	if err != nil {
		log.Println(err)
		return
	}
	v, _, _ := proc.Call()
	sp := fmt.Sprint(v)
	if i, _ := strconv.Atoi(sp); i == 1 {
		fmt.Println("Hi Administrator")
	} else {
		fmt.Println("Hi Friend")
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
