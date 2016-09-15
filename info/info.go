package info

// Name 命令名字
const Name = "CodeFly"

// Usage 说明
const Usage = "Generate aim language data model"

// Version 软件版本
const Version = "0.0.3"

// Author 作者
const Author = "Arror"

// Email 作者电子邮件
const Email = "763911422@qq.com"

// AppHelpTemplate 帮助模板
const AppHelpTemplate = `
Name:
    {{.Name}} - {{.Usage}}
Authors:
    {{range .Authors}}{{ . }}{{end}}{{if .Commands}}
Commands:
{{range .Commands}}{{if not .HideHelp}}    {{join .Names ", "}}{{ "\t" }}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
Global:
    {{range .VisibleFlags}}{{.}}{{end}}{{end}}{{if .Copyright }}
Copyright:
    {{.Copyright}}{{end}}{{if .Version}}
Version:
    {{.Version}}{{end}}
`
