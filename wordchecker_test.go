package shamoji

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
)

func TestWordCheckerDo(t *testing.T) {

	Convey("Given target", t, func() {
		_, err := ioutil.ReadFile("target_test.txt")
		So(err, ShouldBeNil)
	})

}
