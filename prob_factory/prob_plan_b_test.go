package prob_factory

import (
	"fmt"
	"sync"
	"testing"
)

func Test_ProbPlanB(t *testing.T) {
	planB := NewProbGetter()
	for i := 0; i < 5; i++ {
		res, _ := planB.ProbGet(0)
		fmt.Printf("prob %d result is %d \n", 0, res)
	}
	for i := 0; i < 5; i++ {
		res, _ := planB.ProbGet(10)
		fmt.Printf("prob %d result is %d \n", 10, res)
	}
	for i := 0; i < 5; i++ {
		res, _ := planB.ProbGet(20)
		fmt.Printf("prob %d result is %d \n", 20, res)
	}
	for i := 0; i < 5; i++ {
		res, _ := planB.ProbGet(100)
		fmt.Printf("prob %d result is %d \n", 100, res)
	}
	for i := 0; i < 1; i++ {
		res, err := planB.ProbGet(101)
		fmt.Printf("prob %d result is %d, err is %s \n", 101, res, err.Error())
	}
	for i := 0; i < 1; i++ {
		res, err := planB.ProbGet(-2)
		fmt.Printf("prob %d result is %d, err is %s \n", -2, res, err.Error())
	}
}

func Test_ProbPlanB_OutSize(t *testing.T) {
	num := 15
	planB := NewProbGetter()
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			res, _ := planB.ProbGet(10)
			fmt.Printf("prob %d result is %d \n", 10, res)
		}()
	}
	wg.Wait()
}
