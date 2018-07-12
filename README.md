# Sdorica Calculator

## Usage

```
# read usage
go run src/draw/main/main.go -h

# 午茶賦魂連續十抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.01 -char_cnt 1 -target_cnt 1 -prob_step 0.002 -max_draw 120
# output:
# 抽出任一個想要角色的抽數期望值: 54.89
# 抽出特定角色的抽數期望值: 54.89
# 抽出所有想要角色的抽數期望值: 54.89

# 羈絆賦魂連續單抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.0075 -char_cnt 2 -target_cnt 2
# output:
# 抽出任一個想要角色的抽數期望值: 66.67
# 抽出特定角色的抽數期望值: 133.33
# 抽出所有想要角色的抽數期望值: 200.00

# 羈絆賦魂連續十抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.0075 -char_cnt 2 -target_cnt 2 -prob_step 0.0015 -max_draw 120
# output:
# 抽出任一個想要角色的抽數期望值: 42.35
# 抽出特定角色的抽數期望值: 84.69
# 抽出所有想要角色的抽數期望值: 127.04

# 起始機率 0.006, 十連抽不中每隻加成 0.0012, 保底120抽, 活動角色共3隻但是只想要其中2隻
# 連續十抽, 抽數期望值分析
go run src/draw/main/main.go -init_prob 0.006 -char_cnt 3 -target_cnt 2 -prob_step 0.0012 -max_draw 120
# output:
# 抽出任一個想要角色的抽數期望值: 55.88
# 抽出特定角色的抽數期望值: 111.76
# 抽出所有想要角色的抽數期望值: 167.64
```

## Test

```
go test draw/main
```

## Prerequisite

- go command
