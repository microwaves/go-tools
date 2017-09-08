package randomizer

import (
	"reflect"
	"testing"
)

func TestGenerateRandomStringLength(t *testing.T) {
	for n := 1; n <= 100; n++ {
		want := n
		got := GenerateRandomString(want)
		if len(got) != want {
			t.Errorf("Expected '%v' got '%v'", want, got)
		}
	}
}

func TestGenerateRandomStringOutputType(t *testing.T) {
	got := GenerateRandomString(10)
	if reflect.TypeOf(got).Kind() != reflect.String {
		t.Errorf("Expected 'string' got '%v'", got)
	}
}

func TestGenerateRandomStringRandomness(t *testing.T) {
	s := GenerateRandomString(10)
	for n := 0; n <= 10000; n++ {
		nStr := GenerateRandomString(10)
		if s == nStr {
			t.Error("Expected random 'string' got the same")
		}
	}
}
