package bloomfilter

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	testcases := []struct{
		Value interface{}
		Expected uint64
	}{
		{"HelloWorld", 17863347649006167745},
	}

	for _, tc := range testcases {
		got := hash(tc.Value)
		if tc.Expected != got {
			t.Errorf("got %d by hash(%v), %d expected", got, tc.Value, tc.Expected)
		}
	}
}