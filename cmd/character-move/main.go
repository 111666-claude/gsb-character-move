// Command character-move 跑角色移动样例。
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"example.com/move"
)

// Run 执行一次命令行调用，返回退出码。
func Run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("character-move", flag.ContinueOnError)
	flags.SetOutput(stderr)
	sample := flags.String("sample", "landing", "landing / step / ledge / buffer / scan")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	switch *sample {
	case "landing":
		mover := move.New(30, 300, 1000)
		mover.SetColumn(0, 0)
		_, y, _ := mover.Step(0, false, 1000)
		fmt.Fprintf(stdout, "y=%.0f\n", y)
	case "step":
		mover := move.New(30, 300, 1000)
		mover.SetColumn(0, 0)
		mover.SetColumn(1, -30)
		mover.Teleport(0, 0, true)
		x, y, _ := mover.Step(1, false, 1000)
		fmt.Fprintf(stdout, "x=%.0f y=%.0f\n", x, y)
	case "ledge":
		mover := move.New(30, 300, 1000)
		mover.SetColumn(0, 0)
		mover.SetColumn(1, 100)
		mover.Teleport(0, 0, true)
		_, y, onGround := mover.Step(10, false, 100)
		fmt.Fprintf(stdout, "y=%.0f airborne=%s\n", y, map[bool]string{true: "false", false: "true"}[onGround])
	case "buffer":
		mover := move.New(30, 300, 1000)
		mover.SetColumn(0, 0)
		mover.Step(0, true, 100)
		mover.Step(0, false, 100)
		mover.Step(0, false, 100)
		_, _, onGround := mover.Step(0, false, 100)
		fmt.Fprintf(stdout, "airborne=%s\n", map[bool]string{true: "false", false: "true"}[onGround])
	case "scan":
		mover := move.New(30, 300, 1000)
		for index := 0; index < 3000; index++ {
			mover.SetColumn(index, 0)
		}
		for index := 0; index < 3000; index++ {
			mover.Step(0, false, 1)
		}
		fmt.Fprintf(stdout, "scanned=%d\n", mover.Scanned())
	default:
		fmt.Fprintln(stderr, "需要 --sample landing|step|ledge|buffer|scan")
		return 2
	}
	return 0
}

func main() {
	os.Exit(Run(os.Args[1:], os.Stdout, os.Stderr))
}
