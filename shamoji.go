package shamoji

import (
	"fmt"
	"sync"

	"github.com/tchap/go-patricia/patricia"
	"golang.org/x/text/unicode/norm"
)

func NewWordChecker(form norm.Form, words []string) (wc *WordChecker, err error) {

	if len(words) == 0 {
		err = fmt.Errorf("may not be zero length")
	}

	wc = &WordChecker{
		Form: form,
		Trie: generateTrie(form, words),
	}
	return
}

func generateTrie(form norm.Form, words []string) (trie *patricia.Trie) {

	var wg sync.WaitGroup

	wg.Add(len(words))

	finChan := make(chan bool)
	go func() {
		wg.Wait()
		finChan <- true
	}()

	trie = patricia.NewTrie()
	wordChan := make(chan []byte, len(words))
	for _, w := range words {
		go func(word string) {
			defer wg.Done()
			wordChan <- form.Bytes([]byte(word))
		}(w)
	}

LOOP:
	for {
		select {
		case b := <-wordChan:
			trie.Set(b, true)
		case <-finChan:
			break LOOP
		}
	}
	return
}
