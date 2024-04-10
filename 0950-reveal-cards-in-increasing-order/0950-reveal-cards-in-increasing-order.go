func deckRevealedIncreasing(deck []int) []int {
/*
general idea:
- sort the deck
- take the biggest one
- reverse the operation

example: [2,3,5,7,11,13,17]
- [17]
- [13,17]
- [11,17,13] -> put 11 at first element, move last to second
- [7,13,11,17] -> put 7 at first, move last to second
- [5,17,7,13,11]
- [3,11,5,17,7,13]
- [2,13,3,11,5,17,7,13]
*/
    var res []int
    sort.Ints(deck)
    res = append(res, deck[len(deck)-1])
    var top int
    for i:=len(deck)-2;i>=0;i--{
        top = res[len(res)-1]
        res = append([]int{deck[i], top}, res[0:len(res)-1]...)
    }
    return res
}
