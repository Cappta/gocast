package gocast

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/Cappta/gofixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStringAndUInts16(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given any string with the length between 10 and 1000", func() {
			value := AnyString(AnyIntBetween(10, 1000))
			Convey("When casting from string to uint16", func() {
				data := FromStringToUInts16(value)
				Convey("When casting back to value", func() {
					rebuiltValue := FromUInts16ToString(data)
					Convey("Then string value should resemble", func() {
						So(rebuiltValue, ShouldResemble, value)
					})
				})
			})
		})
	})
}
