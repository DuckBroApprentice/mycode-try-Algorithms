package classtwo

import "testing"

func TestJosephus(t *testing.T) {
	type testData struct {
		n    int
		want int
	}
	var testArray []testData = []testData{
		{n: 10, want: 5},
		{n: 5, want: 3},
		{n: 2, want: 1},
	}
	for _, tt := range testArray {
		res := Josephus(tt.n)
		if res != tt.want {
			t.Errorf("input n: %d, output want: %d", tt.n, tt.want)
		}
	}
}
