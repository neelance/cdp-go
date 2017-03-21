// +build generate

package main

import (
	"encoding/json"
	"os"
	"strings"
	"text/template"
)

type Protocol struct {
	Domains []*Domain
}

type Domain struct {
	Domain       string
	Dependencies []string
	Types        []*Type
	Commands     []*Command
	Events       []*Event
	Description  string
	Experimental bool
}

func (d *Domain) GoPackage() string {
	return strings.ToLower(d.Domain)
}

func (d *Domain) Doc() string {
	doc := d.Description
	if d.Experimental {
		doc += " (experimental)"
	}
	return strings.TrimSpace(doc)
}

type Type struct {
	ID           string
	Type         string
	Description  string
	Experimental bool
}

func (t *Type) Doc() string {
	doc := t.Description
	if t.Experimental {
		doc += " (experimental)"
	}
	return strings.TrimSpace(doc)
}

type Command struct {
	Name         string
	Parameters   []*Parameter
	Returns      []*Parameter
	Description  string
	Experimental bool
}

func (c *Command) GoName() string {
	return strings.ToUpper(c.Name[:1]) + c.Name[1:]
}

func (c *Command) GoOptsType() string {
	return c.GoName() + "Opts"
}

func (c *Command) GoResultType() string {
	return c.GoName() + "Result"
}

func (c *Command) Doc() string {
	doc := c.Description
	if c.Experimental {
		doc += " (experimental)"
	}
	return strings.TrimSpace(doc)
}

type Event struct {
	Name         string
	Parameters   []*Parameter
	Description  string
	Experimental bool
}

func (e *Event) GoName() string {
	return strings.ToUpper(e.Name[:1]) + e.Name[1:]
}

func (e *Event) GoType() string {
	return e.GoName() + "Event"
}

func (e *Event) Doc() string {
	doc := e.Description
	if e.Experimental {
		doc += " (experimental)"
	}
	return strings.TrimSpace(doc)
}

type Parameter struct {
	Name string
	TypeRef
	Description  string
	Optional     bool
	Experimental bool
}

func (p *Parameter) GoField() string {
	switch p.Name {
	case "url":
		return "URL"
	}
	return strings.ToUpper(p.Name[:1]) + p.Name[1:]
}

func (p *Parameter) Doc() string {
	doc := p.Description
	switch {
	case p.Optional && p.Experimental:
		doc += " (optional, experimental)"
	case p.Optional:
		doc += " (optional)"
	case p.Optional && p.Experimental:
		doc += " (experimental)"
	}
	return strings.TrimSpace(doc)
}

type TypeRef struct {
	Type  string
	Ref   string `json:"$ref"`
	Items *TypeRef
}

func (t *TypeRef) GoType() string {
	if t.Ref != "" {
		if strings.Contains(t.Ref, ".") {
			return "interface{}" // TODO
		}
		return t.Ref
	}

	switch t.Type {
	case "string":
		return "string"
	case "boolean":
		return "bool"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "array":
		return "[]" + t.Items.GoType()
	case "any", "object":
		return "interface{}"
	default:
		panic("unknown type: " + t.Type)
	}
}

func main() {
	in, err := os.Open("browser_protocol.json")
	if err != nil {
		panic(err)
	}
	var protocol Protocol
	if err := json.NewDecoder(in).Decode(&protocol); err != nil {
		panic(err)
	}
	in.Close()

	os.RemoveAll("protocol")
	os.Mkdir("protocol", 0777)

	for _, d := range protocol.Domains {
		dir := "protocol/" + d.GoPackage()
		os.Mkdir(dir, 0777)
		out, err := os.Create(dir + "/" + d.GoPackage() + ".go")
		if err != nil {
			panic(err)
		}
		if err := t.Execute(out, d); err != nil {
			panic(err)
		}
		out.Close()
	}

}

var t = template.Must(template.New("").Parse(`
{{if .Doc}}// {{.Doc}}{{end}}
package {{.GoPackage}}
{{$domain := .Domain}}
import (
	{{if .Events}}
		"encoding/json"
		"log"
	{{end}}
	"github.com/neelance/cdp-go/rpc"
)

{{if .Doc}}// {{.Doc}}{{end}}
type Domain struct {
	Client *rpc.Client
}

{{range .Types}}
	{{if .Doc}}// {{.Doc}}{{end}}
	type {{.ID}} interface{}
{{end}}

{{range .Commands}}
	{{if .Parameters}}
		type {{.GoOptsType}} struct {
			{{- range .Parameters}}
				{{if .Doc}}// {{.Doc}}{{end}}
				{{.GoField}} {{.GoType}} ` + "`" + `json:"{{.Name}}{{if .Optional}},omitempty{{end}}"` + "`" + `
			{{end}}
		}
	{{end}}

	{{if .Returns}}
		type {{.GoResultType}} struct {
			{{- range .Returns}}
				{{if .Doc}}// {{.Doc}}{{end}}
				{{.GoField}} {{.GoType}} ` + "`" + `json:"{{.Name}}"` + "`" + `
			{{end}}
		}

		{{if .Doc}}// {{.Doc}}{{end}}
		func (d *Domain) {{.GoName}}({{if .Parameters}}opts *{{.GoOptsType}}{{end}}) (*{{.GoResultType}}, error) {
			var result {{.GoResultType}}
			err := d.Client.Call("{{$domain}}.{{.Name}}", {{if .Parameters}}opts{{else}}nil{{end}}, &result)
			return &result, err
		}
	{{else}}
		{{if .Doc}}// {{.Doc}}{{end}}
		func (d *Domain) {{.GoName}}({{if .Parameters}}opts *{{.GoOptsType}}{{end}}) error {
			return d.Client.Call("{{$domain}}.{{.Name}}", {{if .Parameters}}opts{{else}}nil{{end}}, nil)
		}
	{{end}}
{{end}}

{{range .Events}}
	type {{.GoType}} struct {
		{{- range .Parameters}}
			{{if .Doc}}// {{.Doc}}{{end}}
			{{.GoField}} {{.GoType}} ` + "`" + `json:"{{.Name}}"` + "`" + `
		{{end}}
	}

	{{if .Doc}}// {{.Doc}}{{end}}
	func (d *Domain) On{{.GoName}}(listener func(*{{.GoType}})) {
		d.Client.AddListener("{{$domain}}.{{.Name}}", func(params json.RawMessage) {
			var event {{.GoType}}
			if err := json.Unmarshal(params, &event); err != nil {
				log.Print(err)
				return
			}
			listener(&event)
		})
	}
{{end}}
`))
