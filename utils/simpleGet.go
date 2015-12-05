package utils
import (
	"net/http"
	"strconv"
	"io/ioutil"
)

func buildQuery(params []interface{})(string){
	queryString := "?"
	i := 0
	for _,param := range params {
		if i != 0 {queryString += "&"}
		param_,_ := param.(map[string]interface{})
		queryString += param_["name"].(string)
		range_ := param_["range"].([]interface{})
		start := range_[0].(float64)
		queryString += "=" + strconv.FormatFloat(start,'f',-1,64)
		i++
	}
	return queryString
}

func SimpleGet(url string,params []interface{}) ([]byte,error){
	resp,err := http.Get(url + buildQuery(params))
	if(err != nil){return nil,err}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if(err != nil){return nil,err}
	return body,nil
}