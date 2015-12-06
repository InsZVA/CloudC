package main

import (
	"./system/dispatcher"
	"./system"
)

func main(){
	system.LoadConfig()
	dispatcher.Run()
	for i := 0;i < dispatcher.MinNum;i++{
		<- system.WorkerEnd
	}
}