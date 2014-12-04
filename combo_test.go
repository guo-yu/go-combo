package combo

import (
	"fmt"
	"os"
	"testing"
)

var testsJs = []string{
	"./static/js/a.js",
	"./static/js/b.js",
}

var testCss = []string{
	"./static/css/a.css",
	"./static/css/b.css",
}

func TestCombo(t *testing.T) {
	comboStr, err := Combo(testsJs)

	if err != nil {
		t.Error(err)
	}

	t.Log("Passed")
	fmt.Println(comboStr)
}

func TestComboTo(t *testing.T) {
	err := ComboTo("./static/js/dist.js", testsJs)

	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat("./static/js/dist.js"); err == nil {
		t.Log("Passed!")
	} else {
		t.Error("Not found dist files")
	}
}

func TestUriParser(t *testing.T) {
	query, err := parseUri("/?a.js&b.js")

	if err != nil {
		t.Error(err)
		return
	}

	if len(query) == 2 {
		t.Log("Passed!")
	} else {
		t.Error("Not match!")
	}
}
