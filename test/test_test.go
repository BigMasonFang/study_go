package test

import "testing"

func TestIsValidEmail(t *testing.T) {
	// table driven test
	testCases := []struct {
		email string
		want  bool
	}{
		{"example@example.com", true},
		{"bad@error", false},
		// more test case can be added to the table
	}

	for _, data := range testCases {
		result := IsValidEmail(data.email)
		if result != data.want {
			t.Errorf("%v: %t, want: %t", data.email, result, data.want)
		}
	}
}
