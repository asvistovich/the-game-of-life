package main

import (
	"bytes"
	"math/rand"
	"reflect"
)

type Grid struct {
	grid [][]bool
	w, h int
}

// NewGrid returns an empty grid.
func NewGrid(w, h int) *Grid {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Grid{grid: s, w: w, h: h}
}

func (g *Grid) Set(x, y int, b bool) {
	g.grid[y][x] = b
}

// GetState returns active states of selected cell.
func (g *Grid) GetState(x, y int) bool {
	alive := g.CountNeighbors(x, y)
	if alive == 3 {
		return true
	}
	if alive == 2 && g.IsAlive(x, y) {
		return true
	}
	return false
}

// WithValues randomizes initial setup.
func (g *Grid) WithValues() *Grid {
	for i := 0; i < (g.w * g.h / 4); i++ {
		g.Set(rand.Intn(g.w), rand.Intn(g.h), true)
	}
	return g
}

// IsAlive returns cell status.
func (g *Grid) IsAlive(x, y int) bool {
	x += g.w
	x %= g.w
	y += g.h
	y %= g.h
	return g.grid[y][x]
}

// CountNeighbors get state of eight connected cells.
func (g *Grid) CountNeighbors(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && g.IsAlive(x+i, y+j) {
				alive++
			}
		}
	}
	return alive
}

// String returns grid as a string.
func (g *Grid) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			b := byte(' ')
			if g.grid[x][y] {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (g *Grid) Equal(v *Grid) bool {
	if g == nil || v == nil {
		return false
	}
	return reflect.DeepEqual(*g, *v)
}
