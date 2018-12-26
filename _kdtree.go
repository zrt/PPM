package PPM

import (
	"container/list"
)

type KDTree struct {
	hp     []*HPoint
	ls, rs []int
	mn, mx []V
	root   int
}

func NewKDTree(lst *list.List, num int) *KDTree {
	tree := KDTree{make([]*HPoint, num), make([]int, num), make([]int, num), make([]V, num), make([]V, num), -1}
	tmp := lst.Front()
	for i := 0; i < num; i++ {
		tree.hp[i] = tmp.Value.(*HPoint)
		tree.ls[i] = -1
		tree.rs[i] = -1
		tree.mn[i] = tree.hp[i].pos
		tree.mx[i] = tree.hp[i].pos
		tmp = tmp.Next()
	}
	tree.root = tree.Build(0, num-1, 0)
	return &tree
}

// 参数左闭右闭区间
func (t *KDTree) Build(l, r, d int) int {
	if l > r {
		return -1
	}
	mid := (l + r) >> 1
	t.nth(l, mid, r+1, d)
	t.ls[mid] = t.Build(l, mid-1, (d+1)%3)
	t.rs[mid] = t.Build(mid+1, r, (d+1)%3)
	if t.ls[mid] != -1 {
		t.upd(mid, t.ls[mid])
	}
	if t.rs[mid] != -1 {
		t.upd(mid, t.rs[mid])
	}
	return mid
}

func (t *KDTree) upd(x, y int) {
	t.mn[x].Min_(t.mn[y])
	t.mx[x].Max_(t.mx[y])
}

// 左闭右开区间
func (t *KDTree) nth(l, mid, r, d int) {
	//if l == r {
	//	return
	//}
	//pos := rand.Intn(r-l) + l

}
