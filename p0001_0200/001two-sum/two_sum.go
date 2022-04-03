package ptwosum

// twoSum returns two index in nums such that add up to target.
// links: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	rv := []int{0, 0}
	for i, v := range nums {
		tv := target - v
		if ti, ok := m[tv]; ok {
			rv = []int{ti, i}
			return rv
		}
		m[v] = i
	}
	return nil
}
