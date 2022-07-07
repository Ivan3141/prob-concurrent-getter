package prob_factory

import "math/rand"

type ProbPlanA struct {
	//概率1-99发生器，0必不发生，100一定发生
	rands     [99]*rand.Rand
	randsPool [99]*RecyclerPool
}

//概率生成1250个字节的结果,10000个存储量
const ProbBatchSize = 1250

//以prob概率计算结果，如果hit则返回1，否则返回0
func (planA *ProbPlanA) ProbHitCompute(prob int) (int, error) {
	if prob < 0 || prob > 100 {
		return ProbInvalid, SyntaxError{"input out of range"}
	}
	if prob == 0 {
		return ProbMiss, nil
	}
	if prob == 100 {
		return ProbHit, nil
	}
	val := planA.rands[prob-1].Intn(100)
	if val < prob {
		return ProbHit, nil
	}

	return ProbMiss, nil
}

func (planA *ProbPlanA) ProbBitArraryGen(prob int) (*BitArray, error) {
	var probArray []uint8
	for i := 0; i < ProbBatchSize; i++ {
		probByte := uint8(0)
		for j := 7; j >= 0; j-- {
			res, err := planA.ProbHitCompute(prob)
			if err != nil {
				return nil, err
			}
			if res == ProbHit {
				probByte = probByte | (1 << j)
			}
		}
		probArray = append(probArray, probByte)
	}
	probBitArray := NewBitArray(func() (int, error) { return planA.ProbHitCompute(prob) })
	err := probBitArray.Append(probArray)
	if err != nil {
		return nil, err
	}
	return probBitArray, nil
}

func (planA *ProbPlanA) ProbBitArrayReuse(prob int, bitArray *BitArray) error {
	byteSize := len(bitArray.Bytes)
	if byteSize == 0 {
		return SyntaxError{"probArray empty"}
	}
	bitArray.BitOffset = -1
	for i := 0; i < bitArray.BitLength; i++ {
		res, err := planA.ProbHitCompute(prob)
		if err != nil {
			return err
		}
		err = bitArray.Set(i, res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (planA *ProbPlanA) ProbGet(prob int) (res int, err error) {
	if prob < 0 || prob > 100 {
		return ProbInvalid, SyntaxError{"prob out of range"}
	}
	if prob == 0 {
		return ProbMiss, nil
	}
	if prob == 100 {
		return ProbHit, nil
	}
	bitArray := planA.randsPool[prob-1].Get().(*BitArray)
	defer func() {
		if err = planA.randsPool[prob-1].Put(bitArray); err != nil {
			res = ProbInvalid
		}
	}()
	if bitArray.OverFlow() {
		_ = bitArray.Recycle()
	}
	res, err = bitArray.Next()
	if err != nil {
		return ProbInvalid, err
	}
	return res, nil
}
