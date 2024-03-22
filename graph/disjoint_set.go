package graph

type DisjointSet struct {
	parents []int
}

func NewDisjointSet(n int) *DisjointSet {
	res := DisjointSet{
		parents: make([]int, n),
	}

	for i := 0; i < n; i++ {
		res.parents[i] = -1 // initialize the union size to 1 (negative value to store the size of the union, while also saying that it's a union leaderid)
	}

	return &res
}

func (d *DisjointSet) Find(x int) int {
	if d.parents[x] < 0 {
		return x
	}

	path := make([]int, 0)

	for d.parents[x] >= 0 {
		path = append(path, x)
		x = d.parents[x]
	}

	for _, p := range path {
		d.parents[p] = x
	}

	return x
}

func (d *DisjointSet) Union(a, b int) bool {
	if a == b {
		return false
	}

	aId, bId := d.Find(a), d.Find(b)

	if aId == bId {
		return false
	}

	aSize, bSize := d.GetUnionSize(aId), d.GetUnionSize(bId)

	if bSize > aSize {
		d.parents[aId] = bId
		d.parents[bId] = -(aSize + bSize)
	} else {
		d.parents[bId] = aId
		d.parents[aId] = -(aSize + bSize)
	}
	return true
}

func (d *DisjointSet) GetUnionSize(a int) int {
	return -d.parents[d.Find(a)]
}

func (d *DisjointSet) GetTotalUnions() int {
	var count int
	for i := 0; i < len(d.parents); i++ {
		if d.parents[i] < 0 {
			count += 1
		}
	}
	return count
}
