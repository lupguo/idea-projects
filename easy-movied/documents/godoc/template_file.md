## template模版管理
- 目录结构：
  - 我们在 templates 目录中创建了两个模板文件，一个是基础模板 base.html，另一个是具体页面模板 index.html。
- 模板文件：
    - base.html 定义了 HTML 的基本结构，包括头部、主体和页脚，并使用 block 定义可插入的内容区域。
    - index.html 使用 define 定义了具体页面的标题和内容，这些内容将被插入到 base.html 中相应的 block 中。

### 模板中的变量、条件语句、循环

#### 变量
```go
t, err := template.New("example").Parse("Hello, {{.Name}}!")
data := map[string]string{"Name": "Alice"}
err = t.Execute(os.Stdout, data)
```

#### if...end 条件语句
```html
<!-- data := map[string]bool{"IsAdmin": true} -->
{{ if .IsAdmin }}
    <p>Welcome, Admin!</p>
{{ else }}
    <p>Welcome, User!</p>
{{ end }}
```

#### range...end 循环
```html
<!-- data := map[string][]string{"Items": {"Apple", "Banana", "Orange"}} -->
<ul>
{{ range .Items }}
    <li>{{ . }}</li>
{{ end }}
</ul>
```

### template 嵌套

#### base.html
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
</head>
<body>
    {{ template "content" . }}
</body>
</html>
```

#### index.html
```html
{{ define "content" }}
    <h1>{{ .Message }}</h1>
{{ end }}
```

#### go代码
```go
baseTemplate := `
{{ define "content" }}
    <h1>{{ .Message }}</h1>
{{ end }}
`

// 模版文件
t := template.Must(template.New("base").Parse(baseTemplate))
t, _ = t.ParseFiles("base.html", "index.html")

data := map[string]string{"Title": "Home Page", "Message": "Welcome to the Home Page"}

t.ExecuteTemplate(os.Stdout, "base.html", data)
```