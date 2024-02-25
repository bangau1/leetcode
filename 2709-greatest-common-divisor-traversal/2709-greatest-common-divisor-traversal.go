func generatePrimes(n int) []int {
    primes := make([]int, 0)
    isPrime := make([]bool, n+1)
    for i:=0;i<=n;i++{
        isPrime[i] = true
    }
    for i:=2;i*i<=n;i++{
        if !isPrime[i] {
            continue
        }
        
        for p:=i<<1;p<=n;p+=i{
            isPrime[p] = false
        }

    }
    for i:=2;i<=n;i++{
        if isPrime[i]{
            primes = append(primes, i)
        }
    }
    return primes
}
var PRIMES []int
func init() {
    if len(PRIMES) == 0 {
        PRIMES = generatePrimes(10000)
    }
}

func computePrimeFactor(n int) []int {
    c := n
    var factor []int
    for i:=0;i<len(PRIMES) && PRIMES[i]*PRIMES[i] <= n && c > 1;i++{
        if c % PRIMES[i] == 0 {
            factor = append(factor, PRIMES[i])
            for c > 1 && c % PRIMES[i] == 0 {
                c /= PRIMES[i]
            }
        }
        
    }
    if c > 1 {
        factor = append(factor, c)
    }
    return factor
}

func canTraverseAllPairs(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return true
    }

    parents := make([]int, n)
    for i:=0;i<n;i++{
        parents[i] = -1
    }

    primeToLastIdx := make([]int, 10000)
    for i:=0;i<len(primeToLastIdx);i++{
        primeToLastIdx[i] = -1
    }
    setLastIdx := func(idx int) {
        primeFactors := computePrimeFactor(nums[idx])
        for _, factor := range primeFactors {
            if primeToLastIdx[factor] == -1 {
                primeToLastIdx[factor] = idx
            }else{
                union(parents, primeToLastIdx[factor], idx)
                primeToLastIdx[factor] = idx
            }
        }

    }
    for i:=0;i<n;i++{
        if nums[i] == 1 {
            return false
        }
        setLastIdx(i)
    }

    for i:=0;i<n;i++{
        if parents[i] < 0 && -parents[i] == n {
            return true
        }
    }
    return false
   
}


func find(parents []int, x int) int {
	if parents[x] < 0 {
		return x
	}

	path := make([]int, 0)

	for parents[x] >= 0 {
		path = append(path, x)
		x = parents[x]
	}
	for _, p := range path {
		parents[p] = x
	}
	return x
}

func union(parents []int, a, b int) {
	if a == b {
		return
	}
	aId := find(parents, a)
	bId := find(parents, b)

	if aId == bId {
		return
	}

	aSize := -parents[aId]
	bSize := -parents[bId]

	if aSize > bSize {
		parents[bId] = aId
		parents[aId] = -aSize - bSize
	} else {
		parents[aId] = bId
		parents[bId] = -aSize - bSize
	}
}