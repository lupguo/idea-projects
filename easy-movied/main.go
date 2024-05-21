package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var address = `127.0.0.1:1313`

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// 创建一个新的 Echo 实例
	e := echo.New()
	e.HideBanner = true

	// 使用中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 设置模板渲染器
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./assets/templates/*.html")),
	}
	e.Renderer = renderer

	// 定义一个路由
	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"Title":   "Hello, Echo",
			"Message": "Welcome to Echo framework!",
		}
		return c.Render(http.StatusOK, "index.html", data)
	})

	// 启动服务器，监听端口 8080
	log.Infof(`website http://%s`, address)
	e.Logger.Fatal(e.Start(address))
}
