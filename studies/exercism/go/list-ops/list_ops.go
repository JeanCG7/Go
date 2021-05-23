package listops

type binFunc func(a, b int) int
type predFunc func(a int) bool
type unaryFunc func(a int) int

type IntList []int

func (il IntList) Length() int {
	length := 0
	for range il {
		length++
	}
	return length
}

func (il IntList) Append(src IntList) IntList {
	result := make(IntList, il.Length()+src.Length())

	for i := 0; i < il.Length(); i++ {
		result[i] = il[i]
	}

	i := il.Length()
	for j := 0; j < src.Length(); j++ {
		result[i] = src[j]
		i++
	}

	return result
}

func (il IntList) Concat(src []IntList) IntList {
	totalLength := 0
	for _, v := range src {
		totalLength = totalLength + v.Length()
	}
	result := make([]int, totalLength)
	i := 0
	for _, il := range src {
		for _, v := range il {
			result[i] = v
			i++
		}
	}
	return il.Append(result)
}

func (il IntList) Reverse() IntList {
	result := make(IntList, il.Length())
	j := result.Length() - 1

	for i := range il {
		result[j] = il[i]
		j--
	}

	return result
}

func (il IntList) Foldl(fn binFunc, b int) int {
	for i := range il {
		if b != 0 {
			b = fn(il[i], b)
		}
	}

	return b
}

func (il IntList) Foldr(fn binFunc, b int) int {
	reverse := il.Reverse()
	return reverse.Foldl(fn, b)
}

func (il IntList) Filter(fn predFunc) IntList {
	k := 0
	for i, n := range il {
		if fn(n) {
			if i != k {
				il[k] = n
			}
			k++
		}
	}
	il = il[:k]

	return il
}

func (il IntList) Map(fn unaryFunc) IntList {
	for i := range il {
		il[i] = fn(il[i])
	}
	return il
}
