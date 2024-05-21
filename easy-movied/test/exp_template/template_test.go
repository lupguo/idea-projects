package exp_template

import (
	"html/template"
	"os"
	"testing"
)

func TestHtmlTemplateExample1(t *testing.T) {
	// 	baseTemplate := `
	// {{ define "content" }}
	//     <h1>{{ .Message }}</h1>
	// {{ end }}
	// `

	// 模版文件
	// tpl := template.Must(template.New("base").Parse(baseTemplate))
	tpl := template.New("base")
	tpl, err := tpl.ParseFiles("base.html", "index.html")
	if err != nil {
		t.Errorf("parse files got err: %v", err)
	}

	// 渲染数据
	data := map[string]string{"Title": "Home Page", "Message": "Welcome to the Home Page"}

	// 数据填充到模版
	err = tpl.ExecuteTemplate(os.Stdout, "base.html", data)
	if err != nil {
		t.Errorf("exec tpl got err: %v", err)
	}
}
