package main

import "fmt"

func solve (bricks []int, rest int, solution []int) {
	if rest == 0 {
		fmt.Println(solution)
		return
	}
	if len(bricks) == 0 {
		return
	}
	if rest >= bricks[0] {
		solve(bricks[1:], rest - bricks[0], append(solution, bricks[0]))
	}	
	solve (bricks[1:], rest, solution)
}	

func main(){
	bricks:=[]int{7, 11, 24, 3, 28, 4, 6, 12,1,13,31,654,432,8675,21}
	carrying:=48k
	solve(bricks, carrying, make([] int, 0))
}
