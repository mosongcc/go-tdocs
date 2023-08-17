package tdocs

import (
	"bytes"
	_ "embed"
	"html/template"
	"io"
	"log"
	"net/http"
)

//go:embed tdocs.html
var indexHtml string

var funcMap = template.FuncMap{
	/*"mapget": func(k string) []Field {
		return FieldMap[k]
	},*/
}

var tpl = template.Must(template.New("").Funcs(funcMap).Parse(indexHtml))

func Execute(wr io.Writer) (err error) {
	return tpl.Execute(wr, map[string]any{"tables": tables})
}

// Html 输出html页面
func Html() (b []byte, err error) {
	var wr = &bytes.Buffer{}
	err = Execute(wr)
	if err != nil {
		return
	}
	b = wr.Bytes()
	return
}

// HandleFunc HttpHandler输出页面
func HandleFunc(writer http.ResponseWriter, request *http.Request) {
	b, err := Html()
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	_, err = writer.Write(b)
	if err != nil {
		log.Print(err)
	}
}
