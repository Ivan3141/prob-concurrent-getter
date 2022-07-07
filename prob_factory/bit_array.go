package prob_factory

type BitArray struct {
	Bytes     []byte // bits packed into bytes.
	BitLength int    // length in bits.
	BitOffset int    //current visit position
	BitReset  func() (int, error)
}

func (b *BitArray) At(i int) int {
	if i < 0 || i >= b.BitLength {
		return 0
	}
	x := i / 8
	y := 7 - uint(i%8)
	return int(b.Bytes[x]>>y) & 1
}

func (b *BitArray) Next() (int, error) {
	if b.BitOffset+1 == b.BitLength {
		return 0, SyntaxError{"BitOffset over flow"}
	}
	b.BitOffset++
	return b.At(b.BitOffset), nil
}

func (b *BitArray) Set(pos int, val int) error {
	if pos < 0 || pos >= b.BitLength {
		return SyntaxError{"BitOffset over flow"}
	}
	if val != 0 && val != 1 {
		return SyntaxError{"Bit Value Set Error"}
	}
	x := pos / 8
	y := 7 - uint(pos%8)
	if val == 1 {
		b.Bytes[x] = b.Bytes[x] | (1 << y)
	} else {
		b.Bytes[x] = b.Bytes[x] & (^(1 << y))
	}
	return nil
}

func (b *BitArray) Append(bytes []byte) error {
	b.Bytes = append(b.Bytes, bytes...)
	b.BitLength += len(bytes) * 8
	return nil
}

func ParseBitArray(bytes []byte, size int) (ret *BitArray, err error) {

	if len(bytes) == 0 {
		err = SyntaxError{"zero length BIT ARRAY"}
		return
	}

	if size < 1 || size > 8*len(bytes) {
		err = SyntaxError{"size over flow BIT ARRAY"}
		return
	}

	ret = &BitArray{Bytes: bytes, BitLength: size}
	return
}

func (b *BitArray) OverFlow() bool {
	return b.BitOffset == b.BitLength
}

func NewBitArray(fn func() (int, error)) *BitArray {
	return &BitArray{Bytes: make([]byte, 0), BitLength: 0, BitOffset: -1, BitReset: fn}
}

func (b *BitArray) Recycle() error {
	byteSize := len(b.Bytes)
	if byteSize == 0 {
		return SyntaxError{"probArray empty"}
	}
	if !b.OverFlow() {
		return nil
	}
	b.BitOffset = -1
	for i := 0; i < b.BitLength; i++ {
		res, err := b.BitReset()
		if err != nil {
			return err
		}
		err = b.Set(i, res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BitArray) Check() []int {
	var res []int
	for i := 0; i < b.BitLength; i++ {
		res = append(res, b.At(i))
	}
	return res
}
