func maximalRectangle(matrix [][]byte) int {
    h := make([]int, len(matrix[0]))
    var maxArea int
    for r := 0;r<len(matrix);r++{
        for c:=0;c<len(h);c++{
            if matrix[r][c] == '1' {
                h[c] = h[c] + 1
            }else{
                h[c] = 0
            }
        }
        // fmt.Println(h)

        // compute the largest area of the histogram h
        var stack []int
        
        for c:=0;c<len(h);c++{
            for len(stack) > 0 && h[c] < h[stack[len(stack)-1]] {
                currH := h[stack[len(stack)-1]]
                stack = stack[:len(stack)-1]
                currW := c
                if len(stack) > 0 {
                    currW = c - (stack[len(stack)-1] + 1)
                }
                
                currArea := currH * currW
                if maxArea < currArea {
                    maxArea = currArea
                }
            }
            stack = append(stack, c)
        }
        for len(stack) > 0  {
            currH := h[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            currW := len(h)
            if len(stack) > 0 {
                currW = len(h) - (stack[len(stack)-1] + 1)
            }
            
            currArea := currH * currW
            if maxArea < currArea {
                maxArea = currArea
            }
        }

    }

    return maxArea
}