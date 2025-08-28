package utils_test

import (
	"testing"
	"time"

	"github.com/zchrykng/go-datastructures/utils"
)

func TestTimeComparator(t *testing.T) {
	now := time.Now()

	tests := []struct {
		a        time.Time
		b        time.Time
		expected int
	}{
		{now, now, 0},
		{now.Add(24 * 7 * 2 * time.Hour), now, 1},
		{now, now.Add(24 * 7 * 2 * time.Hour), -1},
	}

	for _, test := range tests {
		actual := utils.TimeComparator(test.a, test.b)
		expected := test.expected

		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
