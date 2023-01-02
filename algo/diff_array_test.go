package algo

import (
    "testing"
)

func TestDiffArray_PrefixSum(t *testing.T) {
    da := &DiffArray{}
    res := da.PrefixSum([]int{2,5,6,1})
    t.Log(res)
    for _, ran := range [][]int{{0,0}, {1,3}, {3,2}, {2, 4}} {
        ans, err := da.QueryPrefixSum(ran[0], ran[1])
        if err != nil {
            t.Error(ran[0], ran[1], err)
            continue
        }
        t.Log(ran[0], ran[1], ans)
    }
}

func TestDiffArray_InitDiff(t *testing.T) {
    da := &DiffArray{}
    // init array which all elements value is 0
    initArr := make([]int, 5)
    da.InitDiff(initArr)
    da.ActionDiff("+", 2,1, 3)
    t.Log("current-diff-num: ", da.GetDiffNum())
    t.Log("result-diff-num", da.ResultDiff())
}
