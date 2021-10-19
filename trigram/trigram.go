package trigram

import (
	"math/rand"
	"strings"
)

//go:generate mockgen -source=trigram.go -destination=../mocks/trigram_mock.go -package=mocks

type Trigram string

type TrigramIntf interface {
	Add(doc string) bool
	Query(complexity int) string
	ExtractStringToTrigram(str string) bool
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Extract three string to trigram list
func (t *TrigramIndex) ExtractStringToTrigram(str string) bool {
	if len(str) == 0 {
		return false
	}
	a := []string{}
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		a = append(a, words[i])
		if len(a) == 3 {
			if len(t.Docs) == 0 {
				t.Docs = append(t.Docs, a)
				a = nil
				continue
			}
			for _, v := range t.Docs {
				if !Equal(v, a) && len(a) != 0 {
					t.Docs = append(t.Docs, a)
					a = nil
				}
			}
		}
	}
	return true
}

type TrigramIndex struct {
	//Save all trigram outcomes to a doc
	Docs [][]string
}

//Create a new trigram indexing
func NewTrigramIndex() TrigramIntf {
	t := TrigramIndex{
		Docs: make([][]string, 0),
	}
	return &t
}

//Add new document into this trigram index
func (t *TrigramIndex) Add(doc string) bool {
	return t.ExtractStringToTrigram(doc)
}

//Query a target string to return a sentence
func (t *TrigramIndex) Query(complexity int) string {
	if len(t.Docs) == 0 {
		return ""
	}

	var newWrd string
	for i := 0; i < complexity; i++ {
		randNum := rand.Intn(len(t.Docs))
		if newWrd != "" {
			newWrd = newWrd + " " + t.Docs[randNum][0] + " " + t.Docs[randNum][1] + " " + t.Docs[randNum][2]
		} else {
			newWrd = t.Docs[randNum][0] + " " + t.Docs[randNum][1] + " " + t.Docs[randNum][2]
		}
	}

	return newWrd

}
