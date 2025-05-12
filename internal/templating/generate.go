package templating

import (
	"bytes"
	"text/template"

	"github.com/rabuu/uni-cli/internal/exit"
)

func GenerateString(data TemplateData, templ string) (generated string) {
	t := template.Must(template.New("generated string").Parse(templ))

	var buf bytes.Buffer

	err := t.Execute(&buf, data)
	exit.ExitWithErr(err)

	generated = buf.String()
	return
}
