package main

import (
	"log"
	"os/exec"
	"time"
)

func Run() {
	cmd := exec.Command("hello.exe")

	err := cmd.Start()
	if err != nil {
		log.Println(err.Error())
	}
	err = cmd.Wait()
	cmd = nil
}
func main() {
	for {
		Run()
		time.Sleep(5 * time.Second)
	}
}
