package stats

import (
	"math"
)

func Statistics(xs []float64) (mean, stddev float64) {
	mean = smean(xs)
	variance := svariance(xs)
	stddev = math.Sqrt(variance)
	return
}

func smean(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func svariance(xs []float64) float64 {
	mean := smean(xs)
	sum := 0.0
	for _, x := range xs {
		sum += (x - mean) * (x - mean)
	}
	return sum / float64(len(xs))
}

func Covariance(data1 []float64, data2 []float64) float64 {
	n := len(data1)

	mean1 := smean(data1)
	mean2 := smean(data2)

	covariancen := 0.0
	var a, b float64

	for i := 0; i < n; i++ {
		a = data1[i] - mean1
		b = data2[i] - mean2
		covariancen += a * b
	}

	return covariancen / float64(n)
}

func Correlation(data1 []float64, data2 []float64) float64 {
	cov := Covariance(data1, data2)
	corr := cov / math.Sqrt(svariance(data1)*svariance(data2))
	return corr
}

// Based on algorithm http://home.online.no/~pjacklam/notes/invnorm/
func Norminv(p, m, s float64) float64 {
	a1 := -3.969683028665376e+01
	a2 := 2.209460984245205e+02
	a3 := -2.759285104469687e+02
	a4 := 1.383577518672690e+02
	a5 := -3.066479806614716e+01
	a6 := 2.506628277459239e+00

	b1 := -5.447609879822406e+01
	b2 := 1.615858368580409e+02
	b3 := -1.556989798598866e+02
	b4 := 6.680131188771972e+01
	b5 := -1.328068155288572e+01

	c1 := -7.784894002430293e-03
	c2 := -3.223964580411365e-01
	c3 := -2.400758277161838e+00
	c4 := -2.549732539343734e+00
	c5 := 4.374664141464968e+00
	c6 := 2.938163982698783e+00

	d1 := 7.784695709041462e-03
	d2 := 3.224671290700398e-01
	d3 := 2.445134137142996e+00
	d4 := 3.754408661907416e+00

	p_low := 0.02425
	p_high := 1 - p_low

	var x, q, r float64

	if p > 0 && p < p_low {
		q = math.Sqrt(-2 * math.Log(p))
		x = (((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) / ((((d1*q+d2)*q+d3)*q+d4)*q + 1)
	}

	if p >= p_low && p <= p_high {
		q = p - 0.5
		r = q * q
		x = (((((a1*r+a2)*r+a3)*r+a4)*r+a5)*r + a6) * q / (((((b1*r+b2)*r+b3)*r+b4)*r+b5)*r + 1)
	}

	if p > p_high && p < 1 {
		q = math.Sqrt(-2 * math.Log(1-p))
		x = -(((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) / ((((d1*q+d2)*q+d3)*q+d4)*q + 1)
	}

	return m + x*s
}
