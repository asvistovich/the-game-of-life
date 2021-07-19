package main

type Life struct {
	width  int
	height int

	nextGrid    *Grid
	currentGrid *Grid
}

func NewLife(w, h int) *Life {
	return &Life{
		currentGrid: NewGrid(w, h).WithValues(),
		nextGrid:    NewGrid(w, h),
		width:       w, height: h,
	}
}

func (l *Life) NextGeneration() (completed bool) {
	// Update the state of the next grid from the current grid.
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			l.nextGrid.Set(x, y, l.currentGrid.GetState(x, y))
		}
	}

	// Check is all moves completed.
	if l.currentGrid.Equal(l.nextGrid) {
		return true
	}

	// Swap current and next grid.
	l.currentGrid, l.nextGrid = l.nextGrid, l.currentGrid
	return false
}

func (l *Life) String() string {
	return l.currentGrid.String()
}
