package prob_factory

import (
	"math/rand"
	"time"
)

//for prob test
var TestResutl []int

type ProbGetter interface {
	ProbGet(int) (int, error)
}

func NewProbGetter() ProbGetter {
	return probGeneratorB
}

var probGeneratorB *ProbPlanB

func init() {
	probGeneratorB = new(ProbPlanB)
	for i := 0; i < 99; i++ {
		probGeneratorB.probGenB[i] = new(ProbGenB)
		probGeneratorB.probGenB[i].randGen = rand.New(rand.NewSource(time.Now().Unix() + int64(i)))
		probGeneratorB.probGenB[i].randQueue = NewPoolChain()
	}
	probGeneratorB.probProduceSignal = make(chan int)
	probGeneratorB.probProduce()
}
