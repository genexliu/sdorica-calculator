# Sdorica Calculator

## Prerequisite

- go command

## How to use

```
# read usage
go run src/draw/main/main.go -h

# 羈絆賦魂連續單抽, 抽中第一隻SP的抽數期望值
go run src/draw/main/main.go -init_prob 0.0075 -char_cnt 2
# output:
# expected value: 66.66666666666599

# 羈絆賦魂連續10抽, 抽中第一隻SP的抽數期望值
go run src/draw/main/main.go -init_prob 0.0075 -char_cnt 2 -prob_step 0.0015 -max_draw 120
# output:
# expected value: 42.34676030503299
```

## Test

```
go test draw/main
```
