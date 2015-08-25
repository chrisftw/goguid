package goguid

import (
	"testing"
)

func TestGenV4(t *testing.T) {
	// testing 100 of each test with random values to make sure it passes for many random
	for i := 0; i < 100; i++ {
		guid := GenV4()
		AssertTrue(len(guid) == 36, "The length is not correct", t)
		AssertTrue(guid[14:15] == "4", "The 15th character should be a 4.", t)
		AssertTrue(guid[19:20] == "8" || guid[19:20] == "9" || guid[19:20] == "a" || guid[19:20] == "b", "The 19th character should be a 8, 9, a, or b.", t)
	}
}

func BenchmarkGenV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenV4()
	}
}

func AssertTrue(value bool, message string, t *testing.T) {
	if !value {
		t.Log(message)
		t.Fail()
	}
}
