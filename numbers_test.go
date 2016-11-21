package numbers

import (
	"testing"
)

func TestAddDelim(t *testing.T) {
	s := AddDelimiters(32.3)
	if s != "32.300" {
		t.Errorf("expected 32.3, not '%s'", s)
	}

	s = AddDelimiters(3032.3)
	if s != "3,032" {
		t.Errorf("expected 3,032.3000, not '%s'", s)
	}

	s = AddDelimiters(123131231233032.3)
	if s != "123,131,231,233,032" {
		t.Errorf("expected 3,032.3000, not '%s'", s)
	}

	s = AddDelimiters(1123131231233032.3)
	if s != "1,123,131,231,233,032" {
		t.Errorf("expected 3,032.300, not '%s'", s)
	}
	s = AddDelimiters(12123131231233032.3)
	if s != "12,123,131,231,233,032" {
		t.Errorf("expected 3,032.300, not '%s'", s)
	}

	s = AddDelimitersInt(12123131231233032)
	if s != "12,123,131,231,233,032" {
		t.Errorf("expected 3,032.300, not '%s'", s)
	}
}

type dtest struct {
	num    float64
	result string
}

var dispTests = []dtest{{999.0, "999"},
	{0.3145, "0.315"},
	{1300000, "1.30M"},
	{871300000, "871.30M"},
	{1300000000, "1.30G"},
	{13000000000, "13.00G"},
	{1300000000000, "1.30T"},
	{1300000000000000, "1.30Q"},
	{0.0, "0"},
	{0.0000, "0"},
	{1001.0, "1,001"}}

func TestDisplay(t *testing.T) {
	for i, v := range dispTests {
		s := Display(v.num)
		if s != v.result {
			t.Errorf("%d: expected '%s', got '%s'", i, v.result, s)
		}
	}
}

func TestPercentage(t *testing.T) {
	if Percentage(100.0, 0.0) != 0.0 {
		t.Errorf("expected 0, got %f", Percentage(100.0, 0.0))
	}
	if PercentageMid(100.0, 0.0) != 200.0 {
		t.Errorf("expected 200%%, got %f", PercentageMid(100.0, 0.0))
	}
	if Percentage(110.0, 100.0) != 10.0 {
		t.Errorf("expected 10%%, got %f", Percentage(110.0, 100.0))
	}
	if PercentageMid(110.0, 100.0) != 9.523809523809524 {
		t.Errorf("expected 10%%, got %f", PercentageMid(110.0, 100.0))
	}
}

func TestScale(t *testing.T) {
	if Scale(1.23, 1) != 2.0 {
		t.Errorf("expected 2.0, got %f", Scale(1.23, 1))
	}

	if ScaleDown(1.23, 1) != 1.0 {
		t.Errorf("down scale to 1.0, got %f", ScaleDown(1.23, 1))
	}
	if ScaleDown(2.9, 1) != 2.0 {
		t.Errorf("down scale to 2.0, got %f", ScaleDown(2.9, 1))
	}
	if ScaleDown(2934.0, 1) != 2000.0 {
		t.Errorf("down scale to 2000.0, got %f", ScaleDown(2934.0, 1))
	}
	if Scale(2934.0, 1) != 3000.0 {
		t.Errorf("up scale to 3000.0, got %f", Scale(2934.0, 1))
	}
	if Scale(-2.3, 1) != -3.0 {
		t.Errorf("expected -3.0, got %f", Scale(-2.3, 1))
	}
	if Scale(0.25, 1) != 0.3 {
		t.Errorf("expected 0.3, got %f", Scale(0.25, 1))
	}
	if Scale(0.00003, 1) != 0.00003 {
		t.Errorf("expected 0.00003, got %f", Scale(0.00003, 1))
	}
	if Scale(0.000023, 1) != 0.00003 {
		t.Errorf("expected 0.00003, got %f", Scale(0.000023, 1))
	}
	if Scale(-0.000023, 1) != -0.00003 {
		t.Errorf("expected -0.00003, got %f", Scale(-0.000023, 1))
	}
	if Scale(6210.0, 2) != 6300.0 {
		t.Errorf("Scale(6210, 2) == %v, expected 6300.0", Scale(6210.0, 2))
	}
	if ScaleDown(6210.0, 2) != 6200.0 {
		t.Errorf("ScaleDown(6210, 2) == %v, expected 6200.0", ScaleDown(6210.0, 2))
	}
	if Scale(0.251, 2) != 0.26 {
		t.Errorf("scale 0.251, 2 = %v, expected 0.26", Scale(0.251, 2))
	}
}
