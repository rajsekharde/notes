# Top k frequent elements

Given: An integer array nums and integer k, return k most freuquent elements within the array.

## Approach 1

Store the frequencies of elements in a hash-map

Sort the map by frequencies and get the top element one by one, till we have k elements

-> Store pair{freq, num} in a priority queue (max-heap). The entries are automatically sorted by the first element, with the highest value at the top. Get the top element and pop the pq k times, and store the second value of entries in a vector.

Time complexity:
- hashing: O(n)
- storing in max-heap: O(mlogm)
- getting top k elements: O(klogm)
- Overall: O(mlogm) ~ O(nlogn)

```bash
vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> freq;
        for(int i: nums) freq[i]++;
        priority_queue<pair<int, int>> pq;
        for(auto it: freq) pq.push({it.second, it.first});
        vector<int> ans;
        for(int i=0; i<k; i++) {
            ans.push_back(pq.top().second);
            pq.pop();
        }
        return ans;
    }
```


## Approach 2

Instead of storing all elements in a max-heap, store only top k elements in a min-heap. keep storing elements of hash-map one at a time and pop the top (smallest) element, if size of pq > k. At the end, only the top k elements remain. Then pop all the elements and store in an array.

Time Complexity:
- hashing: O(n)
- storing in min-heap: O(mlogk)
- popping elements: O(klogk)
- overall: O(mlogk) ~ O(nlogk)

```bash
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> freq;
        for(int i: nums) freq[i]++;
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> pq;
        for(auto it: freq) {
            pq.push({it.second, it.first});
            if(pq.size() > k) pq.pop();
        }
        vector<int> ans;
        for(int i=0; i<k; i++) {
            ans.push_back(pq.top().second);
            pq.pop();
        }
        return ans;
```