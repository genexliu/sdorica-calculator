# Sdorica Calculator

## Usage

```
# read usage
go run src/draw/main/main.go -h

# 羈絆賦魂連續單抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.0075 -char_cnt 2 -target_cnt 2
# output:
# 抽出任一個想要角色的抽數期望值: 66.66666666666599
# 抽出特定角色的抽數期望值: 133.33333333333192
# 抽出所有想要角色的抽數期望值: 199.9999999999979

# 羈絆賦魂連續十抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.0075 -char_cnt 2 -target_cnt 2 -prob_step 0.0015 -max_draw 120
# output:
# 抽出任一個想要角色的抽數期望值: 42.34676030503299
# 抽出特定角色的抽數期望值: 84.69352061006583
# 抽出所有想要角色的抽數期望值: 127.04028091509882

# 起始機率 0.006, 十連抽不中每隻加成 0.0012, 保底120抽, 活動角色共3隻但是只想要其中2隻
# 連續十抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.006 -char_cnt 3 -target_cnt 2 -prob_step 0.0012 -max_draw 120
# output:
# 抽出任一個想要角色的抽數期望值: 55.8791508117972
# 抽出特定角色的抽數期望值: 111.7583016235944
# 抽出所有想要角色的抽數期望值: 167.6374524353916
```

## Test

```
go test draw/main
```

## Prerequisite

- go command
