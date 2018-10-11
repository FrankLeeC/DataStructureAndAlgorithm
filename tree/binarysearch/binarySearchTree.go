package main

import (
	"fmt"
	"strconv"
)

var array = []int{81, 94, 11, 96, 12, 35, 17, 95, 28, 58, 41, 75, 15, 27}

func main() {
	t := create()
	walk(t, "")
	b, _ := search(75, t)
	fmt.Println(b)
	delete(35, t)
	fmt.Println("---------------")
	walk(t, "")
	fmt.Println("---------------")
	fmt.Println(findMin(t))
	fmt.Println(findMax(t))
}

type tree struct {
	data int
	p    *tree
	l    *tree
	r    *tree
}

// 创建
func create() *tree {
	if len(array) <= 0 {
		return nil
	}
	t := new(tree)
	t.data = array[0]
	for _, v := range array {
		append(v, t)
	}
	return t
}

func append(data int, t *tree) {
	if data < t.data {
		if t.l == nil {
			tmp := new(tree)
			tmp.data = data
			t.l = tmp
			tmp.p = t
		} else {
			append(data, t.l)
		}
		append(data, t.l)
	} else if data > t.data {
		if t.r == nil {
			tmp := new(tree)
			tmp.data = data
			t.r = tmp
			tmp.p = t
		} else {
			append(data, t.r)
		}
	}
}

// 遍历
func walk(t *tree, s string) {
	if t != nil {
		fmt.Printf("%s%d\n", s, t.data)
		walk(t.l, s+strconv.Itoa(t.data)+"<")
		walk(t.r, s+strconv.Itoa(t.data)+">")
	}
}

// 查询
func search(data int, t *tree) (bool, *tree) {
	if t != nil {
		if t.data == data {
			return true, t
		}
		if t.data < data {
			return search(data, t.r)
		}
		return search(data, t.l)
	}
	return false, nil
}

// 删除
func delete(data int, t *tree) bool {
	if t != nil {
		if t.data == data {
			t.data = deleteMax(t.l).data
		} else if t.data < data {
			return delete(data, t.r)
		}
		return delete(data, t.l)
	}
	return false
}

func deleteMax(t *tree) *tree {
	if t.r == nil {
		t.p.r = t.l
		return t
	}
	return deleteMax(t.r)
}

// 查找最小值
func findMin(t *tree) int {
	if t.l == nil {
		return t.data
	}
	return findMin(t.l)
}

// 查找最大值
func findMax(t *tree) int {
	if t.r == nil {
		return t.data
	}
	return findMax(t.r)
}
