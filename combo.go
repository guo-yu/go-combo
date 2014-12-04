package combo

import "log"
import "errors"
import "strings"
import "net/url"

import . "github.com/tj/go-debug"

import "github.com/turingou/go-fs"
import . "github.com/turingou/go-endswith"

const VERSION = "0.1.0"

var debug = Debug("combo")

func Middleware(uri string) string {
	return ""
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
