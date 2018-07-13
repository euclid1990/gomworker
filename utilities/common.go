package utilities

import (
	"bytes"
	"github.com/euclid1990/gomworker/configs"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func LoadEnv(file string) {
	if file == "" {
		file = ".env"
	}
	err := godotenv.Load(file)
	if err != nil {
		Logf(configs.LOG_CRITICAL, "Error loading %v file", file)
	}
}

func ParseStringTemplate(content string, data interface{}) (string, error) {
	t := template.Must(template.New("").Parse(content))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		Logf(configs.LOG_CRITICAL, "Error when parsing template %v", content)
		return "", err
	}
	return buf.String(), err
}

func GetFileContent(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(content), nil
}

func NotBlankLine(s string) bool {
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, " ")
	return len(s) > 0
}

// Support wildcard
func RemoveFiles(path string) {
	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}
