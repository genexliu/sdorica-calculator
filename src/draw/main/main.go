package main

import (
	"flag"
	"fmt"
)

const (
	defaultBatch   uint    = 10
	defaultMaxDraw uint    = 100000
	threshold      float64 = 0.999
)

var (
	initProb = flag.Float64("init_prob", 0, "initial probability")
	probStep = flag.Float64("prob_step", 0, "the amount of probability increment per ${batch} draws")
	batch    = flag.Uint("batch", defaultBatch, "the amount of draws for probability to increase")
	charCnt  = flag.Uint("char_cnt", 1, "the amount of low probability characters in the pool")
	maxDraw  = flag.Uint("max_draw", defaultMaxDraw, "the amount of draws that guarantees the player a low probability character")
)

func main() {
	flag.Parse()

	fmt.Printf("expected value: %v\n", calc(*initProb, *probStep, *batch, *charCnt, *charCnt, *maxDraw))
	fmt.Printf("expected value for one: %v\n", calc(*initProb, *probStep, *batch, *charCnt, 1, *maxDraw))

	return
}

// calc returns the expected value
func calc(
	initProb float64, probStep float64, batch uint, charCnt uint, targetCnt uint, maxDraw uint,
) float64 {
	var hit = initProb
	var prevMiss = 1.0
	var expNum = 0.0   // numerator
	var expDenom = 1.0 // denominator
	var cnt = uint(1)

	for {
		var p = hit * float64(charCnt)
		if p > threshold {
			break
		}
		if cnt%maxDraw == 0 {
			p = 1.0
			hit = p / float64(charCnt)
		}

		expNum = expNum + float64(cnt)*prevMiss*hit*float64(targetCnt) + float64(cnt)*prevMiss*(p-hit*float64(targetCnt))
		expDenom = expDenom - prevMiss*(p-hit*float64(targetCnt))

		// update for next loop
		prevMiss = prevMiss * (1.0 - p)
		if cnt%batch == 0 {
			hit = hit + probStep
		}
		cnt++
	}
	return expNum / expDenom
}
