## html/template

## 流程说明

1. 创建实例：创建template模版实例，通常使用 `template.Must()、template.New()`方式创建
2. 模版解析：从文件系统中读取对应到template文件，进行解析 ``
3. 数据准备：例如DB、文件等读取数据准备，使用 ``
4. 数据渲染：将数据渲染到已解析到文件中，通常使用 `template.ExecuteTemplate()`

## Go 方法

### template.Parse() - func (t *Template) Parse(text string) (*Template, error)

Parse函数**将文本解析为模板t的主体**。文本中的命名模板定义（{{define ...}}或{{block ...}}语句）定义了与t关联的其他模板，并从t自身的定义中删除。
在对t或任何相关模板执行第一次Execute之前，可以通过连续调用Parse来重新定义模板。

如果一个包含只有空格和注释的主体被认为是空白，它不会替换现有模板的主体。这允许使用Parse添加新的命名模板定义而不覆盖主要模板主体。
更多信息，请参考：https://golang.org/pkg/text/template/#Template.Parse

```go
baseTemplate := `
{{ define "content" }}
    <h1>{{ .Message }}</h1>
{{ end }}
`
tpl := template.Must(template.New("base").Parse(baseTemplate))
```

### template.ParseFiles() - func (t *Template) ParseFiles(filenames ...string) (*Template, error)

ParseFiles函数**解析指定的文件，并将生成的模板与t关联**。如果发生错误，解析过程停止并返回nil；否则返回t。至少要有一个文件。
- 多个目录同名解析，最后一个生效：当解析**具有相同名称但位于不同目录中的多个文件时，最后一个被提及的文件将是结果**。
- 不能重复parse: 如果t或**任何关联的模板已经执行过，则ParseFiles函数会返回错误**。

### template.Must() -  func Must(t *Template, err error) *Template

Must 是一个帮助程序，用于包装对返回的函数的调用（*Template，error），如果错误为非 nil，则会崩溃。
它旨在用于变量初始化，例如 `var t = template.Must(template.New("name").Parse("html"))`

### template.ParseFiles() - func (t *Template) ParseFiles(filenames ...string) (*Template, error)

ParseFiles函数解析指定的文件，并将生成的模板与t关联。

- 如果发生错误，解析过程停止并返回nil；否则返回t。至少要有一个文件。
  当解析具有相同名称但位于不同目录中的多个文件时，最后一个被提及的文件将是结果。
  如果t或任何关联的模板已经执行过，则ParseFiles函数会返回错误。

```go
tpl := template.New("base")
tpl, err := tpl.ParseFiles("base.html", "index.html")
```

### t.ExecuteTemplate()

- 定义: `func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error`
- 说明: ExecuteTemplate 将与具有给定名称的 t 关联的模板应用于指定的数据对象，并将输出写入 wr。
    - 如果在执行模板或写入其输出时发生错误，则执行将停止，但部分结果可能已写入输出编写器。
    - 模板可以安全地**并行执行**，但**如果并行执行共享一个 Writer，则输出可能会交错执行**。

```go
data := map[string]string{"Title": "Home Page", "Message": "Welcome to the Home Page"}
t.ExecuteTemplate(os.Stdout, "base.html", data)
```