package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// --------------------------------------------------------------------
//	1.Implement a stack class for using in the recursive backtracking
//	maze generator and solver.Implemented stack will contain only the
//	basic functionalities of the stacks.
// --------------------------------------------------------------------

type Node struct {
	next * Node
	X int
	Y int
	DIR []string
}

type Stack struct{
	head * Node
	size int
}

func (S * Stack) Push(X,Y int){
	temp:=&Node{
		next:nil,
		X:X,
		Y:Y,
	}
	if S.head==nil {
		S.head=temp
	} else {
		temp.next=S.head
		S.head=temp
	}
	S.size++
}

func (S * Stack) Display(){
	temp:=S.head

	for temp!=nil{
		fmt.Println("X:",temp.X,"Y:",temp.Y)
		temp=temp.next
	}
}

func (S * Stack) Top()(X,Y int){
	if(S.size>0){
		return S.head.X,S.head.Y
	}
	return
}

func (S* Stack) Pop()( X,Y int){
	if(S.size<=0){
		return
	}
	var x,y = S.head.X,S.head.Y
	S.head = S.head.next
	S.size--
	return x,y
}

func (S*Stack) isEmpty() bool{
	return S.head==nil
}

// --------------------------------------------------------------------
//	2.Create the grid for the maze.Grids will contain custom nodes.
// --------------------------------------------------------------------

type MazeNode struct {
	L bool
	R bool
	U bool
	D bool
	visited bool
}

var ROW,COL int

// --------------------------------------------------------------------
//	3.MovementMaps for returning values regarding to moves with axis.
// --------------------------------------------------------------------

func MovementMapY(i string)int{
	if(i=="U"){
		return 1
	}else if(i=="D"){
		return -1
	}else{
		return 0
	}
}

func MovementMapX(i string)int{
	if(i=="R"){
		return 1
	}else if(i=="L"){
		return -1
	}else{
		return 0
	}
}

func Shuffle(array []string) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func InBetween(min,max,number int)bool{
	return (number >= min) && (number <= max)
}

func checkBitSetVar(bool bool) int{
	if bool{
		return 1
	}
	return 0
}

func writeMatrixToFile(matrixHolder [][][]MazeNode,filename string){

	for i := 0; i < len(matrixHolder); i++ {
		fi,err:=os.Create(filename + strconv.Itoa(i+1))
		if(err!=nil){
			panic(err)
		}

		//fmt.Println("Maze:",i+1)
		fmt.Fprintln(fi,strconv.Itoa(ROW)+" "+strconv.Itoa(COL))
		for R := 0; len(matrixHolder[i]) > R; R++ {
			for C := 0; len(matrixHolder[i][R]) > C; C++ {
				var temp = ("x="+strconv.Itoa(C) +" y="+strconv.Itoa(R) +" l="+strconv.Itoa(checkBitSetVar(matrixHolder[i][R][C].L))+" r="+strconv.Itoa(checkBitSetVar(matrixHolder[i][R][C].R))+ " u="+strconv.Itoa(checkBitSetVar(matrixHolder[i][R][C].U))+ " d="+strconv.Itoa(checkBitSetVar(matrixHolder[i][R][C].D)))
				fmt.Fprintln(fi,temp)
			}
		}
		fi.Close()
	}
	fmt.Println("All mazes are generated.")
}

func UnvisitedNeighbors(X,Y int,matrix[][]MazeNode)int{
	var count  = 0
	if(InBetween(0,COL-1,X+1)&&InBetween(0,ROW-1,Y)){ //x+1 y
		if(!matrix[Y][X+1].visited){
			count++
		}
	}
	if(InBetween(0,COL-1,X-1)&&InBetween(0,ROW-1,Y)){ //x-1 y
		if(!matrix[Y][X-1].visited){
			count++
		}
	}
	if(InBetween(0,COL-1,X)&&InBetween(0,ROW-1,Y+1)){ //x y+1
		if(!matrix[Y+1][X].visited){
			count++
		}
	}
	if(InBetween(0,COL-1,X)&&InBetween(0,ROW-1,Y-1)){ //x y-1
		if(!matrix[Y-1][X].visited){
			count++
		}
	}
	return count
}

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

// --------------------------------------------------------------------
// 5. Nodes in the grid(matrix) must be turned to unvisited.
// --------------------------------------------------------------------

func UnvisitMaze(matrix[][]MazeNode){
	for R := 0; len(matrix) > R; R++ {
		for C := 0; len(matrix[R]) > C; C++ {
			matrix[R][C].visited=false
		}
	}
}

// --------------------------------------------------------------------
// 6. Check if the specified move is possible for the current node.
// --------------------------------------------------------------------

func RouteAllowed(node MazeNode,move string)bool{
	if(move=="D"){
		return !node.D
	}else if(move=="U"){
		return !node.U
	}else if(move=="L"){
		return !node.L
	}else{
		return !node.R
	}
}

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

func main(){
	var NumMaze int
	for NumMaze!=-1 {
		fmt.Println("This is a program for generating and solving maze.\nPress -1 or CTRL+C in order to quit.")
		fmt.Print("Enter the number of mazes: ")
		fmt.Scan(&NumMaze)
		if(NumMaze==-1){
			break
		}

		fmt.Print("Enter the number of rows and columns (M and N): ")
		fmt.Scan(&ROW, &COL)		
		MatHolder := make([][][]MazeNode, NumMaze)

		for i := 0; i < NumMaze; i++ {
			MAT := make([][]MazeNode, ROW)
			for R := 0; ROW > R; R++ {
				MAT_LINE := make([]MazeNode, COL)
				for C := 0; COL > C; C++ {
					var temp = MazeNode{true, true, true, true, false}
					MAT_LINE[C] = temp
				}
				MAT[R] = MAT_LINE
			}
			MatHolder[i] = MAT
		}

		GenerateMaze(MatHolder)

		writeMatrixToFile(MatHolder, "maze_")

		var ToFind,EntryX,EntryY,ExitX,ExitY int

		fmt.Print("Enter a maze ID between 1 to ",NumMaze," inclusive to find a path: ")
		fmt.Scan(&ToFind)

		fmt.Print("Enter x and y coordinates of the entry points (x,y) or (column,row): ")
		fmt.Scan(&EntryX,&EntryY)

		fmt.Print("Enter x and y coordinates of the exit points (x,y) or (column,row): ")
		fmt.Scan(&ExitX,&ExitY)

		SolveMaze(MatHolder[ToFind-1],EntryX,EntryY,ExitX,ExitY)
	}
}
