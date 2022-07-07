package prob_factory

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_ProbBitArraryGen(t *testing.T) {
	probGenerator := new(ProbPlanA)
	probGenerator.rands[29] = rand.New(rand.NewSource(time.Now().Unix() + int64(0)))
	bitArray, _ := probGenerator.ProbBitArraryGen(30)
	res := bitArray.Check()
	for i := 0; i < len(res); i++ {
		if res[i] != TestResutl[i] {
			fmt.Println("check error out")
		}
	}
	fmt.Println("check success")
	fmt.Printf("len %d, offset %d \n, content %v \n", bitArray.BitLength, bitArray.BitOffset, bitArray.Check())
}

var testProb int

func Test_ProbBitPool(t *testing.T) {
	testProb = 1
	probGenerator := new(ProbPlanA)
	probGenerator.rands[testProb-1] = rand.New(rand.NewSource(time.Now().Unix() + int64(0)))
	probGenerator.randsPool[testProb-1] = NewRecyclerPool(func() Recycler {
		bitArray, _ := probGenerator.ProbBitArraryGen(testProb)
		return bitArray
	})
	res, err := probGenerator.ProbGet(testProb)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("prob get is %d \n", res)
}

func Test_ProbFactory(t *testing.T) {
	getter := NewProbGetter()
	wg := sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		res, err := getter.ProbGet(0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_0 is %d \n", res)
		}
	}()
	go func() {
		defer wg.Done()
		wg.Add(1)
		res, err := getter.ProbGet(1)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_1 is %d \n", res)
		}
	}()
	go func() {
		defer wg.Done()
		wg.Add(1)
		res, err := getter.ProbGet(30)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_30 is %d \n", res)
		}
	}()
	go func() {
		defer wg.Done()
		wg.Add(1)
		res, err := getter.ProbGet(70)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_70 is %d \n", res)
		}
	}()
	go func() {
		defer wg.Done()
		wg.Add(1)
		res, err := getter.ProbGet(99)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_99 is %d \n", res)
		}
	}()
	go func() {
		defer wg.Done()
		wg.Add(1)
		res, err := getter.ProbGet(100)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_100 is %d \n", res)
		}
	}()
	wg.Wait()
}

func Test_ProbParallel(t *testing.T) {
	getter := NewProbGetter()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		res, err := getter.ProbGet(30)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_30 is %d \n", res)
		}
		wg.Done()
	}()
	go func() {
		res, err := getter.ProbGet(30)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_30 is %d \n", res)
		}
		wg.Done()
	}()
	go func() {
		res, err := getter.ProbGet(30)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_30 is %d \n", res)
		}
		wg.Done()
	}()
	go func() {
		res, err := getter.ProbGet(30)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("res_30 is %d \n", res)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Test_ProbSingleThread(t *testing.T) {
	getter := NewProbGetter()
	res, err := getter.ProbGet(30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("res_30 is %d \n", res)
	}

	res, err = getter.ProbGet(30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("res_30 is %d \n", res)
	}

	res, err = getter.ProbGet(30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("res_30 is %d \n", res)
	}

	res, err = getter.ProbGet(30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("res_30 is %d \n", res)
	}
}
