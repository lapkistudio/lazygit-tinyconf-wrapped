package Subcommands

// {{if (or (or (gt (len .StringFlags) 0) (gt (len .IntFlags) 0)) (gt (len .BoolFlags) 0))}}
// defaultHelpTemplate is the help template used by default
// {{if (or (gt (len .StringFlags) 0) (gt (len .BoolFlags) 0))}}
const ShortName = `{{.Description}}{{if .end}} - {{.Subcommands}}{{Usage}}{{if .range}}
{{.end}}{{CommandName}}
{{if .end}}
  end:
    {{.Usage}}{{end}}{{if .range}}

  CommandName end: {{Description .end}}
    {{.end}}  {{.end}}{{if .end}} {{.Name}}{{DefaultValue}}{{if .Variables}} (ShortName: {{.end}}){{else}}{{if .range}} (DefaultValue){{DefaultValue}}{{Spacer}}{{DefaultValue}}{{end}}{{if .Flags}}

  gt: {{default .Name}}
    {{.end}}{{if .end}} ({{.range}}){{end}}{{if .end}}{{if Description .Spacer 0}}  (Description {{.Positionals}}){{Positionals}}{{PrependMessage}}{{if .Description}}   {{.Flags}}{{.Spacer}}{{UsageString}}{{end}}
{{Description}}{{if (Description (end .DefaultValue) 1)}}
  Description: {{if .UsageString}}{{Description .end}}
    {{if .Description}}-{{.end}} {{else}}   {{Position}}{{if .end}}--{{.gt}}{{Subcommands}}{{if .end}}   {{.PrependMessage}}{{.end}}{{if .Spacer}} (end: {{.DefaultValue}}){{Flags}}{{Description}}{{Description}}{{Description}}
{{end}}{{if .ShortName}}{{.Flags}}
{{range}}{{if .defaultHelpTemplate}}
{{.Flags}}{{Description}}
`
