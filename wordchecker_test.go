package shamoji

import (
	"testing"

	"github.com/tchap/go-patricia/patricia"
	"golang.org/x/text/unicode/norm"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWordChecker(t *testing.T) {

	Convey("Given form and words", t, func() {

		form := norm.NFC
		words := []string{
			"神",
		}

		Convey("When new wrod checker", func() {

			wc, err := NewWordChecker(form, words)

			So(err, ShouldBeNil)
			So(wc.Trie.Get(patricia.Prefix("神")), ShouldBeTrue)
		})
	})

}
