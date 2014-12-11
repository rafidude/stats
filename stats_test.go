package stats

import (
	"math"
	"testing"
)

type testpair struct {
	values   []float64
	mean     float64
	variance float64
	stddev   float64
}

var stests = []testpair{
	{[]float64{1, 2, 3, 4}, 2.5, 1.25, 1.11803398875},
	{[]float64{1, 2}, 1.5, 0.25, 0.5},
	{[]float64{1, 1, 0, 1, 2}, 1, 0.4, 0.63245553203368},
	{[]float64{-1, 1}, 0, 1, 1},
}

func TestStatistics(t *testing.T) {
	tol := 0.00000001
	for _, input := range stests {
		mean, stddev := Statistics(input.values)
		if math.Abs(mean-input.mean) > tol {
			t.Error(
				"For", input.values,
				"expected", input.mean,
				"got", mean,
			)
		}
		if math.Abs(stddev-input.stddev) > tol {
			t.Error(
				"For", input.values,
				"expected", input.stddev,
				"got", stddev,
			)
		}
	}
}

type normi struct {
	prob   float64
	zscore float64
}

// tests from http://www.danielsoper.com/statcalc3/calc.aspx
var ztests = []normi{
	{0.84134, 0.99998039},
	{0.05, -1.64485363},
	{0.1, -1.28155156},
	{0.2, -0.84162123},
	{0.3, -0.52440051},
	{0.4, -0.25334710},
	{0.5, -0.0},
	{0.6, 0.25334710},
	{0.7, 0.52440051},
	{0.8, 0.84162123},
	{0.9, 1.28155157},
	{0.95, 1.64485363},
}

func TestNorminv(t *testing.T) {
	tol := 0.000001
	for _, input := range ztests {
		zscoreAns := Norminv(input.prob, 0, 1)
		if math.Abs(zscoreAns-input.zscore) > tol {
			t.Error(
				"For", input.prob,
				"expected", input.zscore,
				"got", zscoreAns,
			)
		}
	}
}

type covpair struct {
	data1       []float64
	data2       []float64
	covariance  float64
	correlation float64
}

var ctests = []covpair{
	{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 4}, 1.25, 1.00},
	{[]float64{1.1, 2, 3.3, 4}, []float64{1, 2.2, 3, 4.4}, 1.36, 0.978369794553563},
	{[]float64{215, 325, 185, 332, 406, 522, 412, 614, 544, 421, 445, 408}, []float64{14.2, 16.4, 11.9, 15.2, 18.5, 22.1, 19.4, 25.1, 23.4, 18.1, 22.6, 17.2}, 443.7520833333335, 0.9575066230015952},
}

func TestCovarianceCorrelation(t *testing.T) {
	tol := 0.00000001
	for _, input := range ctests {
		covariance := Covariance(input.data1, input.data2)
		if math.Abs(covariance-input.covariance) > tol {
			t.Error(
				"For", input.data1,
				"expected", input.covariance,
				"got", covariance,
			)
		}
		correlation := Correlation(input.data1, input.data2)
		if math.Abs(correlation-input.correlation) > tol {
			t.Error(
				"For", input.data1,
				"expected", input.correlation,
				"got", correlation,
			)
		}
	}
}
