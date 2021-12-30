package algo

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	//res := KmpSearch("hello", "ll")
	//res := PreKMP("ll")
	//demo := PreKMP("abcdabca")
	demo := PreKMP("aabaabaaa")
	fmt.Println(demo)
	//res := Kmp("hello", "abcdabca")
	res := KmpSearch("abxabcabcaby", "abcaby")
	fmt.Println(res)
}
