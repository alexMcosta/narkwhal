package main

import (
	"reflect"
	"testing"
)

func TestMultiAccount(t *testing.T) {
	got := multiAccount("Default,Narkwhal")
	want := []string{"Default", "Narkwhal"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
