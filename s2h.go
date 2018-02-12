package s2h

import (
	"fmt"
	"os"

	"github.com/m0a/s2h/reflect2json"
)

// for save json
func SaveJSON(obj interface{}) {
	v := reflect2json.Reflect2JSON(obj)
	save("test.json", v)
}

// TODO: for save html
func Save(obj interface{}) {
	v := reflect2json.Reflect2JSON(obj)
	save("test.json", v)
}

func save(filename string, json string) {

	os.Remove(filename)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		os.Exit(-1)
	}

	fmt.Fprint(file, json)
}
