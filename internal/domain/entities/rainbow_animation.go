package entities

import "time"

// RainbowAnimation represents the state of rainbow color animation
type RainbowAnimation struct {
	Offset   int
	Interval time.Duration
}

// NewRainbowAnimation creates a new RainbowAnimation
func NewRainbowAnimation(interval time.Duration) *RainbowAnimation {
	return &RainbowAnimation{
		Offset:   0,
		Interval: interval,
	}
}

// NextFrame advances the animation to the next frame
func (r *RainbowAnimation) NextFrame() {
	r.Offset = (r.Offset + 1) % 7 // 7 rainbow colors
}

// GetOffset returns the current color offset
func (r *RainbowAnimation) GetOffset() int {
	return r.Offset
}

// GetInterval returns the animation interval
func (r *RainbowAnimation) GetInterval() time.Duration {
	return r.Interval
}
