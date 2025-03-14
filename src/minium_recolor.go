package main

//Bài 2379
//https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/?envType=daily-question&envId=2025-03-08

func minimumRecolors(blocks string, k int) int {
    blackCount, ans := 0, k
    
    for i := 0; i < len(blocks); i++ {
        if i >= k && blocks[i-k] == 'B' {
            blackCount--
        }
        if blocks[i] == 'B' { 
            blackCount++
        }
        if ans > k-blackCount {
            ans = k-blackCount
        }
    }
    
    return ans
}
