package system
import (
	"../parser"
)

func loadConfig(){
	config,err := parser.ReadAndParseJSON("../test/config.json")
}