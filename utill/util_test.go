package utill

import (
	"fmt"
	"testing"
)

var data = "<div id='hel' class='123'></div>"

func TestName(t *testing.T) {
	fmt.Println(1)
	doc, err := parse(data)
	fmt.Println("1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data, ok := doc.Find("div").Attr("class")
	fmt.Println(ok)
	fmt.Println(data)
}
