# Maze Solver
This is a command line maze solver which uses non-recursive depth-first search 

## User Guide

- First run the main.go which would ask number of maze(s) and the sizes of maze(s) you want to generate respectively.
  - Resulting mazes will be written to text files named as maze_N.txt where N represents the number of mazes.
  - File will contain x and y values in order to represent respective wall information as well.
  
```
go run main.go
```

- After successfully generating maze(s), you are asked to call mazeDrawer.mac file with the following command 
  which will draw the maze into a file named as mazeDrawn.txt .
  
```
./mazeDrawer.mac
```

- Give the number of maze that you want the program (main.go) to solve. 
  (NOTE:Program does not exit after generating the matrix to a .txt file)

- Give the entry point to the program (main.go) as x,y values.

- Give the exit point for the program (main.go) as x,y values.

Andddd voilà here are the respective route that shall save the stranger from the maze just appearing magically on the command line.

## Details

### Maze Generator
```go

// --------------------------------------------------------------------
// 4. The recursive-backtracking algorithm itself for generating
// maze by carving the walls in each iterations until every cell is
// visited.
// --------------------------------------------------------------------

func GenerateMaze(matrixHolder[][][]MazeNode){
	//Initially empty stack
	for maze:=0;maze<len(matrixHolder);maze++ {
		var stack = Stack{}
		var DIR = []string{"U", "D", "L", "R"}
		stack.Push(0, 0)
		matrixHolder[maze][0][0].visited = true
		var visitedCount = 1
		for visitedCount < COL*ROW {
			//current cell
			var x, y = stack.Top()
			var move = DIR[0]
			Shuffle(DIR)
			//choose next cells
			var nx, ny = x + MovementMapX(move), y + MovementMapY(move)
			//check cell
			if (InBetween(0, ROW-1, ny) && InBetween(0, COL-1, nx) && !matrixHolder[maze][ny][nx].visited) {
				//if checked conditions approved than break the wall and add the node
				matrixHolder[maze][ny][nx].visited = true
				visitedCount++
				stack.Push(nx, ny)
				if (move == "L") {
					matrixHolder[maze][y][x].L = false
					matrixHolder[maze][ny][nx].R = false
				} else if (move == "R") {
					matrixHolder[maze][y][x].R = false
					matrixHolder[maze][ny][nx].L = false
				} else if (move == "U") {
					matrixHolder[maze][y][x].U = false
					matrixHolder[maze][ny][nx].D = false
				} else {
					matrixHolder[maze][y][x].D = false
					matrixHolder[maze][ny][nx].U = false
				}
			} else if (UnvisitedNeighbors(x, y, matrixHolder[maze]) == 0) {
				stack.Pop()
			}

		}
	}
}
```
### Maze Solver
```go
// --------------------------------------------------------------------
// 7. The recursive-backtracking algorithm itself for solving the maze
// by traversing nodes and backtracking when gets stuck which runs until
// exit node is found.
// --------------------------------------------------------------------

func SolveMaze(matrix[][]MazeNode,entryX,entryY,exitX,exitY int){
	//initially empty stack
	var stack = Stack{}
	stack.Push(entryX,entryY)
	UnvisitMaze(matrix)
	matrix[entryY][entryX].visited=true
	var DIR = []string{"U", "D", "L", "R"}
	for !stack.isEmpty(){
		var currentX,currentY = stack.Top()
		Shuffle(DIR)
		var move = DIR[0]
		DIR  = DIR[1:]
		var nx,ny = currentX+MovementMapX(move),currentY+MovementMapY(move)
		if(InBetween(0,COL-1,nx)&&InBetween(0,ROW-1,ny)&&!matrix[ny][nx].visited&&RouteAllowed(matrix[currentY][currentX],move)){
			matrix[ny][nx].visited=true
			stack.Push(nx, ny)
			if(nx==exitX&&ny==exitY){
				break
			}
			DIR = []string{"U", "D", "L", "R"}
		}else if (UnvisitedNeighbors(currentX, currentY, matrix) == 0 || len(DIR)==0) {
			stack.Pop()
			DIR = []string{"U", "D", "L", "R"}
		}
	}
	stack.Display()
}
```
