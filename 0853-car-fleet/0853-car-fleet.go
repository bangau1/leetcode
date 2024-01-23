type car struct {
    pos int
    speed int
}
func carFleet(target int, position []int, speed []int) int {
    // sort by position
    // process at the last car (that near the target), calculate the time to reach the target
    // process at the last-1 car, check whether during the last time, it bumps the last car or not. if yes, join them
    // if not bump, then set the time to reach the target to the last car
    n := len(position)
    cars := make([]car, 0)

    for i:=0;i<n;i++{
        cars = append(cars, car{
            position[i],
            speed[i],
        })
    }

    sort.Slice(cars, func(a, b int)bool {
        return cars[a].pos < cars[b].pos
    })

    i := n-1
    var lastCar = cars[i]
    fleet := 1

    // t1 := (target - pos1)/speed1
    // t2 := (target - pos2)/speed2
    // if t2 > t1, car2 won't bump to car1
    // if t2 <= t1, car2 will bump car1 until finish

    for i >= 0 {
        d1 := int64(target - lastCar.pos)
        v1 := int64(lastCar.speed)

        d2 := int64(target - cars[i].pos)
        v2 := int64(cars[i].speed)

        // if t2 > t1, wont bump, increase fleet count
        // if d2/v2 > d1/v1 --> d2 * v1 > v2 * d1
        if d2 * v1 > v2 * d1 {
            fleet++
            lastCar = cars[i]
        }else{
            // car2 bump car1, hence it counts as same car fleet, 
            // do nothing
        }
        i--
    }  

    return fleet
}


// f(t) = v[i]*t + p[i] 
// target = v[i]*t + p[i]