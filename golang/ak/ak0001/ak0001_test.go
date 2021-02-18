package ak

import "testing"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
//返回所有的呢?
func twoSums(nums []int, target int) [][]int {
	m := make(map[int]int)
	var res [][]int
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			res = append(res,[]int{idx,k})
		}
		m[v] = k
	}
	return res
}

func TestXargs(t *testing.T){
	testTwoSum(t,[]int{0,2,3,2,1},4)
}

func testTwoSum(t *testing.T,nums []int,target int){
	t.Log("return [][] -- > ",twoSums(nums,target)) //返回多个符合值
	t.Log("return [] --> ",twoSum(nums,target)) //返回单个
}
