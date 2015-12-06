package system
import (
	"testing"
)

func TestLoadConfig(t *testing.T){
	LoadConfig()
	if Config.Workers != 7{
		t.Errorf("Error!")
	}
}