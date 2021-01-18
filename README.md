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
