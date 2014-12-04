package combo

import (
	"fmt"
	"testing"
)

func TestCombo(t *testing.T) {
	comboStr, err := Combo([]string{"./static/js/a.js", "./static/js/b.js"})

	if err != nil {
		t.Error(err)
	}

	t.Log("Passed")
	fmt.Println(comboStr)
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
