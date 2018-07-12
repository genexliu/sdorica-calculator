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
	initProb  = flag.Float64("init_prob", 0, "低機率角色初始機率")
	probStep  = flag.Float64("prob_step", 0, "${batch}連抽不中機率加成")
	batch     = flag.Uint("batch", defaultBatch, "連抽次數")
	charCnt   = flag.Uint("char_cnt", 1, "低機率角色數量")
	targetCnt = flag.Uint("target_cnt", 1, "低機率角色中想要的角色數量")
	maxDraw   = flag.Uint("max_draw", defaultMaxDraw, "保底抽數")
)

func main() {
	flag.Parse()

	fmt.Printf("抽出任一個想要角色的抽數期望值: %.2f\n", calc(*initProb, *probStep, *batch, *charCnt, *targetCnt, *maxDraw))
	fmt.Printf("抽出特定角色的抽數期望值: %.2f\n", calc(*initProb, *probStep, *batch, *charCnt, 1, *maxDraw))

	var total = 0.0
	for t := *targetCnt; t >= 1; t-- {
		total = total + calc(*initProb, *probStep, *batch, *charCnt, t, *maxDraw)
	}
	fmt.Printf("抽出所有想要角色的抽數期望值: %.2f\n", total)
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
