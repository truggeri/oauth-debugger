package oauthdebugger

import (
	"html/template"
	"net/http"
	"path"
)

const gcpFuncSourceDir = "serverless_function_source_code"

func renderTemplate(w http.ResponseWriter, file string) error {
	filePath := path.Join(gcpFuncSourceDir, file)
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, nil); err != nil {
		return err
	}
	return nil
}
