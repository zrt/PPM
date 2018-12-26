package PPM

import (
	"container/list"
	"fmt"
)

type HashTable struct {
	hash []*list.List
	N    int
	r    float64
}

func NewHashTable(hp *list.List, num int, r float64) *HashTable {
	r *= 2
	t := &HashTable{make([]*list.List, 2*num), 2 * num, r}
	for i := 0; i < t.N; i++ {
		t.hash[i] = list.New()
	}
	one := *NewV(1, 1, 1)
	bar := Pbar{}
	bar.Init(num)
	fmt.Println("gen hashtable...")
	for e := hp.Front(); e != nil; e = e.Next() {
		hpp := e.Value.(*HPoint)
		p := hpp.pos
		mn := p.Div(r).Sub(one)
		mx := p.Div(r).Add(one)
		for i := int(mn.X); i <= int(mx.X); i++ {
			for j := int(mn.Y); j <= int(mx.Y); j++ {
				for k := int(mn.Z); k <= int(mx.Z); k++ {
					t.insert(i, j, k, hpp)
				}
			}
		}
		bar.Tick()
	}
	fmt.Println()
	return t
}

func (t *HashTable) GetTable(pos V) *list.List {
	return t.hash[t.get(pos)]
}

func (t *HashTable) insert(i, j, k int, hp *HPoint) {
	idx := t.gett(i, j, k)
	t.hash[idx].PushFront(hp)
}

func (t *HashTable) get(v V) int {
	x := v.Div(t.r)
	return t.gett(int(x.X), int(x.Y), int(x.Z))
}

func (t *HashTable) gett(x, y, z int) int {
	sum := uint((x * 73856093) ^ (y * 19349663) ^ (z * 83492791))
	return int(sum % uint(t.N))
}
