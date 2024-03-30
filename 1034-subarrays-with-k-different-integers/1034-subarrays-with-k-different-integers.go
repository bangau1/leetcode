import (
    "container/list"
)
func subarraysWithKDistinct(nums []int, k int) int {
    // general idea:
    // - need to maintain a counter map
    // - if len(counter) < k, keep increasing the window
    // - if len(counter) > k, reduce the window, by increase the left part of the window
    // - if len(counter) == k, then 
    //   - left to right is the maximum length of such subarray on that window. To know how many subarray in total for a given window, do it in naive way?
    //   - try to increase the left (in temp variable) until the len(counter) < k
    //
    // further improve on how to avoid naive way when len(counter) == k
    // - we keep the lastIndex on certain data structure
    // - then get the first minimum value, then the subtotal+= minValue - left + 1
    counter := make(map[int]int)
    lastIdxWindow := NewLastIdxWindow()

    var left, total int
    n := len(nums)
    
    for r := 0;r < n;r++{
        counter[nums[r]]++
        lastIdxWindow.Append(nums[r], r)

        for left <= r && len(counter) > k {
            counter[nums[left]]--
            if counter[nums[left]] == 0 {
                delete(counter, nums[left])
                lastIdxWindow.RemoveNum(nums[left])
            }
            left++
        }

        if len(counter) == k {
            total += lastIdxWindow.GetLeft() - left + 1
        }

    }

    return total 
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

func (w *LastIdxWindow) Append(num int, idx int) {
    element, found := w.lastIdxMap[num]

    if found {
        w.window.Remove(element)
    }
    newElement := w.window.PushBack(idx)
    w.lastIdxMap[num] = newElement
}

func (w *LastIdxWindow) GetLeft() int {
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