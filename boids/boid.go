package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position Vector2d
	velocity Vector2d
}

func (b Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if 
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		id:       bid,
		position: Vector2d{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2d{(rand.Float64() * 2) - 1, (rand.Float64() * 2) - 1},
	}
	boids[bid] = &b
	go b.start()
}
