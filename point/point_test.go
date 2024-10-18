package point_test

import (
	"testing"

	"github.com/dwethmar/vork/point"
	"github.com/google/go-cmp/cmp"
)

func TestPoint_Add(t *testing.T) {
	expected := point.Point{X: 1, Y: 1}
	if diff := cmp.Diff(expected, (point.Point{}).Add(1, 1)); diff != "" {
		t.Errorf("Add() mismatch (-want +got):\n%s", diff)
	}
}

func TestNew(t *testing.T) {
	expected := point.Point{X: 1, Y: 1}
	if diff := cmp.Diff(expected, point.New(1, 1)); diff != "" {
		t.Errorf("New() mismatch (-want +got):\n%s", diff)
	}
}

func TestZero(t *testing.T) {
	expected := point.Point{X: 0, Y: 0}
	if diff := cmp.Diff(expected, point.Zero()); diff != "" {
		t.Errorf("Zero() mismatch (-want +got):\n%s", diff)
	}
}
