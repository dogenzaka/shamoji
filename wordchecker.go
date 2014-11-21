package shamoji

import (
	"io"

	"github.com/tchap/go-patricia/patricia"
	"golang.org/x/text/unicode/norm"
)

type WordChecker struct {
	Form norm.Form
	Trie *patricia.Trie
}

func (wc *WordChecker) Do(target io.Reader) (ok bool, err error) {
	return
}
