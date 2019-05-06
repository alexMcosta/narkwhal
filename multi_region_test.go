package main

import (
	"reflect"
	"testing"
)

func TestMultiRegion(t *testing.T) {

	checkSums := func(t *testing.T, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("processes one region", func(t *testing.T) {
		regions := "us-west-1"
		got := multiRegion(regions)
		want := []string{
			"us-west-1",
		}

		checkSums(t, got, want)
	})

	t.Run("processes multiple regions", func(t *testing.T) {
		regions := "us-west-1,us-east-2,ca-central-1"
		got := multiRegion(regions)
		want := []string{
			"us-west-1",
			"us-east-2",
			"ca-central-1",
		}

		checkSums(t, got, want)
	})

	t.Run("processes the EU flag", func(t *testing.T) {
		regions := "EU"
		got := multiRegion(regions)
		want := []string{
			"eu-central-1",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"eu-north-1",
		}

		checkSums(t, got, want)
	})

	t.Run("processes the AP flag", func(t *testing.T) {
		regions := "AP"
		got := multiRegion(regions)
		want := []string{
			"ap-south-1",
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-southeast-1",
			"ap-southeast-2",
		}

		checkSums(t, got, want)
	})

	t.Run("processes the AM flag", func(t *testing.T) {
		regions := "AM"
		got := multiRegion(regions)
		want := []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"ca-central-1",
			"sa-east-1",
		}

		checkSums(t, got, want)
	})

	t.Run("can process a group flag with one region", func(t *testing.T) {
		regions := "AM,eu-west-3"
		got := multiRegion(regions)
		want := []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"ca-central-1",
			"sa-east-1",
			"eu-west-3",
		}

		checkSums(t, got, want)
	})

	t.Run("can process a group flag with multiple regions", func(t *testing.T) {
		regions := "AP,eu-west-3,us-west-1"
		got := multiRegion(regions)
		want := []string{
			"ap-south-1",
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-southeast-1",
			"ap-southeast-2",
			"eu-west-3",
			"us-west-1",
		}

		checkSums(t, got, want)
	})

	t.Run("can process the ALL flag", func(t *testing.T) {
		regions := "ALL"
		got := multiRegion(regions)
		want := []string{
			"ap-south-1",
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-southeast-1",
			"ap-southeast-2",
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"ca-central-1",
			"sa-east-1",
			"eu-central-1",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"eu-north-1",
		}

		checkSums(t, got, want)
	})
}

func BenchmarkMultiRegion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regions := "ALL"
		multiRegion(regions)
	}
}
