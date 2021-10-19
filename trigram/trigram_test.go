package trigram

import "testing"

func TestAddTrigram(t *testing.T) {
	ti := NewTrigramIndex()
	res := ti.Add("Code is my life and i love it")
	if !res {
		t.Errorf("Unable to add to trigram on first try")
	}
	res = ti.Add("Search the whole codebase for bifrost issues")
	if !res {
		t.Errorf("Unable to add to trigram on second try")
	}
	res = ti.Add("I write a lot of Codes to make time go faster")
	if !res {
		t.Errorf("Unable to add to trigram on third try")
	}
}
func TestEmptyLessQuery(t *testing.T) {
	ti := NewTrigramIndex()
	res := ti.Add("Code is my life and i love it, Search the whole codebase for bifrost issues. I write a lot of Codes to make time go faster")
	if !res {
		t.Errorf("Unable to add to trigram on first try")
	}

	ret := ti.Query(5) //less than 5 should still return a structured sentence
	if ret == "" {
		t.Errorf("Error on Empty Response")
	}

	ret = ti.Query(20)
	if len(ret) <= 5 {
		t.Errorf("Error on less than expected character length")
	}
}
