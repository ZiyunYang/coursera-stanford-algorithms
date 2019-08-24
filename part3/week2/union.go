package main

func find(p int, clusterID []int) int {
	for {
		if p == clusterID[p] {
			return p
		}
		clusterID[p] = clusterID[clusterID[p]]
		p = clusterID[p]
	}

}

func union(p, q, count int, sz, clusterID []int) int{
	i := find(p, clusterID)
	j := find(q, clusterID)
	if i == j {
		return count
	}
	if sz[i] < sz[j] {
		clusterID[i] = j
		sz[j] += sz[i]
	} else {
		clusterID[j] = i
		sz[i] += sz[j]
	}
	count--
	return count


}