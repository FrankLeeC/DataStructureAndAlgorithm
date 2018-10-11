package test

import (
	"testing"
)

var array = []int{81, 94, 11, 96, 12, 35, 17, 95, 28, 58, 41, 75, 15}

// var array = []int{22, 2, 222, 2, 2, 2222}

func TestBubbleSort(t *testing.T) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	t.Logf("%v\n", array)
}

func TestShellSort(t *testing.T) {
	n := len(array)
	var i int
	for gap := n / 2; gap > 0; gap /= 2 {
		for j := gap; j < n; j++ {
			tmp := array[j]
			for i = j; i >= gap; i -= gap {
				if tmp < array[i-gap] {
					array[i] = array[i-gap]
				} else {
					break
				}
			}
			array[i] = tmp
		}
	}
	t.Logf("%v\n", array)
}

/*
插入排序：N-1次循环，每次循环，当前位置及以前的数字都排好序。
最好时间复杂度：O(N)
最差时间复杂度：O(N^2)
平均时间复杂度：O(N^2)
*/
func TestInsertSort(t *testing.T) {
	var j, p, tmp int
	for p = 1; p < len(array); p++ {
		tmp = array[p]
		for j = p; j > 0 && array[j-1] > tmp; j-- {
			array[j] = array[j-1]
		}
		array[j] = tmp
	}
	t.Logf("%v\n", array)
}

// 快速排序，每轮，找一个数做为基准（第一个），从右往左找小的交换，然后从左往右找大的交换，往复直至首尾指针相遇或交错。然后递归基准数的两侧子序列。
func TestQuickSort(t *testing.T) {
	quickSort(array)
	t.Logf("%v\n", array)
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	n := len(arr)
	i, j := 0, n-1
	p := 0
	for i <= j {
		// 先从右往左，遇到arr[j] < arr[p]就调换两者
		for j >= 0 {
			if arr[j] < arr[p] {
				arr[p], arr[j] = arr[j], arr[p]
				p = j
				j--
				break // 找到就退出循环，开始从左往右找大的
			}
			j--
		}
		// 判断条件是i <= j，不是i < j。当从左往右开始第一次比较并且i == j时也有可能arr[i] > arr[p]
		for i <= j {
			if arr[i] > arr[p] {
				arr[i], arr[p] = arr[p], arr[i]
				p = i
				i++
				break // 找到就退出循环，再次从右往左找小的
			}
			i++
		}
	}
	// i >= j并且无满足条件的数需要交换，退出本次寻找。此时k左边小于k，右边大于k
	// 递归k两侧
	quickSort(arr[0:p])
	quickSort(arr[p+1:])
}

// 每次从无序区找出最小的放到无序区第0个位置，有序区扩展一个数，无序区缩减一个数。重复直到无序区只剩一个数，添加至有序区后全部有序。
func TestSelectSort(t *testing.T) {
	SelectSort(array)
	t.Logf("%v\n", array)
}

func SelectSort(arr []int) {
	if len(arr) == 1 {
		return
	}
	i := argmin(arr)
	if i > 0 {
		arr[0], arr[i] = arr[i], arr[0]
	}
	SelectSort(arr[1:])
}

func argmin(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	i := 0
	m := arr[0]
	for j := 1; j < len(arr); j++ {
		if arr[j] < m {
			m = arr[j]
			i = j
		}
	}
	return i
}

// 堆排序，无序区构造大根堆，将堆顶与最后一个元素交换，剩下的做为无序区，重复至无序区仅剩一个元素，添加至有序区后完全有序。
func TestHeapSort(t *testing.T) {
	n := len(array)
	p := n - 1
	for ; p >= 1; p-- {
		k := n/2 - 1
		for i := k; i >= 0; i-- {
			adjust(i, 2*i+1, 2*i+2, n)
		}
		array[0], array[p] = array[p], array[0]
		n--
	}
	t.Logf("%v\n", array)
}

func adjust(k, l, r, n int) {
	if array[l] > array[k] {
		array[k], array[l] = array[l], array[k]
	}
	if r < n {
		if array[r] > array[k] {
			array[r], array[k] = array[k], array[r]
		}
	}
}

// 归并排序，将数组分为两列，将两列排序，然后归并为一个。递归分割，当只有一个元素时，该子序列为有序的。
func TestMergeSort(t *testing.T) {
	array = MergeSort(array)
	t.Logf("%v\n", array)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	m := len(arr) / 2
	a := MergeSort(arr[0:m])
	b := MergeSort(arr[m:])
	return Merge(a, b)
}

func Merge(a, b []int) []int {
	i, j := 0, 0
	c := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	for i < len(a) {
		c = append(c, a[i])
		i++
	}
	for j < len(b) {
		c = append(c, b[j])
		j++
	}
	return c
}

// 计数排序，找出原数组的最小值a和最大值b，创建数组s大小为(b-a+1)，遍历原数组，迭代变量为e，将s[e-min]加一，进行计数。
// 对数组s，从第1个开始，每一个位置的值增加前一个的值，表示当前下标所代表的值在结果数组中的最后一个位置。
// 遍历原数组，从s中取出该元素的位置，将其减一即为在结果数组中的位置（因为计数是从1开始的，数组是从0开始数），并将s中的值减一。
func TestCountSort(t *testing.T) {
	min := array[0]
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
		if array[i] > max {
			max = array[i]
		}
	}
	a := make([]int, (max - min + 1), (max - min + 1))
	for _, v := range array {
		a[v-min]++
	}
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
	}
	rs := make([]int, len(array), len(array))
	for _, v := range array {
		rs[a[v-min]-1] = v
		a[v-min]--
	}
	t.Logf("%v\n", rs)
}

// 基数排序，对每个元素的基数排序，第n次就按照每个元素的第n位（从右往左）大小排序。
func TestRadixSort(t *testing.T) {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	c := 0 // 需要几轮基数排序，最大数有几位就是几轮
	for max > 0 {
		c++
		max /= 10
	}

	for i := 1; i <= c; i++ {
		m := make([][]int, 10, 10)
		for _, v := range array {
			org := v
			var mod int
			for k := 0; k < i; k++ {
				mod = v % 10
				v /= 10
			}
			a := m[mod]
			if len(a) == 0 {
				a = make([]int, 0)
			}
			a = append(a, org)
			m[mod] = a
		}

		array = make([]int, 0, len(array))
		for _, v := range m {
			for _, val := range v {
				array = append(array, val)
			}
		}
	}

	t.Logf("%v\n", array)
}
