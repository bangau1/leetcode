class Solution {

    static class Node {
        int Value;
        int Index;

        public Node(int value, int index) {
            this.Value = value;
            this.Index = index;
        }
    }

    boolean isNodeInValidWindow(Node node, int minIndex, int maxIndex) {
        return node.Index >= minIndex && node.Index <= maxIndex;
    }

    public int[] maxSlidingWindow(int[] nums, int k) {
        LinkedList<Node> list = new LinkedList<Node>();
    
        if (k == 1) {
            return nums;
        }
        var result = new int[nums.length - k + 1];
        var resultInd = 0;
        // process first k window element
        for (int i = 0; i < k; i++) {
            Node node = new Node(nums[i], i);

            if (i == 0) {
                list.add(node);
                continue;
            } 

            // remove element until the currentNode is smaller than the last element
            while (list.size() > 0 && node.Value >= list.getLast().Value ) {
                list.removeLast();
            }
            list.add(node);
        }
        result[resultInd] = list.getFirst().Value;
        resultInd += 1;

        // then proceed sliding window 
        for (int i = k; i < nums.length; i++){
            var first = list.getFirst();

            if (!isNodeInValidWindow(first, i-k+1 , i) ){
                list.removeFirst();
            } 
            
            Node node = new Node(nums[i], i);
             // remove element until the currentNode is smaller than the last element
            while (list.size() > 0 && node.Value >= list.getLast().Value ) {
                list.removeLast();
            }
            list.add(node);


            result[resultInd] = list.getFirst().Value;
            resultInd += 1;
        }
        return result;
    }
}