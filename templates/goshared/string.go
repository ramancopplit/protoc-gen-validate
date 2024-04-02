package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if {{ accessor . }} != "" {
	{{ end }}

	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetLen }} {
			err := {{ err . "msg:value length must be " $r.GetLen " runes, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		{{ else }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
			err := {{ err . "msg:value length must be " $r.GetMinLen " runes, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		{{ end }}
	}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				err := {{ err . "msg:value length must be between " $r.GetMinLen " and " $r.GetMaxLen " runes, inclusive, validate_error_code:"  $r.GetErrorCode }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				err := {{ err . "msg:value length must be at least " $r.GetMinLen " runes, validate_error_code:"  $r.GetErrorCode }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			err := {{ err . "msg:value length must be at most " $r.GetMaxLen " runes, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
			if len({{ accessor . }}) != {{ $r.GetLenBytes }} {
				err := {{ err . "msg:value length must be " $r.GetLenBytes " bytes, validate_error_code:"  $r.GetErrorCode }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinBytes }} {
				err := {{ err . "msg:value length must be " $r.GetMinBytes " bytes, validate_error_code:"  $r.GetErrorCode }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinBytes }} || l > {{ $r.GetMaxBytes }} {
					err := {{ err . "msg:value length must be between " $r.GetMinBytes " and " $r.GetMaxBytes " bytes, inclusive, validate_error_code:"  $r.GetErrorCode }}
					if !all { return err }
					errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinBytes }} {
				err := {{ err . "msg:value length must be at least " $r.GetMinBytes " bytes, validate_error_code:"  $r.GetErrorCode }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxBytes }}
		if len({{ accessor . }}) > {{ $r.GetMaxBytes }} {
			err := {{ err . "msg:value length must be at most " $r.GetMaxBytes " bytes, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !strings.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			err := {{ err . "msg:value does not have prefix " (lit $r.GetPrefix) ", validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !strings.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			err := {{ err . "msg:value does not have suffix " (lit $r.GetSuffix) ", validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Contains }}
		if !strings.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			err := {{ err . "msg:value does not contain substring " (lit $r.GetContains) ", validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.NotContains }}

		if strings.Contains({{ accessor . }}, {{ lit $r.GetNotContains }}) {
			err := {{ err . "msg:value contains substring " (lit $r.GetNotContains) ", validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			err := {{ err . "msg:value must be a valid IP address, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			err := {{ err . "msg:value must be a valid IPv4 address, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			err := {{ err . "msg:value must be a valid IPv6 address, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			err = {{ errCause . "err" "msg:value must be a valid email address, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			err = {{ errCause . "err" "msg:value must be a valid hostname, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetAddress }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			if ip := net.ParseIP({{ accessor . }}); ip == nil {
				err := {{ err . "msg:value must be a valid hostname, or ip address, validate_error_code:"  $r.GetErrorCode }}
				if !all { return err }
				errors = append(errors, err)
			}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			err = {{ errCause . "err" "msg:value must be a valid URI, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		} else if !uri.IsAbs() {
			err := {{ err . "msg:value must be absolute, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			err = {{ errCause . "err" "msg:value must be a valid URI, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUuid }}
		if err := m._validateUuid({{ accessor . }}); err != nil {
			err = {{ errCause . "err" "msg:value must be a valid UUID, validate_error_code:"  $r.GetErrorCode }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Pattern }}
		if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
			err := {{ err . "msg:value does not match regex pattern " (lit $r.GetPattern) ", validate_error_code:"  $r.GetErrorCode}}
			if (lit $r.GetErrorCode) != 0{
               err.code = (lit $r.GetErrorCode)
			}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}

`
