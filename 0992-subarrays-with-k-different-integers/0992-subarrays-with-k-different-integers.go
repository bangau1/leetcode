import (
    "container/list"
)

var counter = make([]int, 20000+1)

/*
general idea:
- need to maintain a counter map (or use array as hashmap and additional countDist to track the distinct count)
- if countDist < k, keep increasing the window (increasing the right)
- if countDist > k, reduce the window, by increasnig the left
- if countDist == k, then we can start to increment the total good subarray. Observation:
  a. [left...right] is the maximum subarray that start with nums[left] and ends with nums[right] element. 
  b. how many good subarray is there from left...right subarray, that ends with nums[right] position?
      - Let the mostLeftIdx = min(lastIdx(A1), lastIdx(A2), ..., lastIdx(Ak)), where A1 to Ak denotes the k distinct numbers in the subarray.
      - then total subarray is = mostLeftIdx - left + 1.
      - For example, given subarray: [1,2,1,2] and k = 2.
        - left = 0, right = 3
        - mostLeftIdx = min(lastIdx(1), lastIdx(2)) = min(2, 3) = 2
        - then total subarray is = 2 - 0 + 1 = 3. That is {[1,2,1,2], [2,1,2], [1,2]}

    Question is: how to compute (and update) mostLeftIdx efficiently in O(1)?
    - We can use 2 data structure to help on this: 
      - double linked list to track the last index of distinct number
      - hashmap to track the pointer of such distinct number inside the double linked list, so we can easily delete that element in O(1), without iterating the double linked list 

*/
func subarraysWithKDistinct(nums []int, k int) int {

    arrFills(counter, 0) // to track the number occurences
    countDist := 0 // to track how many distinct number
    lastIdxWindow := NewLastIdxWindow() // to track the last index of certain number

    var left, total int
    n := len(nums)
    
    for right := 0;right < n;right++{

        counter[nums[right]]++
        if counter[nums[right]] == 1 {
            countDist++ // increment the countDistinct
        }
        lastIdxWindow.UpdateLastIdx(nums[right], right)

        // reduce the window by increasing the left pointer
        for left <= right && countDist > k {
            counter[nums[left]]--
            // if the nums[left] occurences is 0, then we need to update the countDist and remove the lastIdx of such element
            if counter[nums[left]] == 0 {
                countDist--
                lastIdxWindow.RemoveNum(nums[left])
            }
            left++
        }

        if countDist == k {
            total += lastIdxWindow.GetMostLeftIndex() - left + 1
        }

    }

    return total 
}

func arrFills(arr []int, val int) {
    for i:=0;i<len(arr);i++{
        arr[i] = val
    }
}

type LastIdxWindow struct {
    window *list.List
    lastIdxMap map[int]*list.Element
}

func NewLastIdxWindow() *LastIdxWindow {
    return &LastIdxWindow {
        window: list.New(),
        lastIdxMap: make(map[int]*list.Element),
    }
}

func (w *LastIdxWindow) UpdateLastIdx(num int, idx int) {
    element, found := w.lastIdxMap[num]

    if found {
        w.window.Remove(element)
    }
    newElement := w.window.PushBack(idx)
    w.lastIdxMap[num] = newElement
}

func (w *LastIdxWindow) GetMostLeftIndex() int {
    return w.window.Front().Value.(int)
}

func (w *LastIdxWindow) RemoveNum(num int) {
    element, found := w.lastIdxMap[num]

    if found {
        w.window.Remove(element)
        delete(w.lastIdxMap, num)
    }
}

func init() { debug.SetGCPercent(-1) }