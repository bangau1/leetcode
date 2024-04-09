func timeRequiredToBuy(tickets []int, k int) int {
    n := len(tickets)
    if k < n-1 {
        ticketsAfterK := 0
        for i:=k+1;i<n;i++{
            ticketsAfterK += tickets[i]
        }
        tickets = tickets[0:k+1]
        tickets = append(tickets, ticketsAfterK)
    }

    if k > 0 {
        ticketsBeforeK := 0
        for i:=0;i<k;i++{
            ticketsBeforeK += tickets[i]
        }
        tickets = tickets[k:]
        tickets = append([]int{ticketsBeforeK}, tickets...)
    }
    fmt.Println("after", tickets)

    kIdx := k
    time := 0
    tmp := -1
    for len(tickets) > 0 {
        if tickets[0] > 0 {
            time += 1
            tickets[0] -= 1
            
                tmp = tickets[0]
                
            if tickets[0] == 0 {
                tickets = tickets[1:]
            }else{
                tickets = append(tickets, tmp)
            }
        }else{

        }
    }
    return time
}