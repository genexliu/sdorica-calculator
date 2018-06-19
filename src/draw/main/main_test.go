package main

import (
	"math"
	"testing"
)

const (
	numericalError = 0.0001
)

func TestCalc(t *testing.T) {
	assertAppoxEqual(10, calc(0.1, 0, defaultBatch, 1, 1, defaultMaxDraw), t)
	assertAppoxEqual(33.3333, calc(0.03, 0, defaultBatch, 1, 1, defaultMaxDraw), t)

	assertAppoxEqual(454.5454, calc(0.0022, 0, defaultBatch, 1, 1, defaultMaxDraw), t)
	assertAppoxEqual(64.9350, calc(0.0022, 0, defaultBatch, 7, 7, defaultMaxDraw), t)
	assertAppoxEqual(56.8181, calc(0.0022, 0, defaultBatch, 8, 8, defaultMaxDraw), t)

	assertAppoxEqual(55.5555, calc(0.006, 0, defaultBatch, 3, 3, defaultMaxDraw), t)
	assertAppoxEqual(66.6666, calc(0.0075, 0, defaultBatch, 2, 2, defaultMaxDraw), t)

	// 1*0.2 + 2*(1-0.2)*1
	assertAppoxEqual(1.8, calc(0.2, 0.1, 1, 1, 1, 2), t)

	// 1*0.2 + 2*(1-0.2)*0.3 + 3*[(1-0.2)*(1-0.3)]*1
	assertAppoxEqual(2.36, calc(0.2, 0.1, 1, 1, 1, 3), t)

	// 1*0.2 + 2*(1-0.2)*0.3 + 3*[(1-0.2)*(1-0.3)]*0.4 + 4*[(1-0.2)*(1-0.3)*(1-0.4)]*1
	assertAppoxEqual(2.696, calc(0.2, 0.1, 1, 1, 1, 4), t)

	// 1*0.2 + 2*(1-0.2)*0.3 + 3*[(1-0.2)*(1-0.3)]*0.4 + 4*[(1-0.2)*(1-0.3)*(1-0.4)]*0.5 + 5*[(1-0.2)*(1-0.3)*(1-0.4)*(1-0.5)]*1
	assertAppoxEqual(2.864, calc(0.2, 0.1, 1, 1, 1, 5), t)

	// 1*0.2 + 2*(1-0.2)*0.2 + 3*[(1-0.2)*(1-0.2)]*0.3 + 4*[(1-0.2)*(1-0.2)*(1-0.3)]*0.3 + 5*[(1-0.2)*(1-0.2)*(1-0.3)*(1-0.3)]*1
	assertAppoxEqual(3.2016, calc(0.2, 0.1, 2, 1, 1, 5), t)

	// 1*0.2 + 2*(1-0.2)*0.2 + 3*[(1-0.2)*(1-0.2)]*0.2 + 4*[(1-0.2)*(1-0.2)*(1-0.2)]*0.3 + 5*[(1-0.2)*(1-0.2)*(1-0.2)*(1-0.3)]*1
	assertAppoxEqual(3.3104, calc(0.2, 0.1, 3, 1, 1, 5), t)

	// 1*0.4 + 2*(1-0.4)*0.4 + 3*[(1-0.4)*(1-0.4)]*0.6 + 4*[(1-0.4)*(1-0.4)*(1-0.6)]*0.6 + 5*[(1-0.4)*(1-0.4)*(1-0.6)*(1-0.6)]*1
	assertAppoxEqual(2.1616, calc(0.2, 0.1, 2, 2, 2, 5), t)

	// 1*0.4 + 2*(1-0.4)*0.4 + 3*[(1-0.4)*(1-0.4)]*0.4 + 4*[(1-0.4)*(1-0.4)*(1-0.4)]*0.6 + 5*[(1-0.4)*(1-0.4)*(1-0.4)*(1-0.6)]*1
	assertAppoxEqual(2.2624, calc(0.2, 0.1, 3, 2, 2, 5), t)
}

func assertAppoxEqual(a, b float64, t *testing.T) {
	if math.Abs(a-b) > numericalError {
		t.Errorf("%v and %v are not equal", a, b)
	}
}
