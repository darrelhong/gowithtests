package main

func Sum(numbers []int) int {
	return Reduce(numbers, func(accumulator, currentValue int) int {
		return accumulator + currentValue
	}, 0)
}

func SumAll(slices ...[]int) []int {
	return Reduce(slices, func(acc, curr []int) []int {
		return append(acc, Sum(curr))
	}, []int{})
}

func Reduce[T, U any](collection []T, acc func(U, T) U, initialValue U) U {
	result := initialValue
	for _, item := range collection {
		result = acc(result, item)
	}
	return result
}
