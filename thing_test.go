package thing

import "testing"

func TestGiveMeThing(t *testing.T) {
	retval, err := GiveMeThing()
	if (err != nil) || (retval != "thing") {
		t.Fatalf("wanted thing, didn't get it")
	}
}
func TestGiveMeThings(t *testing.T) {
	retval, err := GiveMeThings()
	if (err != nil) || (retval != "things") {
		t.Fatalf("wanted things, didn't get them")
	}
}
