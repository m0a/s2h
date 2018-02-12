package s2h

import (
	"fmt"
	"os"

	"github.com/m0a/s2h/reflect2json"
	"text/template"
	"bytes"
)

// for save json
func SaveJSON(obj interface{}) {
	v := reflect2json.Reflect2JSON(obj)
	save("test.json", v)
}


type JSONTemplateData struct {
	JSON        string
	JS string
	STYLE string
}

// save html
func Save(obj interface{}) {
	fs := FS(false)
	indexJs, err := fs.Open("/static/index.js")
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error can't open index.js err: %s", err.Error())
		os.Exit(-1)
	}
	indexCss, err := fs.Open("/static/index.css")
	if err != nil {
		fmt.Fprint(os.Stderr,"Error can't open index.css")
		os.Exit(-1)
	}

	var templateData JSONTemplateData

	{
		buf := new(bytes.Buffer)
		buf.ReadFrom(indexJs)
		templateData.JS = buf.String()
	}
	{
		buf := new(bytes.Buffer)
		buf.ReadFrom(indexCss)
		templateData.STYLE = buf.String()
	}

	{
		v := reflect2json.Reflect2JSON(obj)
		templateData.JSON = v;
	}


	tmpl, err := template.New("index.html").Parse(`<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>STRUCT2HTML</title>
  <style>{{.STYLE}}</style>
<body>
  <script> window.testJSON = {{.JSON}};
  </script>
  <script>{{.JS}}</script> 
</body>
</html>
`)
	if err != nil {
		fmt.Fprint(os.Stderr,"Error can't open template")
		os.Exit(-1)
	}

	filename:= "test.html"
	os.Remove(filename)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		os.Exit(-1)
	}

	tmpl.Execute(file, templateData)


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
