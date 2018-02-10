package s2h

import (
	"github.com/m0a/s2h/struct2json"
	"os"
	"encoding/json"
	"fmt"
)



func Save(obj interface{}) {
	v := struct2json.Create(obj)
	save("test.json",v)
}

func save(filename string, v struct2json.Struct2json) {
	file, err := os.OpenFile("test.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer  file.Close()
	if err != nil {
		os.Exit(-1)
	}
	data, err := json.Marshal(v)
	if err != nil {
		os.Exit(-1)
	}
	fmt.Fprint(file,string(data))
}