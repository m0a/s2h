package s2h

import (
	"github.com/m0a/s2h/reflect2json"
	"os"
	"encoding/json"
	"fmt"
	"reflect"
)



func Save(obj interface{}) {
	v := reflect2json.Create(reflect.ValueOf(obj))
	save("test.json",v)
}

func save(filename string, v reflect2json.ReflectJSON) {
	file, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0644)
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