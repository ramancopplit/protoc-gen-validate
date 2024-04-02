package goshared

const constTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			{{- if isEnum $f }}
			err := {{ errWithCode . $r.GetErrorCode  "value must equal " (enumVal $f $r.GetConst) }}
			{{- else }}
			err := {{ errWithCode . $r.GetErrorCode  "value must equal " $r.GetConst }}
			{{- end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
