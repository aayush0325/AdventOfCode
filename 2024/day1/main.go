package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/aayush0325/AdventOfCode/utils"
)

func main(){

	result, err := utils.ReadFileAsInt("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dif := 0;
	a := make([]int,0);
	b := make([]int,0);

	for i, value := range result {
		if i % 2 == 0 {
			a = append(a, value);
		} else {
			b = append(b, value);
		}
	}	

	sort.Ints(a)
	sort.Ints(b)
	
	n := len(a)

	for i := 0 ; i < n ; i++ {
		dif += int(math.Abs(float64(a[i] - b[i])))
	}

	fmt.Println("The answer for 1st part:",dif) // first part

	sum := 0

	mapB := make(map[int]int)

	for i := 0 ; i < n ; i++ {
		mapB[b[i]]++
	}

	for i := 0 ; i < n ; i++ {
		sum += a[i] * mapB[a[i]]
	}

	fmt.Println("The answer for 2nd part:",sum)
}


