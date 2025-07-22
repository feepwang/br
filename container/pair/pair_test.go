package pair

import "testing"

func TestPair(t *testing.T) {
	// Test creating a pair with different types
	p1 := Pair[int, string]{First: 42, Second: "hello"}
	if p1.First != 42 {
		t.Errorf("Expected First to be 42, got %d", p1.First)
	}
	if p1.Second != "hello" {
		t.Errorf("Expected Second to be 'hello', got %s", p1.Second)
	}

	// Test with same types
	p2 := Pair[int, int]{First: 1, Second: 2}
	if p2.First != 1 || p2.Second != 2 {
		t.Errorf("Expected (1, 2), got (%d, %d)", p2.First, p2.Second)
	}

	// Test zero values
	var p3 Pair[string, bool]
	if p3.First != "" || p3.Second != false {
		t.Errorf("Expected zero values ('', false), got ('%s', %t)", p3.First, p3.Second)
	}
}
