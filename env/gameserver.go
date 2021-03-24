package env

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


var consulAddr string 
var natsAddr string

func init() {

	dataIo, err := ioutil.ReadFile("./env.json")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(dataIo))
    }

	envMap:=make(map[string]string)

	data := (string(dataIo))

	json.Unmarshal([]byte(data),&envMap)

	var ok bool
	consulAddr, ok = envMap["consul_addr"];
	if ok {
		
	} else {

	}

	natsAddr, ok = envMap["nats_addr"];
	if ok {
		
	} else {

	}

	fmt.Println("============consul: " + consulAddr + ", nats: " + natsAddr)

}

func GetRsAddr() string {
	return consulAddr
}


func GetNcAddr() string {
	return natsAddr
}



