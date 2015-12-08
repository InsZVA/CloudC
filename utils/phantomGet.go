package utils
import (
	"io/ioutil"
	"os/exec"
)

func PhantomGet(url string,params []interface{}) ([]byte,error){
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