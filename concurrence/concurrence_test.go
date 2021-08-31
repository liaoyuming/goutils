package concurrence

import (
	"errors"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConcurrentExecFunc(t *testing.T) {

	Convey("TestConcurrentExecFunc", t, func() {
		keys := []int{1, 2, 3, 4, 5, 6, 7}

		Convey("TestConcurrentExecFunc success", func() {
			now := time.Now()
			errs := ConcurrentIterate(keys, func(index int, key interface{}) error {
				k, ok := key.(int)
				if !ok {
					t.Fatalf("key type not equals int")
				}
				time.Sleep(100 * time.Millisecond)
				t.Log(index, key, k)
				return nil
			})
			t.Logf("ConcurrentIterate cost: %v", time.Since(now))

			So(len(errs), ShouldEqual, 0)
		})

		Convey("TestConcurrentExecFunc param error", func() {
			errs := ConcurrentIterate(1, func(index int, key interface{}) error {
				return nil
			})
			t.Log(errs)
			So(len(errs), ShouldEqual, 1)
		})

		Convey("TestConcurrentExecFunc failed", func() {
			errs := ConcurrentIterate(keys, func(index int, key interface{}) error {
				return errors.New("test")
			})
			t.Log(errs)
			So(len(errs), ShouldBeGreaterThan, 0)
		})
	})
}

func TestConcurrentExecMultiFunc(t *testing.T) {
	Convey("test", t, func() {
		Convey("success", func() {

			errs := ConcurrentExecMultiFunc(func() error {
				t.Log(1)
				return nil
			}, func() error {
				t.Log(2)
				return nil
			})
			So(errs, ShouldBeNil)
		})

		Convey("error", func() {

			errs := ConcurrentExecMultiFunc(func() error {
				t.Log(3)
				return errors.New("test error")
			}, func() error {
				t.Log(4)
				return nil
			})
			So(len(errs), ShouldEqual, 1)
		})
	})
}
