package cloud

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	slice := []string{"a", "b", "c", "z", "a", "b", "c", "d", "e"}
	got := removeDuplicates(slice)
	want := []string{"a", "b", "c", "z", "d", "e"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
