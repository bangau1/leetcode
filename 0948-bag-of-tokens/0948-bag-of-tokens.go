func bagOfTokensScore(tokens []int, power int) int {
    // face up: power >= tokens[i], power-=tokens[i], score+=1
    // face down: score >= 1, power += tokens[i], score -= 1
    // we want to maximize the score
    // - face up -> to increase the score, but reduce power
    // - face down -> to prolong the power, but reduce the score
    // 
    // basically it's a problem to maximize the number of token being face up
    
    sort.Ints(tokens)
    score := 0
    maxScore := 0
    for len(tokens) > 0 {
        
        if power >= tokens[0]{
            power -= tokens[0]
            score += 1
            tokens = tokens[1:]

            if maxScore < score {
                maxScore = score
            }
        }else if score  >= 1 {
            power += tokens[len(tokens)-1]
            score -= 1
            tokens = tokens[0:len(tokens)-1]
        }else{
            break
        }
    }

    return maxScore
}