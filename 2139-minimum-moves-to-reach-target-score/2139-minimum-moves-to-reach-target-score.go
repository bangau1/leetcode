func minMoves(target int, maxDoubles int) int {
    moves := 0
    double := 0
    for target > 1 {
        if double < maxDoubles {
            if target % 2 == 1 {
                target -= 1
                moves += 1
            }
            target = target/2
            moves+= 1
            double++
        }else{
            moves += target-1
            target = 1
        }
        
    }
    return moves
}