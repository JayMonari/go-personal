package main

type Ball struct{ x, y, speedX, speedY int }

const ballRune = '\u24cf'

func (b *Ball) Update() {
	b.x += b.speedX
	b.y += b.speedY
}

func (b *Ball) Bounce(maxWidth int, maxHeight int) {
	if b.x == 0 || b.x == maxWidth-1 {
		b.speedX *= -1
	}
	if b.y == 0 || b.y == maxHeight-1 {
		b.speedY *= -1
	}
}
