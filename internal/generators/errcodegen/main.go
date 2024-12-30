package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"go/format"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
	"text/template"
)

//go:embed twilio_errors.json
var errorCodeJsonBytes []byte

//go:embed codes.tmpl
var tmpl string

type errInfo struct {
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
	LogLevel string `json:"log_level,omitempty"`
	LogType  string `json:"log_type,omitempty"`
	Product  string `json:"product,omitempty"`
	// date_created, last_updated
	// secondary_message
	// docs
	// causes
	// solutions
	// description
}

type tmplError struct {
	Code    int
	Message string
}

type tmplData struct {
	ErrorList []errInfo
}

const (
	filename = `errors_gen.go`
)

func main() {
	rawErrorList := make([]errInfo, 0, 2500)

	if err := json.Unmarshal(errorCodeJsonBytes, &rawErrorList); err != nil {
		panic(err)
	}

	errorMap := make(map[int]errInfo)

	for _, entry := range rawErrorList {
		if !relevantError(entry) {
			continue
		}

		// Sigh...
		entry.Message = strings.ReplaceAll(entry.Message, `‘`, `'`)
		entry.Message = strings.ReplaceAll(entry.Message, `’`, `'`)

		errorMap[entry.Code] = entry
	}

	codes := slices.Collect(maps.Keys(errorMap))
	slices.Sort(codes)

	errorList := make([]errInfo, 0, len(errorMap))

	for _, code := range codes {
		errorList = append(errorList, errorMap[code])
	}

	td := tmplData{
		ErrorList: errorList,
	}

	tplate, err := template.New("codetmpl").Parse(tmpl)
	if err != nil {
		log.Fatalf("error parsing template: %s", err)
	}

	var buffer bytes.Buffer
	err = tplate.Execute(&buffer, td)
	if err != nil {
		log.Fatalf("error executing template: %s", err)
	}

	contents, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Fatalf("error formatting generated file: %s", err)
	}

	// contents := buffer.Bytes()

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file (%s): %s", filename, err)
	}

	if _, err := f.Write(contents); err != nil {
		f.Close()
		log.Fatalf("error writing to file (%s): %s", filename, err)
	}

	if err := f.Close(); err != nil {
		log.Fatalf("error closing file (%s): %s", filename, err)
	}

}

func relevantError(err errInfo) bool {
	if err.Product == "Lookup" {
		return true
	}

	if err.Product != "" {
		return false
	}

	if err.LogLevel == "WARNING" {
		return false
	}

	if err.LogType == "" {
		return false
	}

	if strings.HasPrefix(err.Message, `Transcriptions`) {
		return false
	}

	if strings.HasPrefix(err.Message, `Broadcast`) {
		return false
	}

	if strings.Contains(err.Message, `'Template`) {
		return false
	}

	return true
}
