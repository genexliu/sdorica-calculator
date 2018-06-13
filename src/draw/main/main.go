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

	fmt.Printf("expected value: %v\n", calc(*initProb, *probStep, *batch, *charCnt, *maxDraw))

	return
}

// calc returns the expected value
func calc(
	initProb float64, probStep float64, batch uint, charCnt uint, maxDraw uint,
) float64 {
	var hit = initProb
	var prevMiss = 1.0
	var exp = 0.0
	var cnt = uint(1)

	for ; cnt < maxDraw; cnt++ {
		var p = hit * float64(charCnt)
		// the probability is so high that it is almost a guaranteed hit
		if p > threshold {
			break
		}

		exp = exp + float64(cnt)*prevMiss*p

		// update for next loop
		prevMiss = prevMiss * (1.0 - p)
		//fmt.Printf("%v %v %v\n", cnt, p, exp)
		if cnt%batch == 0 {
			hit = hit + probStep
		}
	}
	// plus guaranteed hit
	return exp + float64(cnt)*prevMiss*1.0
}
