package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

// JSON type structure
type JSON map[string]interface{}

func main() {
	funcMap := template.FuncMap{
		"month":  formatAsMonth,
		"year":   formatAsYear,
		"toList": formatArrayAsList,
	}

	const jsonFile = "data.json"
	jsondata, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	const templateFile = "template.tex"
	t, err := template.New(templateFile).Funcs(funcMap).Delims("[[", "]]").ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	resume := JSON{}
	if err := json.Unmarshal(jsondata, &resume); err != nil {
		panic(err)
	}

	if _, err := os.Stat("out"); os.IsNotExist(err) {
		os.Mkdir("out", 0755)
	}

	file, err := os.Create("out/resume.tex")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	if err := t.Execute(file, resume); err != nil {
		panic(err)
	}

	if err := exec.Command("pdflatex", "-output-directory", "out", "out/resume.tex").Run(); err != nil {
		log.Fatal("Cannot create pdf", err)
	}
}

func formatAsMonth(date string) string {
	return formatAsDate(date, "Jan")
}

func formatAsYear(date string) string {
	return formatAsDate(date, "2006")
}

func formatAsDate(date string, format string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format(format)
}

func formatArrayAsList(elements []interface{}) string {
	list := make([]string, len(elements))
	for i, v := range elements {
		list[i] = v.(string)
	}
	return strings.Join(list, ", ") + "."
}
