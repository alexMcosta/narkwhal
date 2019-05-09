package main

import "testing"

func TestTimeMeasures(t *testing.T) {
	got := timeMeasures("1d")
	want := "24h"

	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}
