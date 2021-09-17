package main

import (
	"fmt"
	"testing"
)

func TestMathAdd1(t *testing.T) {
	first := 111.0
	second := 222.0
	want := 333.0

	got := add(first, second)

	if got != want {
		t.Fatalf(`Add %v plus %v returned %v. Wanted %v`, first, second, got, want)
	}
}
func TestMathAdd2(t *testing.T) {
	first := 245.0
	second := 0.0
	want := 245.0

	got := add(first, second)

	if got != want {
		t.Fatalf(`Add %v plus %v returned %v. Wanted %v`, first, second, got, want)
	}
}
func TestMathAdd3(t *testing.T) {
	first := 333.0
	second := 123.0
	want := 456.0

	got := add(first, second)

	if got != want {
		t.Fatalf(`Add %v plus %v returned %v. Wanted %v`, first, second, got, want)
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		first, second, want float64
	}{
		{1, 2, 3},
		{4, 5, 9},
		{123, 500, 623},
		{987, 111, 1098},
		{5, 5, 10},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Add:%v,%v", tt.first, tt.second)

		t.Run(testname, func(t *testing.T) {
			got := add(tt.first, tt.second)

			if got != tt.want {
				t.Errorf(`Add %v plus %v should be %v. Got %v`, tt.first, tt.second, tt.want, got)
			}
		})
	}
	t.Run("integration", func(t *testing.T) {
		first := 100.0
		second := 200.0
		want := 30000.0

		got := multiply(add(first, second), subtract(second, first))
		if got != want {
			t.Errorf(`Should be %v. Got %v`, want, got)
		}
	})
}
