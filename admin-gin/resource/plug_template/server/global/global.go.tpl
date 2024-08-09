package global

{{- if .HasGlobal }}

import "admin-gin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}