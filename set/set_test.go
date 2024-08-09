package set

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	// Ensure the HashSet implements the Set interface.
	var _ Set[int] = &HashSet[int]{}

	// Create a new HashSet.
	set := New[int](10)

	// Test Add and Contains
	set.Add(1)
	if !set.Contains(1) {
		t.Errorf("Set should contain 1 after adding")
	}

	set.Add(1) // Add the same element twice.
	if set.Len() != 1 {
		t.Errorf("Set should not contain duplicate elements")
	}

	// Test Remove
	set.Remove(1)
	if set.Contains(1) {
		t.Errorf("Set should not contain 1 after removing")
	}

	// Test Len
	if set.Len() != 0 {
		t.Errorf("Set length should be 0 after clearing all elements")
	}

	// Test Clear
	set.Add(1)
	set.Add(2)
	set.Clear()
	if set.Len() != 0 {
		t.Errorf("Set length should be 0 after clear")
	}

	// Test Keys
	set.Add(1)
	set.Add(2)
	keys := set.Keys()
	if len(keys) != 2 || (keys[0] != 1 && keys[0] != 2) || (keys[1] != 1 && keys[1] != 2) {
		t.Errorf("Keys should contain 1 and 2")
	}
}
