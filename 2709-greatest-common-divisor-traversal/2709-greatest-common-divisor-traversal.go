func generatePrimes(n int) []int {
    primes := make([]int, 0)
    isPrime := make([]bool, n+1)
    for i:=0;i<=n;i++{
        isPrime[i] = true
    }
    isPrime[3] = true
    for i:=3;i*i<=n;i+=2{
        if !isPrime[i] {
            continue
        }
        
        for p:=i*i;p<=n;p+=i{
            isPrime[p] = false
        }

    }
    primes = append(primes, 2)
    for i:=3;i<=n;i++{
        if i&1==1 && isPrime[i] {
            primes = append(primes, i)
        }
    }
    return primes
}
var PRIMES []int
var PRIME_FACTORS [][]int

func init() {
    if len(PRIMES) == 0 {
        PRIMES = generatePrimes(320)
    }
    PRIME_FACTORS = make([][]int, 100001)
    for i:=2;i<=100000;i++{
        PRIME_FACTORS[i] = computePrimeFactor(i)
    }
}

func computePrimeFactor(n int) []int {
    c := n
    var factor []int
    for i:=0;i<len(PRIMES) && c > 1 && PRIMES[i]*PRIMES[i] <= n ;i++{
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
        if nums[i] == 1 {
            return false
        }
    }

    primeToLastIdx := make(map[int]int)
    setLastIdx := func(idx int) {
        primeFactors := PRIME_FACTORS[nums[idx]]
        for _, factor := range primeFactors {
            if _, ok:=primeToLastIdx[factor]; !ok {
                primeToLastIdx[factor] = idx
            }else{
                union(parents, primeToLastIdx[factor], idx)
                primeToLastIdx[factor] = idx
            }
        }

    }
    for i:=0;i<n;i++{
        setLastIdx(i)
    }

    unionId := find(parents, 0)
    return -parents[unionId] == n
   
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