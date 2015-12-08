package utils
import (
	"net/http"
	"strconv"
	"io/ioutil"
)

type paramNode struct{
	name string
	rangeEnum []string
	pos int
}

func buildQuery(paramStack []*paramNode)([]string){
	length := len(paramStack)
	pos := length - 1
	var querys []string
	//Initialize
	for i := 0;i < length;i++{
		paramStack[i].pos = 0
	}
	for{
		//Create a piece of query
		query := "?"
		for i := 0;i < length;i++{
			if(i != 0){query += "&"}
			query += paramStack[i].name + "=" + paramStack[i].rangeEnum[paramStack[i].pos]
		}
		querys = append(querys,query)
		for pos >= 0 && paramStack[pos].pos >= len(paramStack[pos].rangeEnum) - 1{
			paramStack[pos].pos = 0
			pos--
		}
		if(pos < 0){break}
		paramStack[pos].pos++
		pos = length - 1
	}
	return querys
}

func buildStack(params []interface{})([]*paramNode){
	var stack []*paramNode 
	for _,param := range params {
		node := new(paramNode)
		param_,_ := param.(map[string]interface{})
		node.name = param_["name"].(string)
		type_ := param_["type"].(string)
		switch(type_){
			case "int":
				ranges := param_["range"].([]interface{})
				for _,rangeItem := range ranges{
					rangeItem_ := rangeItem.(map[string]interface{})
					start := int(rangeItem_["start"].(float64))
					end := int(rangeItem_["end"].(float64))
					for ex:= start;ex <= end;ex++{
						node.rangeEnum = append(node.rangeEnum,strconv.Itoa(ex))
					}
				}
		}
		stack = append(stack,node)
	}
	return stack
}

func SimpleGet(url string,params []interface{}) ([]byte,error){
	querys := buildQuery(buildStack(params))
	var resps []byte
	for _,query := range querys{
		resp,err := http.Get(url + query)
		if(err != nil){return nil,err}
		defer resp.Body.Close()
		body,err := ioutil.ReadAll(resp.Body)
		resps = append(resps,body...)
		if(err != nil){return nil,err}
	}
	return resps,nil
}