# character-move

角色移动控制器：贴地吸附、台阶限制、跳跃缓冲与每步 O(1)，只用 Go 标准库。

```
go run ./cmd/character-move --sample=landing
go run ./cmd/character-move --sample=step
go run ./cmd/character-move --sample=edge
go run ./cmd/character-move --sample=buffer
go run ./cmd/character-move --sample=scan
go test ./...
go vet ./...
```

## 口径（README 为准）

- **坐标**：`y` 越小越高，列的地面高度是脚底落地的 `y`；没设置过的列按 0 算。
- **贴地吸附**：下落中越过地面时，脚底必须落在该列地面高度上，垂直速度清零并置为贴地。
- **台阶**：目标列地面比当前脚底高出不超过 `stepHeight` 时可以直接上去；
  高出正好等于 `stepHeight` 也算可上；超过则水平被挡住，本帧不改变所在列。
- **跳跃缓冲**：按下跳跃时记下时刻；只要在 `jumpBufferMs` 内落地，
  这次跳跃必须在落地当帧生效（落地后立刻离地起跳）。
- **不变量**：贴地时脚底等于该列地面高度；同一串输入重复执行结果相同。
- **代价**：`Step` 每步 O(1)，`Scanned` 不随列数线性增长。
- **规模**：10 万列、每秒 10 万步，单步摊还 O(1)。

## 输出契约（不改格式）

```
y=..
x=.. y=..
airborne=..
scanned=..
```

## 目录

```
move.go              贴地、台阶与缓冲
move_test.go         单元测试
cmd/character-move/  命令行入口
```
