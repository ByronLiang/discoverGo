package algo

import "testing"

func TestInitLocationTrie(t *testing.T) {
	cit, cct, err := InitLocationTrie("./data/location.json")
	if err != nil {
		t.Fatal(err)
	}
	cid, ok := cct.ParseCountryCodeToId("JPN")
	if !ok {
		t.Log("can not find")
		return
	}
	t.Log("JPN", cid)
	code, ok := cit.ParseCountryIdToCode(cid)
	if ok {
		t.Log(code)
	}
}

func BenchmarkCountryCodeTrie_ParseCountryCodeToId(b *testing.B) {
	_, cct, err := InitLocationTrie("./data/location.json")
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		cct.ParseCountryCodeToId("JPN")
	}
}

func BenchmarkCountryIdTrie_ParseCountryIdToCode(b *testing.B) {
	cid, _, err := InitLocationTrie("./data/location.json")
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		cid.ParseCountryIdToCode(11980)
	}
}
