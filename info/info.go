package info

// Name 命令名字
const Name = "CodeFly"

// Usage 说明
const Usage = "iOS开发者工具集"

// Version 软件版本
const Version = "0.0.2"

// Author 作者
const Author = "Arror"

// Email 作者电子邮件
const Email = "763911422@qq.com"

// AppHelpTemplate 帮助模板
const AppHelpTemplate = `
名称:
    {{.Name}} - {{.Usage}}
作者:
    {{range .Authors}}{{ . }}{{end}}{{if .Commands}}
命令:
{{range .Commands}}{{if not .HideHelp}}    {{join .Names ", "}}{{ "\t" }}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
全局选项:
    {{range .VisibleFlags}}{{.}}{{end}}{{end}}{{if .Copyright }}
版权:
    {{.Copyright}}{{end}}{{if .Version}}
版本:
    {{.Version}}{{end}}
`
