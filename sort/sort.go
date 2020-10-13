package sort

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var arrayLength = 8
var maxNum = 1000

func Index() {
	sort(insertSort, "insert sort")

	sort(bubbleSort1, "bubble sort v1")

	sort(bubbleSort2, "bubble sort v2")

	sort(mergeSort, "merge sort")

	sort(quickSort, "quick sort")

	sort(heapSort, "heap sort")

	sort(bucketSort, "bucket sort")

	sort(radixSortFromLow2High, "radix sort from low to high")

	sort(radixSortFromHigh2Low, "radix sort from high to low")

}

// radix sort from high to low
func radixSortFromHigh2Low(arr []int) []int {
	// element range from 0 to 999, please let maxNum = 1000
	bitSize := int(math.Log10(float64(maxNum)))
	res := helpRadix(arr, bitSize)
	return res
}

func helpRadix(arr []int, k int) []int {
	if k <= 0 || len(arr) == 0{
		return arr
	}

	record := make([][]int, 10) // 0 ~ 9
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		cur := v / int(math.Pow10(k-1)) % 10
		record[cur] = append(record[cur], v)
	}

	res := make([]int, 0, len(arr))
	for _, v := range record {
		ans := helpRadix(v, k-1)
		res = append(res, ans...)
	}

	return res
}

// radix sort from low to high
func radixSortFromLow2High(arr []int) []int {
	// element range from 0 to 999, please let maxNum = 1000
	bitSize := int(math.Log10(float64(maxNum))) - 1
	for k := bitSize; k >= 0; k-- {
		record := make([]int, 10) // 0 ~ 9
		for _, v := range arr {
			cur := v / int(math.Pow10(2-k)) % 10
			record[cur]++
		}
		total := -1
		for i, v := range record {
			total += v
			record[i] = total
		}
		// 3. 放置
		res := make([]int, len(arr))
		for i := len(arr) - 1; i >= 0; i-- {
			bucket := arr[i] / int(math.Pow10(2-k)) % 10
			res[record[bucket]] = arr[i]
			record[bucket]--
		}
		arr = res
	}
	return arr
}

// bucket sort
func bucketSort(arr []int) []int {
	record := make([]int, maxNum+1)

	// 1. 统计个数
	for _, v := range arr {
		record[v]++
	}

	// 2. 累加和
	total := -1
	for i, v := range record {
		total += v
		record[i] = total
	}

	// 3. 放置
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res[record[arr[i]]] = arr[i]
		record[arr[i]]--
	}

	return res
}

// heap sort
func heapSort(arr []int) []int {
	h := heap{}

	h.arr = arr
	// 1. build heap
	for i := range arr {
		h.add(i)
	}

	// 2. sort
	for i := 0; i < len(arr); i++ {
		h.remove()
	}

	arr = h.sorted
	return arr
}

func (h *heap) add(i int) {
	swap := func(arr []int, i, j int) {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}

	cur := i
	for h.getParent(cur) != -1 && h.arr[cur] > h.arr[h.getParent(cur)] {
		swap(h.arr, cur, h.getParent(cur))
		cur = h.getParent(cur)
	}
}

func (h *heap) remove() {
	h.sorted = append(h.sorted, h.arr[0])

	swap := func(arr []int, i, j int) {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	cur := 0
	for {
		l := h.getLeftChild(cur)
		r := h.getRightChild(cur)
		if l != -1 && r != -1 {
			if h.arr[cur] < h.arr[l] || h.arr[cur] < h.arr[r] {
				if h.arr[l] > h.arr[r] {
					swap(h.arr, l, cur)
					cur = l
				} else {
					swap(h.arr, r, cur)
					cur = r
				}
			} else {
				return
			}
		} else if l != -1 {
			if h.arr[cur] < h.arr[l] {
				swap(h.arr, l, cur)
				cur = l
			} else {
				return
			}
		} else {
			return
		}
	}

}

func (h *heap) getParent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func (h *heap) getLeftChild(i int) int {
	l := 2*i + 1
	if l < len(h.arr) {
		return l
	}
	return -1
}

func (h *heap) getRightChild(i int) int {
	l := 2*i + 2
	if l < len(h.arr) {
		return l
	}
	return -1
}

type heap struct {
	arr    []int
	sorted []int
}

// quickSort
func quickSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	j := partition(arr)
	quickSort(arr[:j])
	if j < len(arr)-1 {
		quickSort(arr[j+1:])
	}
	return arr
}

// quickSort partition
func partition(arr []int) int {
	if len(arr) <= 1 {
		return len(arr) - 1
	}

	key := arr[0]
	lo := 1
	hi := len(arr) - 1
	for {
		for lo < hi && arr[lo] < key {
			lo++
		}
		for arr[hi] > key {
			hi--
		}
		if lo >= hi {
			break
		}
		tmp := arr[lo]
		arr[lo] = arr[hi]
		arr[hi] = tmp
	}
	arr[0] = arr[hi]
	arr[hi] = key
	return hi
}

// insertSort
func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curElement := arr[i]
		k := i - 1
		for ; k >= 0 && arr[k] > curElement; k-- {
			arr[k+1] = arr[k]
		}
		// 注意是 k+1
		arr[k+1] = curElement
	}
	return arr
}

// bubble sort version 1
func bubbleSort1(arr []int) []int {
	swap := func(arr []int, i, j int) {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}
	for k := 0; k < len(arr); k++ {
		for i := 1; i < len(arr)-k; i++ {
			if arr[i-1] > arr[i] {
				swap(arr, i-1, i)
			}
		}
	}
	return arr
}

// bubble sort version 2
func bubbleSort2(arr []int) []int {
	swap := func(arr []int, i, j int) {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}
	swapped := true
	for k := 0; k < len(arr) && swapped; k++ {
		swapped = false
		for i := 1; i < len(arr)-k; i++ {
			if arr[i-1] > arr[i] {
				swap(arr, i-1, i)
				swapped = true
			}
		}
	}
	return arr
}

// mergeSort
func mergeSort(arr []int) []int {
	// 1. 递归终止
	if len(arr) <= 1 {
		return arr
	}

	// 2. 递归
	prePart := mergeSort(arr[:len(arr)/2])
	postPart := mergeSort(arr[len(arr)/2:])

	// 3. 合并
	prePoint := 0
	postPoint := 0
	res := make([]int, 0, len(arr))
	for prePoint < len(prePart) && postPoint < len(postPart) {
		if prePart[prePoint] < postPart[postPoint] {
			res = append(res, prePart[prePoint])
			prePoint++
		} else {
			res = append(res, postPart[postPoint])
			postPoint++
		}
	}
	for prePoint < len(prePart) {
		res = append(res, prePart[prePoint])
		prePoint++
	}
	for postPoint < len(postPart) {
		res = append(res, postPart[postPoint])
		postPoint++
	}
	return res
}

// util
func sort(sortFunc func([]int) []int, name string) {
	arr := generateArray()
	fmt.Printf("----- before sort by: %s-----\n", name)
	printResult(arr)
	res := sortFunc(arr)
	fmt.Printf("----- after sort by: %s-----\n", name)
	printResult(res)
	fmt.Println()
}

// util
func generateArray() []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, arrayLength)
	for i, _ := range arr {
		arr[i] = rand.Intn(maxNum)
	}
	return arr
}

// util
func printResult(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
