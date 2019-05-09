package main

import "testing"

func TestTimeMeasures(t *testing.T) {

	t.Run("test with days", func(t *testing.T) {
		got := timeMeasures("1d")
		want := "24h"

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("test with weeks", func(t *testing.T) {
		got := timeMeasures("1w")
		want := "168h"

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("test with miliseconds", func(t *testing.T) {
		got := timeMeasures("1ms")
		want := "1ms"

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
