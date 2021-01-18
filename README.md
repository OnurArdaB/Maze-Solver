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

- Give the number of maze that you want the program (main.go) to solve. (NOTE:Program does not exit after generating the matrix to a .txt file)

- Give the entry point to the program (main.go) as x,y values.

- Give the exit point for the program (main.go) as x,y values.

Andddd voila here are the respective route that shall save the stranger from the maze just appearing magically on the command line.

