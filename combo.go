package combo

import (
	"errors"
	"log"
	"net/url"
	"strings"

	. "github.com/tj/go-debug"
	. "github.com/turingou/go-endswith"
	"github.com/turingou/go-fs"
)

const VERSION = "0.1.0"

var debug = Debug("combo")

func Middleware(uri string) string {
	return ""
}

func ComboTo(dist string, files []string) error {
	distStr, err := Combo(files)

	if err != nil {
		return err
	}

	return fs.WriteFile(dist, distStr)
}

func Combo(files []string) (string, error) {
	debug("%s", files)

	if len(files) == 0 {
		return "", errors.New("Not enough files to combo")
	}

	content := []string{}

	for _, file := range files {
		fileStr, err := fs.ReadFile(file)

		if err != nil {
			return "", err
		}

		content = append(content, fileStr)
	}

	return strings.Join(content, ""), nil
}

func parseUri(uri string) ([]string, error) {
	u, err := url.Parse(uri)
	files := []string{}

	if err != nil {
		log.Fatal(err)
		return files, err
	}

	query := u.Query()

	if len(query) == 0 {
		return files, errors.New("invalid query string")
	}

	// Ingore the staic `version` in this version
	for filename, _ := range query {
		if EndsWithIn(filename, []string{".css", ".js"}) {
			files = append(files, filename)
		}
	}

	return files, nil
}
