package thing2

import "testing"

func TestGiveMeThing2(t *testing.T) {
	retval, err := GiveMeThing2()
	if (err != nil) || (retval != "thing") {
		t.Fatalf("wanted thing, didn't get it")
	}
}
func TestGiveMeThings2(t *testing.T) {
	retval, err := GiveMeThings2()
	if (err != nil) || (retval != "things") {
		t.Fatalf("wanted things, didn't get them")
	}
}
