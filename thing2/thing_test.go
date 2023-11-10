package thing2

import "testing"

func TestGiveMeThing2(t *testing.T) {
	retval, err := GiveMeThing2()
	if (err != nil) || (retval != "thing2") {
		t.Fatalf("wanted thing2, didn't get it")
	}
}
func TestGiveMeThings2(t *testing.T) {
	retval, err := GiveMeThings2()
	if (err != nil) || (retval != "things2") {
		t.Fatalf("wanted things2, didn't get them")
	}
}
