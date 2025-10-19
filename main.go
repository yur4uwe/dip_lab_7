package main

import (
	"graph"
	"math"
	"math/cmplx"
)

func f(t ...float64) []float64 {
	res := make([]float64, len(t))
	for i := range t {
		res[i] = math.Cos((2*math.Pi*t[i])/10+1) + math.Cos((2*math.Pi*t[i])/40+math.Pi/2)
	}
	return res
}

func dft(x []complex128, inv bool) []complex128 {
	N := len(x)
	X := make([]complex128, N)

	for k := 0; k < N; k++ {
		var sum complex128
		for n := 0; n < N; n++ {
			omega := 2 * math.Pi * float64(k*n) / float64(N)
			if inv {
				omega = -omega
			}
			sum += complex(math.Cos(omega), math.Sin(omega)) * x[n]
		}
		if inv {
			sum /= complex(float64(N), 0)
		}
		X[k] = sum
	}

	return X
}

func amplitudeSpectrum(X []complex128) []float64 {
	A := make([]float64, len(X)/2)
	for k := range A {
		A[k] = cmplx.Abs(X[k])
	}
	return A
}

func phaseSpectrum(X []complex128) []float64 {
	Ph := make([]float64, len(X)/2)
	for k := range Ph {
		Ph[k] = math.Atan2(imag(X[k]), real(X[k]))
	}
	return Ph
}

func realPart(X []complex128) []float64 {
	Re := make([]float64, len(X))
	for i := range X {
		Re[i] = real(X[i])
	}
	return Re
}

func complexSlice(x []float64) []complex128 {
	X := make([]complex128, len(x))
	for i := range x {
		X[i] = complex(x[i], 0)
	}
	return X
}

func main() {
	t := graph.LinearArray(0, 480, 400)

	y := f(t...)

	X := dft(complexSlice(y), false)
	XInv := dft(X, true)

	g := graph.NewGraph()
	ls := graph.NewLS()
	ls.Solid()
	invLs := graph.NewLS()
	invLs.Dots(3)

	g.Plot(t, y, ls)
	g.Plot(t, realPart(XInv), invLs)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/dft_plot.png"); err != nil {
		panic(err)
	}

	g.Clear()

	freq_ls := graph.NewLS()
	freq_ls.Pillars(3)

	amplitude := amplitudeSpectrum(X)
	freqs := graph.IntLinearArray(0, len(amplitude)-1)

	g.Plot(freqs, amplitude, freq_ls)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/dft_freq_domain_plot.png"); err != nil {
		panic(err)
	}

	g.Clear()

	phase := phaseSpectrum(X)
	g.Plot(freqs, phase, ls)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/dft_phase_domain_plot.png"); err != nil {
		panic(err)
	}
}
