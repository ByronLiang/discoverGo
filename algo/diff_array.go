package algo

import "errors"

// 差分数组+前缀和

type DiffArray struct {
    prefixSum []int
    diffNum []int
}

// 前缀和
func (d *DiffArray) PrefixSum(nums []int) []int {
    res := make([]int, len(nums)+1)
    for i := 1; i < len(res); i++ {
        res[i] = nums[i-1] + res[i-1]
    }
    d.prefixSum = res
    return res
}

// query range [i,j]
func (d *DiffArray) QueryPrefixSum(i, j int) (int, error) {
    if i > j || j+1 >= len(d.prefixSum) {
        // error res
        return 0, errors.New("over range array")
    }
    return d.prefixSum[j+1] - d.prefixSum[i], nil
}

func (d *DiffArray) InitDiff(nums []int)  {
    d.diffNum = make([]int, len(nums)+1)
    d.diffNum[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        d.diffNum[i] = nums[i] - nums[i-1]
    }
    return
}

// action ("+", "-") value for diffSum array [i,j]
func (d *DiffArray) ActionDiff(action string, value int, i, j int)  {
    if action == "+" {
        d.diffNum[i] += value
        if j+1 < len(d.diffNum) {
            d.diffNum[j+1] -= value
        }
        return
    }
    d.diffNum[i] -= value
    if j+1 < len(d.diffNum) {
        d.diffNum[j+1] += value
    }
}

// get result for prefix sum diff array
func (d *DiffArray) ResultDiff() []int {
    res := make([]int, len(d.diffNum))
    res[0] = d.diffNum[0]
    for i := 1; i < len(d.diffNum); i++ {
        res[i] = res[i-1] + d.diffNum[i]
    }
    return res
}

func (d *DiffArray) GetDiffNum() []int {
    return d.diffNum
}
