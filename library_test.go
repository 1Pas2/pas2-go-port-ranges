package pas2

import (
	"reflect"
	"testing"
)

func TestRanges(t *testing.T) {
	ports, e := Parse("443, 80, 8000-8002,80")
	if e != nil || !reflect.DeepEqual(ports, []int{80, 443, 8000, 8001, 8002}) {
		t.Fatal(ports, e)
	}
}
func TestInvalid(t *testing.T) {
	for _, s := range []string{"0", "65536", "3-1", "80,", "a", "1-2-3", "+80"} {
		if _, e := Parse(s); e == nil {
			t.Fatal(s)
		}
	}
}
func TestEmpty(t *testing.T) {
	ports, e := Parse("")
	if e != nil || len(ports) != 0 {
		t.Fatal(ports, e)
	}
}
