package main

import (
	"fmt"
	"container/list"
	"bufio"
)

const (
	DIRECTION_YY = 1
	DIRECTION_BG = 2
	NO_DIRECTION = 0
)
/*
problem statememt:
http://www.spoj.com/problems/MARTIAN/
 */

type MartianCell struct {
	CellDirection int
	SumBg int64
	SumYy int64
}

func main() {
	
	reader := bufio.NewReader(os.Stdin)

	str ,err  := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There is some problem with input", err)
		return
	}

	str = str[:len(str)-1]
	strs  := strings.Split(str , " ")
	if len(strs) != 2 {
		fmt.Println("There is some problem with m & n", err)
		return
	}

	m, err := strconv.Atoi(strs[0])
	if err != nil {
		fmt.Println("There is some problem with m", err)
		return
	}

	n, err  := strconv.Atoi(strs[1])
	if err != nil {
		fmt.Println("There is some problem with n", err)
		return
	}

	yy := make([][]int64, m)
	for i:=0; i<m; i++ {
		str , err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("There is some problem with yy input", err)
			return
		}

		str = str[:len(str)-1]
		strs  := strings.Split(str , " ")

		if len(strs) != n {
			fmt.Println("There is some problem with yy input  n on line ", i)
			return
		}
		yy[i] = make([]int64, n)
		for j:=0; j <n; j ++ {
			yy[i][j], err  = strconv.ParseInt(strs[j], 10, 64)
			if err != nil {
				fmt.Println("There is some problem with yy input ", i, err)
				return
			}
		}
	}
	m :=4
	n := 4
	yy := [][]int64{{0,0,0,10},{1,3,10,0},{4,2,1,3},{1,1,20,0}}
	fmt.Println("YY INput is :", m, n, yy)
	
	//bg := make([][]int64, m)
	//for i:=0; i<m; i++ {
	//	str , err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("There is some problem with yy input", err)
	//		return
	//	}
	//
	//	str = str[:len(str)-1]
	//	strs  := strings.Split(str , " ")
	//
	//	if len(strs) != n {
	//		fmt.Println("There is some problem with yy input  n on line ", i)
	//		return
	//	}
	//	bg[i] = make([]int64, n)
	//	for j:=0; j <n; j ++ {
	//		bg[i][j], err  = strconv.ParseInt(strs[j], 10, 64)
	//		if err != nil {
	//			fmt.Println("There is some problem with yy input ", i, err)
	//			return
	//		}
	//	}
	//}
	bg := [][]int64{{10,0,0,0},{1,1,1,30},{0,0,5,5},{5,10,10,10}}
	fmt.Println("BG INput is :", m, n, bg)
	
	var maxMinedQuantity int64
	
	minedData := make([][]MartianCell, m)
	for i :=m-1; i>=0; i-- {
		for j := n - 1; j >= 0; j-- {
			minedData[i] = make([]MartianCell, n)
		}
	}
	
	for i :=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			CalculateMaxMinerals(i, j, m, n, yy, bg, minedData, &maxMinedQuantity)
			fmt.Println("minedData", minedData)
			fmt.Println()
		}
	}
	
	fmt.Println("Max Quantity : ", maxMinedQuantity)
	
}


func Max(a, b int64) int64{
	if a >=b {
		return a
	}
	return b
}

func CalculateMaxMinerals(i, j, m, n int, yy,bg [][]int64, mc [][]MartianCell, maxMinedQuantity *int64) {
	// calculate north sum
	var ySum, bSum int64
	if mc[i][j].CellDirection == NO_DIRECTION  || mc[i][j].CellDirection == DIRECTION_YY {
		if i == m - 1 {
			for k := i; k >= 0; k-- {
				bSum = bSum + bg[k][j]
			}
		} else {
			bSum = mc[i + 1][j].SumBg - bg[i + 1][j]
		}
		fmt.Println("bsum", bSum)
		mc[i][j].SumBg = bSum
	}
	
	if mc[i][j].CellDirection == NO_DIRECTION  || mc[i][j].CellDirection == DIRECTION_BG{
		if j == n - 1 {
			for k := j; k >= 0; k-- {
				ySum = ySum + yy[i][k]
			}
		} else {
			ySum = mc[i][j + 1].SumYy - yy[i][j + 1]
		}
		fmt.Println("ysum", ySum)
		mc[i][j].SumYy = ySum
	}
	
	switch mc[i][j].CellDirection {
	case DIRECTION_YY:
		if mc[i][j].SumBg > mc[i][j].SumYy {
			mc[i][j].CellDirection = DIRECTION_BG
			// rework this formula
			(*maxMinedQuantity) = (*maxMinedQuantity) + mc[i][j].SumBg - mc[i][j].SumYy
			fmt.Println("maxq", *maxMinedQuantity)
			for k:=i-1; k>=0; k-- {
				if mc[k][j].CellDirection == DIRECTION_YY {
					for l:=j-1; l>=0; l-- {
						mc[k][l].CellDirection = NO_DIRECTION
					}
				}
				mc[k][j].CellDirection = DIRECTION_BG
				mc[k][j].SumBg = mc[i][j].SumBg
			}
		}
	case DIRECTION_BG:
		if mc[i][j].SumYy > mc[i][j].SumBg {
			mc[i][j].CellDirection = DIRECTION_YY
			// rework this formula
			(*maxMinedQuantity) = (*maxMinedQuantity) + mc[i][j].SumYy - mc[i][j].SumBg
			fmt.Println("maxq", *maxMinedQuantity)
			for k:=j-1; k>=0; k-- {
				if mc[i][k].CellDirection == DIRECTION_BG {
					for l:=i-1; l>=0; l-- {
						mc[l][k].CellDirection = NO_DIRECTION
					}
				}
				mc[i][k].CellDirection = DIRECTION_YY
				mc[i][k].SumYy = mc[i][j].SumYy
			}
		}
	case NO_DIRECTION:
		if mc[i][j].SumBg >= mc[i][j].SumYy {
			mc[i][j].CellDirection = DIRECTION_BG
			// rework this formula
			(*maxMinedQuantity) = (*maxMinedQuantity) + mc[i][j].SumBg
			for k:=i-1; k>=0; k-- {
				mc[k][j].CellDirection = DIRECTION_BG
				mc[k][j].SumBg = mc[i][j].SumBg
			}
		} else if mc[i][j].SumYy > mc[i][j].SumBg {
			mc[i][j].CellDirection = DIRECTION_YY
			// rework this formula
			(*maxMinedQuantity) = (*maxMinedQuantity) + mc[i][j].SumYy
			for k:=j-1; k>=0; k-- {
				mc[i][k].CellDirection = DIRECTION_YY
				mc[i][k].SumYy = mc[i][j].SumYy
			}
		}
		fmt.Println("maxq", *maxMinedQuantity)
	default:
		panic("invalid direction")
	}
	return
}                                        
