# Conway's Game of Life


## Rules
At each step in time, the following transitions occur:
1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
   The initial pattern constitutes the seed of the system. The first generation is created by applying the
   above rules simultaneously to every cell in the seed—births and deaths occur simultaneously, and the
   discrete moment at which this happens is sometimes called a tick (in other words, each generation is a
   pure function of the preceding one). The rules continue to be applied repeatedly to create further
   generations.

## Install

Just get the source and build/install as normal

```
go get -u github.com/asvistovich/the-game-of-life
go install github.com/asvistovich/the-game-of-life
```

## Examples

[a link](https://github.com/asvistovich/the-game-of-life/blob/main/examples/20210719092327.gif)

![20210719092327.gif](https://media.giphy.com/media/57UG2EAm8HJeNVQOWl/giphy.gif)
