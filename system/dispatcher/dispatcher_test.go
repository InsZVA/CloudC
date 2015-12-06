package dispatcher

import (
	".."
	"testing"
)

func TestLoadTaskQueue(t *testing.T){
	LoadTaskQueue()
	if TaskQueue[0].databaseTemplate != "tecentKt"{
		t.Errorf("error")
	}
}

func TestRun(t *testing.T){
	system.LoadConfig()
	tq := Run()
	if tq[0][0].databaseTemplate != "tecentKt"{
		t.Errorf("error")
	}
}