// Package move 是角色移动控制器：贴地吸附、台阶限制、跳跃缓冲与每步 O(1)。
// 缺陷：落地不吸附、台阶差一（正好等于步高被挡住）、跳跃缓冲不生效、
// 默认地面取值不统一、每步扫描全部列。
package move

import "math"

// JumpSpeed 是起跳初速度。
const JumpSpeed = 300.0

// Mover 是角色控制器。
type Mover struct {
	stepHeight   float64
	jumpBufferMs float64
	gravity      float64
	columns      map[int]float64
	x            float64
	y            float64
	vy           float64
	onGround     bool
	nowMs        float64
	jumpAtMs     float64
	hasJump      bool
	scanned      int
}

// New 建控制器，角色从地面以上 100 处下落。
func New(stepHeight, jumpBufferMs, gravity float64) *Mover {
	return &Mover{stepHeight: stepHeight, jumpBufferMs: jumpBufferMs, gravity: gravity,
		columns: map[int]float64{}, y: -100}
}

// SetColumn 设置某一列的地面高度（y 越小越高）。
func (m *Mover) SetColumn(column int, groundY float64) { m.columns[column] = groundY }

// GroundOf 读某一列的地面高度，没设置时是 0。
func (m *Mover) GroundOf(column int) float64 { return m.columns[column] }

// Teleport 把角色直接放到指定位置。
func (m *Mover) Teleport(x, y float64, onGround bool) {
	m.x, m.y, m.onGround = x, y, onGround
	m.vy = 0
}

func (m *Mover) groundAt(column int) float64 { return m.columns[column] }

// Step 推进一帧，返回坐标与是否在地面。
// 缺陷：台阶判定用了 <=（正好等于步高被挡）；落地只清速度不吸附坐标；
// 跳跃只在按键当帧且在地面才生效。
func (m *Mover) Step(inputX float64, jump bool, dtMs float64) (float64, float64, bool) {
	m.scanned += len(m.columns)
	m.nowMs += dtMs
	dt := dtMs / 1000.0
	if jump && m.onGround {
		m.vy = -JumpSpeed
		m.onGround = false
	}
	wantX := m.x + inputX*dt
	column := int(math.Round(wantX))
	ground := m.groundAt(column)
	if ground <= m.y-m.stepHeight {
		column = int(math.Round(m.x))
		ground = m.groundAt(column)
	} else {
		m.x = wantX
		if m.onGround {
			m.y = ground
		}
	}
	if !m.onGround {
		m.vy += m.gravity * dt
		m.y += m.vy * dt
		if m.y >= ground {
			m.onGround = true
			m.vy = 0
		}
	}
	return m.x, m.y, m.onGround
}

// OnGround 是当前是否贴地。
func (m *Mover) OnGround() bool { return m.onGround }

// Scanned 是累计扫描量（规模观测）。
func (m *Mover) Scanned() int { return m.scanned }
