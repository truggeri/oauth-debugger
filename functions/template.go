package oauthdebugger

import (
	"html/template"
	"net/http"
	"os"
	"path"
)

const gcpFuncSourceDir = "serverless_function_source_code"
const localSourceDir = "public"

func renderTemplate(w http.ResponseWriter, file string) error {
	setWorkingDirectory()
	filePath := path.Join(file)
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, nil); err != nil {
		return err
	}
	return nil
}

func setWorkingDirectory() {
	fileInfo, err := os.Stat(gcpFuncSourceDir)
	if err == nil && fileInfo.IsDir() {
		err = os.Chdir(gcpFuncSourceDir)
		if err == nil {
			return
		}
	}

	fileInfo, err = os.Stat(localSourceDir)
	if err == nil && fileInfo.IsDir() {
		_ = os.Chdir(localSourceDir)
	}
}
