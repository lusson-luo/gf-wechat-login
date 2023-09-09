package geo_test

import (
	"fmt"
	"login-demo/internal/packed/geo"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestGetDistance(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		location1 := "39.90469,116.40717"
		location2 := "31.23037,121.47370"

		distance := geo.GetDistance(location1, location2)
		fmt.Printf("两地之间的直线距离为：%.2f千米\n", distance)
		t.Assert(1067, int(distance))
	})
}
