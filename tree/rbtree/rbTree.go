package main

import (
	"fmt"
)

var array = []int{81, 94, 11, 96, 12, 35, 17, 95, 28, 58, 41, 75, 15, 27, 36}

var black, red = true, false

var tree *rbtree

type rbtree struct {
	data int
	c    bool    // true -> black  false -> red
	p    *rbtree // parent
	l    *rbtree // left child
	r    *rbtree // rigth child
}

func main() {
	tree = root(tree)
	for _, v := range array {
		if tree == nil {
			tree = new(rbtree)
			tree.data = v
			insertNode(tree, nil)
		} else {
			n := new(rbtree)
			n.c = red
			n.data = v
			insertNode(n, tree)
		}
		tree = root(tree)
		walk(tree, "")
		fmt.Println("------------------")
	}
	fmt.Println("=====delete=====")
	// deleteData(28, tree)

	// deleteData(75, tree)

	// deleteData(81, tree)

	// deleteData(11, tree)

	deleteData(15, tree)
	deleteData(27, tree)
	deleteData(11, tree)
	tree = root(tree)
	walk(tree, "")
}

func insertNode(n *rbtree, t *rbtree) {
	if t == nil { // case 1
		insertCase1(tree)
	} else {
		p := findParent(n, t)
		if p.c == black { // case 2
			insertCase2(n, p)
		} else { // p.c == red
			u := sibling(p)
			if u != nil && u.c == red { // case 3
				insertCase3(n, p)
			} else { // case 4    u.c == black or u == nil
				insertCase4(n, p)
			}
		}
	}
}

func findParent(n *rbtree, t *rbtree) *rbtree {
	if n == t {
		return nil
	}
	if t.data > n.data {
		if t.l == nil || t.l.data == n.data {
			return t
		}
		return findParent(n, t.l)
	}
	if t.r == nil || t.r.data == n.data {
		return t
	}
	return findParent(n, t.r)
}

func insertCase1(n *rbtree) {
	tree.c = black
}

func insertCase2(n *rbtree, p *rbtree) {
	n.p = p
	if p.data > n.data {
		p.l = n
	} else {
		p.r = n
	}
}

func insertCase3(n *rbtree, p *rbtree) {
	u := sibling(p)
	g := parent(p)
	n.p = p
	if p.data > n.data {
		p.l = n
	} else {
		p.r = n
	}
	g.c = red
	p.c = black
	u.c = black
	p = findParent(g, tree)
	insertNode(g, p)
}

func insertCase4(n *rbtree, p *rbtree) {
	g := parent(p)
	if n.data < p.data && g.l == p {
		p.l = n
		n.p = p
		rightRotate(g)
		g.c = red
		p.c = black
	} else if n.data >= p.data && g.r == p {
		p.r = n
		n.p = p
		leftRotate(g)
		g.c = red
		p.c = black
	} else if n.data >= p.data && g.l == p {
		p.r = n
		n.p = p
		leftRotate(p)
		rightRotate(g)
		n.c = black
		g.c = red
	} else {
		p.l = n
		n.p = p
		rightRotate(p)
		leftRotate(g)
		n.c = black
		g.c = red
	}
}

func root(t *rbtree) *rbtree {
	if t == nil {
		return nil
	}
	if t.p == nil {
		return t
	}
	return root(t.p)
}

func parent(t *rbtree) *rbtree {
	return t.p
}

func grandParent(t *rbtree) *rbtree {
	if p := parent(t); p != nil {
		return parent(p)
	}
	return nil
}

func sibling(t *rbtree) *rbtree {
	if p := parent(t); p != nil {
		if p.l == t {
			return p.r
		}
		return p.l
	}
	return nil
}

func uncle(t *rbtree) *rbtree {
	if g := grandParent(t); g != nil {
		p := parent(t)
		if g.l == p {
			return g.r
		}
		return g.l
	}
	return nil
}

func leftRotate(t *rbtree) {
	p := parent(t)
	if p == nil {
		r := t.r
		t.r = nil
		r.p = nil
		if r.l != nil {
			t.r = r.l
			r.l.p = t
		}
		t.p = r
		r.l = t
	} else {
		if p.l == t {
			tmp := t.r.l
			p.l = t.r
			t.r.p = p
			p.l.l = t
			t.p = p.l
			t.r = nil
			if tmp != nil {
				t.r = tmp
				tmp.p = t
			}
		} else {
			tmp := t.r.l
			p.r = t.r
			t.r.p = p
			p.r.l = t
			t.p = p.r
			t.r = nil
			if tmp != nil {
				t.r = tmp
				tmp.p = t
			}
		}
	}

}

func rightRotate(t *rbtree) {
	p := parent(t)
	if p == nil {
		l := t.l
		t.l = nil
		l.p = nil
		if l.r != nil {
			t.l = l.r
			l.r.p = t
		}
		t.p = l
		l.r = t
	} else {
		if p.l == t {
			tmp := t.l.r
			p.l = t.l
			t.l.p = p
			p.l.r = t
			t.p = p.l
			t.l = nil
			if tmp != nil {
				t.l = tmp
				tmp.p = t
			}
		} else {
			tmp := t.l.r
			p.r = t.l
			t.l.p = p
			p.r.r = t
			t.p = p.r
			t.l = nil
			if tmp != nil {
				t.l = tmp
				tmp.p = t
			}

		}
	}
}

func walk(t *rbtree, s string) {
	if t != nil {
		fmt.Printf("%s%d(%v)\n", s, t.data, t.c)
		walk(t.l, fmt.Sprintf("%d(%v)%s", t.data, t.c, "<"))
		walk(t.r, fmt.Sprintf("%d(%v)%s", t.data, t.c, ">"))
	}
}

func search(data int, t *rbtree) *rbtree {
	if t == nil {
		return nil
	}
	if t.data == data {
		return t
	}
	if t.data > data {
		return search(data, t.l)
	}
	return search(data, t.r)
}

// 删除
func deleteData(data int, t *rbtree) {
	d := search(data, t)
	if p := parent(d); p == nil && d.l == nil && d.r == nil { // 删除的是唯一一个结点
		tree = nil
		return
	}
	if d.r != nil {
		x := findRightMin(d)
		d.data = x.data
		deleteNode(x)
	} else if d.l != nil {
		x := findLeftMax(d)
		d.data = x.data
		deleteNode(x)
	} else { // 要删除的结点没有非空子结点
		deleteNode(d)
	}
}

func findRightMin(t *rbtree) *rbtree {
	r := t.r
	for r.l != nil {
		r = r.l
	}
	return r
}

func findLeftMax(t *rbtree) *rbtree {
	l := t.l
	for l.r != nil {
		l = l.r
	}
	return l
}

func removeNode(x *rbtree) {
	if p := parent(x); p != nil {
		if p.l == x {
			p.l = nil
		} else {
			p.r = nil
		}
	}
}

func deleteNode(x *rbtree) {
	if x.c == red {
		deleteCase1(x)
		return
	}
	if x.c == black && (x.l != nil || x.r != nil) {
		deleteCase2(x)
		return
	}
	deleteCase3(x)
}

// 要删除的结点是红色结点。
// 此时一定没有（非空）子结点。
// 假如有一个红色子结点，违背性质4，假如有一个黑节点，违背性质5。
// 一定不是根结点
func deleteCase1(x *rbtree) {
	removeNode(x)
	x = nil
}

// 要删除的结点是黑色结点。
// 有一个（非空）子结点，且一定是红色，如果是黑色，则原树违背性质5。
func deleteCase2(x *rbtree) {
	var c *rbtree
	if x.l != nil {
		c = x.l
	}
	if x.r != nil {
		c = x.r
	}
	if x.p.l == x {
		x.p.l = c
	} else {
		x.p.r = c
	}
	c.p = x.p
	c.c = black
	x = nil
}

func deleteCase3(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if s.c == red { // case1
		deleteCase3_1(x)
		removeNode(x)
		x = nil
		return
	}
	if p.c == red { // case2-5
		if s.l == nil && s.r == nil {
			deleteCase3_2(x)
			removeNode(x)
			x = nil
			return
		}
		if s.l != nil && s.r == nil {
			deleteCase3_3(x)
			removeNode(x)
			x = nil
			return
		}
		if s.r != nil && s.l == nil {
			deleteCase3_4(x)
			removeNode(x)
			x = nil
			return
		}
		if s.r != nil && s.l != nil {
			deleteCase3_5(x)
			removeNode(x)
			x = nil
			return
		}
	}
	if p.c == black { // case6-9   p is black
		if s.l != nil && s.r == nil {
			deleteCase3_6(x)
			removeNode(x)
			x = nil
			return
		}
		if s.r != nil && s.l == nil {
			deleteCase3_7(x)
			removeNode(x)
			x = nil
			return
		}
		if s.l != nil && s.r != nil {
			deleteCase3_8(x)
			removeNode(x)
			x = nil
			return
		}
		if s.l == nil && s.r == nil {
			deleteCase3_9(x)
			removeNode(x)
			x = nil
			return
		}
	}
}

// S为红色。
// 由于性质4，P必须为黑色，又由于性质5，一定有SL，SR且必须为黑色， SL,SR一定没有非空子结点。
func deleteCase3_1(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		sl := s.l
		leftRotate(p)
		s.c = black
		sl.c = red
	} else {
		sr := s.r
		rightRotate(p)
		s.c = black
		sr.c = red
	}

}

// 当S没有（非空）子结点时。且P是红色。
func deleteCase3_2(x *rbtree) {
	p := parent(x)
	if p.l == x {
		leftRotate(p)
	} else {
		rightRotate(p)
	}
}

// 当S左子结点非空时。且P是红色。
// 由于性质5，SL必须为红色。
func deleteCase3_3(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		rightRotate(s)
		leftRotate(p)
		p.c = black
	} else {
		leftRotate(s)
		rightRotate(p)
		p.c = black
	}
}

// 当S右子结点非空时。且P是红色。
// 由于性质5，SR必须为红色。
func deleteCase3_4(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		sr := s.r
		leftRotate(p)
		p.c = black
		s.c = red
		sr.c = black
	} else {
		sl := s.l
		rightRotate(p)
		p.c = black
		s.c = red
		sl.c = black
	}
}

// 当S结点有两个非空子结点时。且P是红色。
// 由于性质5，两个都必须为黑色。
func deleteCase3_5(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		sr := s.r
		leftRotate(p)
		p.c = black
		s.c = red
		sr.c = black
	} else {
		sl := s.l
		rightRotate(p)
		p.c = black
		s.c = red
		sl.c = black
	}
}

// 当S有一个左子结点时。且P是黑色。
// S的左子结点一定是红色，否则违背性质5。
func deleteCase3_6(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		sl := s.l
		rightRotate(s)
		leftRotate(p)
		sl.c = black
	} else {
		sr := s.r
		leftRotate(s)
		rightRotate(p)
		sr.c = black
	}
}

// 当S有一个左子结点时。且P是黑色。
// S的右子结点一定是红色，否则违背性质5。
func deleteCase3_7(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		sr := s.r
		leftRotate(p)
		sr.c = black
	} else {
		sl := s.l
		rightRotate(p)
		sl.c = black
	}

}

// 当S有两个子结点时。且P是黑色。
// 由于性质5，S的两个结点一定都是红色。
func deleteCase3_8(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	if p.l == x {
		sr := s.r
		leftRotate(p)
		sr.c = black
	} else {
		sl := s.l
		rightRotate(p)
		sl.c = black
	}
}

// 当P,S均为黑色，且S没有（非空）子结点时。
func deleteCase3_9(x *rbtree) {
	p := parent(x)
	s := sibling(x)
	g := grandParent(x)
	s.c = red
	if g == nil { // P是根节点
		deleteCase3_9_1(x)
		return
	}
	if g.c == red { // case3_9_2_1  (case3_9_2_1_1-case3_9_2_1_4)
		u := uncle(x)
		if u.l.c == black && u.r.c == black {
			deleteCase3_9_2_1_1(x)
			return
		}
		if u.l.c == red && u.r.c == red {
			deleteCase3_9_2_1_2(x)
			return
		}
		if u.l.c == red && u.r.c == black {
			deleteCase3_9_2_1_3(x)
			return
		}
		if u.l.c == black && u.r.c == red {
			deleteCase3_9_2_1_4(x)
			return
		}
	}
	if g.c == black { // case3_9_2_2 (case3_9_2_2_1,case3_9_2_2_2)
		u := uncle(x)
		if u.c == black {
			if u.l.c == black && u.r.c == black {
				u.c = red
				deleteCase3_9(p) // 递归
				return
			}
			if p.l == x {
				if u.l.c == red && u.r.c == black {
					deleteCase3_9_2_2_1_2(x)
					return
				}
				if u.r.c == red {
					deleteCase3_9_2_2_1_3(x)
					return
				}
			}
			if p.r == x {
				if u.l.c == black && u.r.c == red {
					deleteCase3_9_2_2_1_2(x)
					return
				}
				if u.l.c == red {
					deleteCase3_9_2_2_1_3(x)
					return
				}
			}
		}
		if u.c == red {
			deleteCase3_9_2_2_2(x)
		}
	}

}

// P是根节点。
func deleteCase3_9_1(x *rbtree) {
	return
}

// 祖父结点G为红色。
// U结点一定是黑色
// 如果UL, UR也是黑色
func deleteCase3_9_2_1_1(x *rbtree) {
	g := grandParent(x)
	p := parent(x)
	if p.l == x {
		leftRotate(g)
	} else {
		rightRotate(g)
	}
}

// 祖父结点G为红色。
// U结点一定是黑色
// 如果UL，UR都是红色
// 由于性质4，ULL,ULR,URL,URR都必须是黑色
// 由于性质5，ULL,ULR,URL,URR都没有（非空）子结点。
func deleteCase3_9_2_1_2(x *rbtree) {
	g := grandParent(x)
	p := parent(x)
	u := uncle(x)
	if p.l == x {
		ul := u.l
		leftRotate(g)
		ul.c = black
		ul.l.c = red
		ul.r.c = red
	} else {
		ur := u.r
		rightRotate(g)
		ur.c = black
		ur.l.c = red
		ur.r.c = red
	}
}

// 祖父结点G为红色。
// U结点一定是黑色
// 如果UL为红色，UR为黑色
// 由于性质4，ULL,ULR必须是黑色
// 由于性质5，UR,ULL,ULR都没有（非空）子结点。
func deleteCase3_9_2_1_3(x *rbtree) {
	g := grandParent(x)
	u := uncle(x)
	p := parent(x)
	if p.l == x {
		ul := u.l
		leftRotate(g)
		ul.c = black
		ul.l.c = red
		ul.r.c = red
	} else {
		rightRotate(g)
	}
}

// 祖父结点G为红色。
// U结点一定是黑色
// 如果UL为黑色，UR为红色
// 由于性质4，URL,URR一定为黑色
// 由于性质5,UL，URL，URR都没有（非空）子结点。
func deleteCase3_9_2_1_4(x *rbtree) {
	g := grandParent(x)
	p := parent(x)
	u := uncle(x)
	if p.l == x {
		leftRotate(g)
	} else {
		ur := u.r
		rightRotate(g)
		ur.c = black
		ur.l.c = red
		ur.r.c = red
	}
}

// 祖父结点G为黑色。
// 如果U为黑色，UR为黑色，UL为红色
// 由于性质4和性质5，ULL,ULR,UR都是黑色结点，且没有（非空）子结点。
func deleteCase3_9_2_2_1_2(x *rbtree) {
	u := uncle(x)
	g := grandParent(x)
	p := parent(x)
	if p.l == x {
		ul := u.l
		rightRotate(u)
		leftRotate(g)
		ul.c = black
	} else {
		ur := u.r
		leftRotate(u)
		rightRotate(g)
		ur.c = black
	}
}

// 祖父结点G为黑色。
// 如果U为黑色，UR是红色。UL红黑都可以
func deleteCase3_9_2_2_1_3(x *rbtree) {
	g := grandParent(x)
	u := uncle(x)
	p := parent(x)
	if p.l == x {
		ur := u.r
		leftRotate(g)
		ur.c = black
	} else {
		ul := u.l
		rightRotate(g)
		ul.c = black
	}
}

// 祖父结点G为黑色。
// 如果U为红色
func deleteCase3_9_2_2_2(x *rbtree) {
	g := grandParent(x)
	u := uncle(x)
	p := parent(x)
	if p.l == x {
		ul := u.l
		leftRotate(g)
		u.c = black
		ul.c = red
		if ul.l.c == red {
			ul.l.c = black
			ul.l.l.c = red
			ul.l.r.c = red
		}
		if ul.r.c == red {
			ul.r.c = black
			ul.r.l.c = red
			ul.r.r.c = red
		}
	} else {
		ur := u.r
		rightRotate(g)
		u.c = black
		ur.c = red
		if ur.l.c == red {
			ur.l.c = black
			ur.l.l.c = red
			ur.l.r.c = red
		}
		if ur.r.c == red {
			ur.r.c = black
			ur.r.l.c = red
			ur.r.r.c = red
		}
	}
}
