package ak

import (
	"testing"
)
//滑动窗口
func lengthOfLongestSubstring_(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring__(s string)int{
	if len(s) == 0{
		return 0
	}
	hm := map[uint8]int{}
	result , left ,right := 0,0,-1
	for left < len(s){
		if left != 0 {
			delete(hm,s[left-1])
		}
		for right+1<len(s)&&hm[s[right+1]]==0{
			right++
			hm[s[right]]++
		}
		result = max(result,right-left+1)
		left++
	}
	return result
}


func TestXargs(t *testing.T){
	t.Log(lengthOfLongestSubstring_("cccdd"))
	t.Log(lengthOfLongestSubstring__("cccdd"))

}

