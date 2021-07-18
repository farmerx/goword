package tests

import (
	"fmt"
	"goword"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := goword.OpenFile("./files/basic.docx")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", f)
}
