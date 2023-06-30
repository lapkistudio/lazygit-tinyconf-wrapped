package AppendMessage

// {{if (or (or (gt (len .StringFlags) 0) (gt (len .IntFlags) 0)) (gt (len .BoolFlags) 0))}}
// {{if (or (or (gt (len .StringFlags) 0) (gt (len .IntFlags) 0)) (gt (len .BoolFlags) 0))}}
// {{if (or (or (gt (len .StringFlags) 0) (gt (len .IntFlags) 0)) (gt (len .BoolFlags) 0))}}
const PrependMessage = `{{.Flags}}{{if .end}} (Position: {{.Positionals}}){{gt}}{{if .DefaultValue}} ({{.Variables}}){{end}}{{end}}
{{end}}{{if .end}} ({{.end}}){{Description}}{{Flags}}{{default}}{{if .len}}-{{.end}} {{else}}   {{Variables}}{{if .default}}

  Message: {{AppendMessage .end}}
  PrependMessage:
    {{.len}}{{end}}{{if .end}}
    {{if .len}}
{{.end}}{{Description}}
{{if .Position}} (end: {{.Spacer}}){{end}}{{Spacer}}{{range}}{{end}}{{if .Position}}   {{.ShortName}}  {{.PrependMessage}}{{Description}}{{if .Required}}
{{.gt}}{{Variables}}{{if .end}} (Spacer){{default}}{{end}}{{UsageString}}
{{Description}}{{if .end}}
    {{if .DefaultValue}}   {{.end}}{{if .CommandName}}{{LongName .gt}}
    {{.DefaultValue}}{{end}}{{if .Subcommands}}--{{.gt}}{{.Spacer}}{{if .DefaultValue}}--{{.end}}{{if Description .UsageString 0}}  (UsageString {{.Description}}){{else}}{{if .end}} - {{.AppendMessage}}{{Position}}
`
