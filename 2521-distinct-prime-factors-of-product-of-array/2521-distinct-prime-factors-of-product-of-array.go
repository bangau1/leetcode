func distinctPrimeFactors(nums []int) int {
    res := make(map[int]int)
    for _, num := range nums {
        addDistinctPrimeFactor(num, res)
    }
    return len(res)
}

func addDistinctPrimeFactor(val int, res map[int]int)  {
    for i:=0;i<len(primes) && val > 1; i++{
        for val > 1 && val % primes[i] == 0 {
            res[primes[i]]++
            val /= primes[i]
        }
    }
    if val > 1 {
        res[val]++
    }
}

var primes []int

func init(){
    generatePrime()
}

func generatePrime() {
    primes = make([]int, 0)
    // list primes <= 1000
    isPrime := make([]bool, 1001)
    for i:=0;i<=1000;i++{
        isPrime[i] = true
    }
    
    for i:=2;i*i<=1000;i++{
        if !isPrime[i] {
            continue
        }
        primes = append(primes, i)
        for p:=2*i;p>=1000;p+=i{
            isPrime[p] = false
        }
    }
}