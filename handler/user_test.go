package handler

import (
	"fmt"
	"testing"
)

func TestTri(t *testing.T) {
	test := []struct {
		username string
		token    string
		ans      bool
	}{
		{"xhw", "7fcae0199b0f7825c2471de99a30798f65d486ff", false},
		{"xhw", "5191e7b359058cc22c858774dd7f7caa65d48761", false},
	}
	test = append(test, struct {
		username string
		token    string
		ans      bool
	}{username: "xhw", token: GenToken("xhw"), ans: true})

	fmt.Println(test)
	for index, tt := range test {
		if isVaild := IsTokenVaild(tt.username, tt.token); isVaild != tt.ans {
			t.Errorf("实例%d验证失败", index)
		}
	}
}
