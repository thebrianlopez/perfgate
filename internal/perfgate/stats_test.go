package perfgate

import (
	"math"
	"testing"
)

func TestComputeEmpty(t *testing.T) {
	s := Compute(nil)
	if s.N != 0 {
		t.Errorf("expected N=0, got %d", s.N)
	}
}

func TestComputeSingle(t *testing.T) {
	s := Compute([]float64{1.0})
	if s.N != 1 {
		t.Errorf("expected N=1, got %d", s.N)
	}
	if s.Mean != 1.0 {
		t.Errorf("expected Mean=1.0, got %f", s.Mean)
	}
	if s.Median != 1.0 {
		t.Errorf("expected Median=1.0, got %f", s.Median)
	}
	if s.Min != 1.0 || s.Max != 1.0 {
		t.Errorf("expected Min=Max=1.0, got Min=%f Max=%f", s.Min, s.Max)
	}
}

func TestComputeMean(t *testing.T) {
	s := Compute([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	if s.Mean != 3.0 {
		t.Errorf("expected Mean=3.0, got %f", s.Mean)
	}
}

func TestComputeMedianOdd(t *testing.T) {
	s := Compute([]float64{5.0, 1.0, 3.0})
	if s.Median != 3.0 {
		t.Errorf("expected Median=3.0, got %f", s.Median)
	}
}

func TestComputeMedianEven(t *testing.T) {
	s := Compute([]float64{1.0, 2.0, 3.0, 4.0})
	if s.Median != 2.5 {
		t.Errorf("expected Median=2.5, got %f", s.Median)
	}
}

func TestComputeP95(t *testing.T) {
	samples := make([]float64, 100)
	for i := range samples {
		samples[i] = float64(i + 1)
	}
	s := Compute(samples)
	// P95 of 1..100 is the 95th value = 95.0
	if s.P95 != 95.0 {
		t.Errorf("expected P95=95.0, got %f", s.P95)
	}
}

func TestComputeMinMax(t *testing.T) {
	s := Compute([]float64{3.0, 1.0, 4.0, 1.0, 5.0, 9.0})
	if s.Min != 1.0 {
		t.Errorf("expected Min=1.0, got %f", s.Min)
	}
	if s.Max != 9.0 {
		t.Errorf("expected Max=9.0, got %f", s.Max)
	}
}

func TestComputeStdDevKnown(t *testing.T) {
	// stddev of [2, 4, 4, 4, 5, 5, 7, 9] = 2.0 (population)
	s := Compute([]float64{2, 4, 4, 4, 5, 5, 7, 9})
	if math.Abs(s.StdDev-2.0) > 0.0001 {
		t.Errorf("expected StdDev≈2.0, got %f", s.StdDev)
	}
}

func TestComputeStdDevZero(t *testing.T) {
	s := Compute([]float64{5.0, 5.0, 5.0})
	if s.StdDev != 0.0 {
		t.Errorf("expected StdDev=0.0 for constant samples, got %f", s.StdDev)
	}
}
