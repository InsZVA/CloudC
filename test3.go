package main

import (
	"./system/dispatcher"
	"./system"
)

func main(){
	system.LoadConfig()
	go system.IO()
	dispatcher.Run()
	for i := 0;i < dispatcher.MinNum;i++{
		<- system.WorkerEnd
	}
	for !system.Commit{}
}