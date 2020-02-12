type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: make([]int, 0, 10000),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	for len(this.pings) > 0 && this.pings[0]+3000 < t {
		this.pings = this.pings[1:]
	}

	return len(this.pings)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
