package gocast

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/Cappta/gofixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBytesAndUInts16(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given any uint16 slice with the length between 10 and 1000", func() {
			uints16 := AnyUints16(AnyIntBetween(10, 1000))
			Convey("When casting from uints16 to bytes", func() {
				data := FromUInts16ToBytes(uints16)
				Convey("When casting back to uint16", func() {
					rebuiltUints16, err := FromBytesToUInts16(data)
					Convey("Then error should be nil", func() {
						So(err, ShouldBeNil)
					})
					Convey("Then uint16 should resemble", func() {
						So(rebuiltUints16, ShouldResemble, uints16)
					})
				})
				Convey("When data is uneven", func() {
					data = append(data, AnyByte())
					Convey("When casting back to uint16", func() {
						rebuiltUints16, err := FromBytesToUInts16(data)
						Convey("Then error should not be nil", func() {
							So(err, ShouldNotBeNil)
						})
						Convey("Then uint16 should be nil", func() {
							So(rebuiltUints16, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}
