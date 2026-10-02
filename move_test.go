package move

import "testing"

func TestNewStartsAirborne(t *testing.T) {
	mover := New(30, 300, 1000)
	if mover.OnGround() {
		t.Fatal("新控制器应该在空中")
	}
}

func TestDefaultGroundIsZero(t *testing.T) {
	if got := New(30, 300, 1000).GroundOf(7); got != 0 {
		t.Fatalf("没设置过的列地面应该是 0：%v", got)
	}
}

func TestSetColumnIsReadBack(t *testing.T) {
	mover := New(30, 300, 1000)
	mover.SetColumn(2, 40)
	if got := mover.GroundOf(2); got != 40 {
		t.Fatalf("第二列地面应该是 40：%v", got)
	}
}

func TestScannedStartsAtZero(t *testing.T) {
	if New(30, 300, 1000).Scanned() != 0 {
		t.Fatal("Scanned 初始应该是 0")
	}
}

func TestStepReturnsThreeValues(t *testing.T) {
	x, y, onGround := New(30, 300, 1000).Step(0, false, 100)
	if x != 0 || y <= -100 || onGround {
		t.Fatalf("第一帧应该还在下落：%v %v %v", x, y, onGround)
	}
}
