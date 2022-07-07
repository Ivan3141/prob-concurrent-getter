package prob_factory

import (
	"math/rand"
)

//如果没有数据，则填100000个计算结果进去
const probPlanBBatchSize = 100000

type ProbGenB struct {
	randGen   *rand.Rand
	randQueue *poolChain
}

type ProbPlanB struct {
	//概率1-99发生器，0必不发生，100一定发生
	probProduceSignal chan int
	probGenB          [99]*ProbGenB
}

//以prob概率计算结果，如果hit则返回1，否则返回0
func (planB *ProbPlanB) probHitCompute(prob int) (int, error) {
	if prob < 0 || prob > 100 {
		return ProbInvalid, SyntaxError{"input out of range"}
	}
	if prob == 0 {
		return ProbMiss, nil
	}
	if prob == 100 {
		return ProbHit, nil
	}
	val := planB.probGenB[prob-1].randGen.Intn(100)
	if val < prob {
		return ProbHit, nil
	}

	return ProbMiss, nil
}

func (planB *ProbPlanB) ProbGet(prob int) (res int, err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			res, _ = planB.probHitCompute(prob)
			err = nil
		}
	}()
	if prob < 0 || prob > 100 {
		return ProbInvalid, SyntaxError{"prob out of range"}
	}
	if prob == 0 {
		return ProbMiss, nil
	}
	if prob == 100 {
		return ProbHit, nil
	}
	if res, ok := planB.probGenB[prob-1].randQueue.PopTail(); ok {
		return res.(int), nil
	}
	planB.probProduceSignal <- prob
	res, _ = planB.probHitCompute(prob)
	return res, nil
}

func (planB *ProbPlanB) probProduce() {
	lastGetProb := -1
	go func() {
		for {
			select {
			case prob := <-planB.probProduceSignal:
				if prob != lastGetProb {
					planB.probMakeUp(prob)
				}
				lastGetProb = prob
			}
		}
	}()
}

func (planB *ProbPlanB) probMakeUp(prob int) {
	if prob < 1 || prob > 99 {
		return
	}
	for i := 0; i < probPlanBBatchSize; i++ {
		res, _ := planB.probHitCompute(prob)
		planB.probGenB[prob-1].randQueue.PushHead(res)
	}
}
