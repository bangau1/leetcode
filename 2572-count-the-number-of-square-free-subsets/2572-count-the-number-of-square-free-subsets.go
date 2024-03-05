type primeFactor struct {
    factors map[int]int
    mask int
    squareFree bool 
}

func squareFreeSubsets(nums []int) int {
    
    maskSize := 1 << len(primes)
    dp := make([]int, maskSize)
    dp[0] = 1 // empty set 

    var pf primeFactor
    for i:=0;i<len(nums);i++{

        pf = getPrimeFactor(nums[i])
        // skip including this number if it isn't square free
        if !pf.squareFree {
            continue
        }

        for prevMask:=0;prevMask<maskSize;prevMask++{
            if prevMask & pf.mask == 0 { // no common prime
                dp[pf.mask|prevMask] = (dp[pf.mask|prevMask] + dp[prevMask]) % MOD
            }
        }
    }
    
    total := 0
    for _, subset := range dp {
        total = (total + subset) % MOD
    }
    if total - 1 < 0 {
        return total + MOD -1
    }
    return (total-1)  % MOD
}
const MOD = 1000000000+7
// // since the mask size is smaller than the len(nums), there may be overlap
// // to calculate it correctly, we need to store the pair{idx, mask} -> total subset found at idx and given mask
// // dp[idx][0] = 1
// // dp[idx]
// func helper(idx int, nums []int, pfs []primeFactor, globalMask int, total *int) {
//     if idx == len(nums){
//         return
//     }
//     prevSubset := dp[idx][globalMask]

//     // include the idx
//     if pfs[idx].mask & globalMask == 0 && pfs[idx].squareFree {
//         *total = *total + 1
//         if prevSubset == 0 {
//             prevSubset = 1
//         }

//         dp[idx][globalMask | pfs[idx].mask] = prevSubset << 1

//         helper(idx+1, nums, pfs, globalMask | pfs[idx].mask, total)
//     }else{
//         helper(idx+1, nums, pfs, globalMask | pfs[idx].mask, total)
//     }
    
// }

// A square-free integer is an integer that is divisible by no square number other than 1.
// square number: 1, 4, 9, 16?
// square number to prime factor:
// - 1 = 1
// - 4 = 2 x 2
// - 9 = 3 x 3
// - 16 = 4 x 4 = 2^8
// - 25 = 5^2
// - 36 = 6 x 6 = 2^2 x 3^2
// - 49 = 7^2
// - 100 = 10^2 = 2^2 x 5^2
// in other word, a sqare number is = prime factorization, and each prime has >= power of 2.
// constraint:
// - nums.length <= 1000
// - nums[i] <= 30
// max product is: 30^1000
// since nums[i] <= 30, we can list down all primes below 30:
// {2,3,5,7,11,13,17,19,23,29}
var primes = []int{2,3,5,7,11,13,17,19,23,29}

// the bool = squareFree
func getPrimeFactor(val int) primeFactor {
    if val == 1 {
        return primeFactor{nil, 0, true}
    }

    res := make(map[int]int)
    mask := 0
    squareFree := true

    for i:=0;i<len(primes) && val > 1;i++{
        for val > 1 && val % primes[i] == 0 {
            res[primes[i]]++
            val /= primes[i]
            mask |= (1 << i)
            if res[primes[i]] >= 2 {
                squareFree = false
            }
        }
    }
    if val > 1 {
        res[val]++
    }
    return primeFactor{res, mask, squareFree}
}