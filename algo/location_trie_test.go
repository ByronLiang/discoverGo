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
