package oauthdebugger

import (
	"html/template"
	"net/http"
	"path"
)

const GCP_SOURCE_DIR = "serverless_function_source_code"

func renderTemplate(w http.ResponseWriter, file string, data interface{}) error {
	filePath := path.Join(GCP_SOURCE_DIR, file)
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	return err
}
