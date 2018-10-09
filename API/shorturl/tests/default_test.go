package tests

import beetest "github.com/astaxie/beego/testing"
import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"fmt"
)
type ShortResult struct {
	Urlshort string
	Urllong string
}

func TestShort(t *testing.T)  {
	request:=beetest.Get("/v1/shorten")
	request.Param("longurl","http://www.beego.com")
	response,_:=request.Response()
	defer response.Body.Close()
	contens,_:=ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(contens,&s)
	fmt.Println(string(contens))
	if s.Urlshort!=""{

	}
}
func TestExpand(t *testing.T)  {
	request:=beetest.Get("/v1/expand")
	request.Param("shorturl","Fla5F")
	response,_:=request.Response()
	defer response.Body.Close()
	contents,_:=ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(contents,&s)
	fmt.Println(string(contents))
	if s.Urllong==""{
		t.Fatal("urllong is empty")
	}
}