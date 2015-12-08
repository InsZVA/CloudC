package utils
import (
	"testing"
	"fmt"
)

func TestPeek(t *testing.T){
	ret,_ := Peek([]byte("<meta charset=\"utf8\">79;5"),"(.+);(.+)")
	if(ret[0][1] != "1"){
		fmt.Println(ret)
	}
}