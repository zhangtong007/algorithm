package leetcode2671

type FrequencyTracker struct {
	// frequency : count
	frequency map[int]int
	// value: count
	values map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		frequency: map[int]int{},
		values:    map[int]int{},
	}
}

func (this *FrequencyTracker) Add(number int) {
	this.frequency[this.values[number]]--
	this.values[number]++
	this.frequency[this.values[number]]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.values[number] == 0 {
		return
	}
	this.frequency[this.values[number]]--
	this.values[number]--
	this.frequency[this.values[number]]++
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.frequency[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
