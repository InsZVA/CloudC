package utils
import (
	"regexp"
)

func Peek(bytes []byte,pattern string) ([][]string,error){
	reg,err := regexp.Compile(pattern)
	if(err != nil){return nil,err}
	return reg.FindAllStringSubmatch(string(bytes),-1),nil
}