package concurrence

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoWithDefaultRecovery(t *testing.T) {
	type args struct {
		f func()
	}

	Convey("test", t, func() {
		Convey("testdata case1", func() {
			done := make(chan bool)

			SetPanicLogFunc(func(format string, v ...interface{}) {
				t.Log(fmt.Sprintf(format, v...))
				done <- true
			})

			params := args{
				f: func() {
					panic("test")
				},
			}
			GoWithDefaultRecovery(params.f)
			So(<-done, ShouldBeTrue)
		})
	})
}
