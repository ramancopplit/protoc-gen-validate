package goshared

const durationcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if dur != {{ durLit $r.Const }} {
					err := {{ errWithCode . $r.GetErrorCode  "value must equal " (durStr $r.Const) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{  if $r.Lt }}  lt  := {{ durLit $r.Lt }};  {{ end }}
			{{- if $r.Lte }} lte := {{ durLit $r.Lte }}; {{ end }}
			{{- if $r.Gt }}  gt  := {{ durLit $r.Gt }};  {{ end }}
			{{- if $r.Gte }} gte := {{ durLit $r.Gte }}; {{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLt $r.GetGt }}
						if dur <= gt || dur >= lt {
							err := {{ errWithCode . $r.GetErrorCode  "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur <= gt {
							err := {{ errWithCode . $r.GetErrorCode  "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
						if dur < gte || dur >= lt {
							err := {{ errWithCode . $r.GetErrorCode  "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur < gte {
							err := {{ errWithCode . $r.GetErrorCode  "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur >= lt {
						err := {{ errWithCode . $r.GetErrorCode  "value must be less than " (durStr $r.GetLt) }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
						if dur <= gt || dur > lte {
							err := {{ errWithCode . $r.GetErrorCode  "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur <= gt {
							err := {{ errWithCode . $r.GetErrorCode  "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
						if dur < gte || dur > lte {
							err := {{ errWithCode . $r.GetErrorCode  "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur < gte {
							err := {{ errWithCode . $r.GetErrorCode  "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur > lte {
						err := {{ errWithCode . $r.GetErrorCode  "value must be less than or equal to " (durStr $r.GetLte) }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Gt }}
				if dur <= gt {
					err := {{ errWithCode . $r.GetErrorCode  "value must be greater than " (durStr $r.GetGt) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.Gte }}
				if dur < gte {
					err := {{ errWithCode . $r.GetErrorCode  "value must be greater than or equal to " (durStr $r.GetGte) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{ if $r.In }}
				if _, ok := {{ lookup $f "InLookup" }}[dur]; !ok {
					err := {{ errWithCode . $r.GetErrorCode  "value must be in list " $r.In }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.NotIn }}
				if _, ok := {{ lookup $f "NotInLookup" }}[dur]; ok {
					err := {{ errWithCode . $r.GetErrorCode  "value must not be in list " $r.NotIn }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
`
