package SkyLine_ExternalLibraries_Quantum

import (
	"fmt"
	"math"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// SkyLine Quantum is a currently beta module of the SkyLine programming language that implements quantum concepts
// that can be respectively used within the cyber security space such as Grovers algorithm, QKD and other various algorithms
// Code contained in the library written by the developers (Totally_Not_A_Haxxer) SHOULD NOT be used in any direct production ready software
// the developers of the skyline programming language do not recomend these due to them not being fully tested and currently being in beta.
//
//
// [WARN] -> Any use case of Quantum mathematics should be tested further and in the future case's of skyline you should use Qiskit with a plugin
// [WARN] -> or any other library developed for go that can be plugged in. Without further to say, below you will find a list of current algorithms
// [WARN] -> that are un-stable and are working demos.
//
// Grovers search algorithm
// QKD ( Quantum Key Distribution)
// ...
//
//
// ------- DEVELOPER NOTES ---------
//
// It is important that when working with functions they all end with M or S where M stands for the module number
// in this case the Module number is the number of the file in the order that the file was generated and S stands for the function
// number within the current file.
//

type OracleFunction func(int) bool

func Grovers_Search_Algorithm_M1_S0(OC OracleFunction, n int) int {
	Itter := int(math.Round(math.Pi / 4 * math.Sqrt(float64(n))))
	state := make([]float64, n)
	for idx := 0; idx < n; idx++ {
		state[idx] = 1 / math.Sqrt(float64(n))
	}

	for l := 0; l < Itter; l++ {
		for j := 0; j < n; j++ {
			if OC(j) {
				state[j] *= -1
			}
		}
		avg := 0.0
		for _, amplitude := range state {
			avg += amplitude / float64(n)
		}
		for k := 0; k < n; k++ {
			state[k] = 2*avg - state[k]
		}
	}
	maxampl := 0.0
	idx := 0
	for i, amplitude := range state {
		if amplitude > maxampl {
			maxampl = amplitude
			idx = i
		}
	}
	return idx
}

func InitateGrover() {
	oracle := func(x int) bool {
		spec := []int{1, 4, 6, 7, 9}
		return spec[x] == 7
	}
	idx := Grovers_Search_Algorithm_M1_S0(oracle, 5)
	fmt.Println(idx)
}
