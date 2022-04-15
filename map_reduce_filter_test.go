package MR

import (
	"testing"
)
//must import testing
//must be named something_test.go
//tests must be named func TestXxx(t *testing.T) where Xxx leads with capital
//benchmarks must be named func BenchmarkXxx(b *testing.B) and ^
//benchmark output gives tag, how many times loop ran, and speed of loop

func slicesEqual[T comparable](sliceOne, sliceTwo []T) bool {
	if len(sliceOne) != len(sliceTwo) {
		return false
	}

	for i := 0; i < len(sliceOne); i++ {
		if sliceOne[i] != sliceTwo[i] {
			return false
		}
	}
	return true
}

func TestSlicesEqual(t *testing.T) {
	numsOne := []int{1,2,3,4,5}
	numsTwo := []int{1,2,3}
	numsThree := []int{1,2,3,4,6}
	numsFour := []int{1,2,3,4,5}

	if !slicesEqual(numsOne, numsFour) {
		t.Errorf("Unique slices with same elements should be the same")
		return
	}

	if !slicesEqual(numsOne, numsOne) {
		t.Errorf("Same slice shold be the same")
		return
	}

	if slicesEqual(numsOne, numsTwo) {
		t.Errorf("Slices of different length should not be equal")
		return
	}

	if slicesEqual(numsOne, numsThree) {
		t.Errorf("Slices of same length with different elements should not be equal")
		return
	}

	return
}


func TestMapSquare(t *testing.T) {
	n := 100
	nums := make([]int, n)
	squaredNums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
		squaredNums[i] = i * i
	}
	square := func(x int) int { return x * x }

	mapTest := Map(nums, square)
	if !slicesEqual(mapTest, squaredNums){
		t.Errorf("squared mapping gave %v", mapTest)
	}
	return
}

func TestReduce(t *testing.T){
	answer := 0
	n := 100
        nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
		answer += i
	}

	add := func(x,y int) int { return x + y }
	reduceNums := Reduce(nums,0,add)
	if reduceNums != answer {
		t.Errorf("reduce using add resulted in %d, should be %d",
			reduceNums, answer)
	}
	return
}

func TestFilter(t *testing.T){
        n := 100
        nums := make([]int, n)
	answer := make([]int, 0, n)
	isOdd := func(x int) bool { return x%2 != 0 }
        for i := 0; i < n; i++ {
                nums[i] = i
                if isOdd(i) {
			answer = append(answer, i)
		}
        }

        filterNums := Filter(nums, isOdd)
	if !slicesEqual(answer, filterNums) {
		t.Errorf("slices aren't equal. Is: %v, should be %v",
			filterNums, answer)
	}
	return
}

func TestAltFilter(t *testing.T){
        n := 100
        nums := make([]int, n)
	answer := make([]int, 0, n)
	isOdd := func(x int) bool { return x%2 != 0 }
        for i := 0; i < n; i++ {
                nums[i] = i
                if isOdd(i) {
			answer = append(answer, i)
		}
        }

        filterNums := AltFilter(nums, isOdd)
	if !slicesEqual(answer, filterNums) {
		t.Errorf("slices aren't equal. Is: %v, should be %v",
			filterNums, answer)
	}
	return
}


func BenchmarkMapSquare(b *testing.B){
	n := 100
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
        square := func(x int) int { return x * x }

        for i := 0; i < b.N; i++ {
                Map(nums, square)
        }
        return
}

func BenchmarkMapPlusOne(b *testing.B){
	n := 100
	        nums := make([]int, n)
        for i := 0; i < n; i++ {
                nums[i] = i
        }
        addOne := func(x int) int { return x + 1 }

        for i := 0; i < b.N; i++ {
                Map(nums, addOne)
        }
        return
}

func BenchmarkMapIncrement(b *testing.B){
	n := 100
	nums := make([]int, n)
        for i := 0; i < n; i++ {
                nums[i] = i
        }
        increment := func(x int) int { x++; return x }

        for i := 0; i < b.N; i++ {
                Map(nums, increment)
        }
        return
}

func BenchmarkReduce(b *testing.B){
	n := 100
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	add := func(x, y int) int { return x + y }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(nums, 0, add)
	}
}

func BenchmarkFilter(b *testing.B){
	n := 100
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}

	isOdd := func(x int) bool { return x%2!=0 }

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Filter(nums, isOdd)
	}
	return
}

func BenchmarkAltFilter(b *testing.B){
	        n := 100
        nums := make([]int, n)
        for i := 0; i < n; i++ {
                nums[i] = i
        }

        isOdd := func(x int) bool { return x%2!=0 }

        b.ResetTimer()

        for i := 0; i < b.N; i++ {
                AltFilter(nums, isOdd)
        }
        return
}

