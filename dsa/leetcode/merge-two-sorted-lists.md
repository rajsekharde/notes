# Merge two sorted lists

Given: Heads of two sorted lists list1, list2. Merge them into one sorted list and return its head.

## Approach 1

Iterate through the two lists and store elements in a min-heap. Pop elements one by one and store in
another list.

Time complexity:
m, n: no. of elements in the two lists
Iteration: O(m + n)
Storing in min-heap: O((m + n)log(m + n))
Popping top element n times: O((m + n)log(m + n))

Overall: O((m + n)log(m + n))

```bash
ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        priority_queue<int, vector<int>, greater<int>> pq;
        ListNode *temp;
        temp = list1;
        while(temp != nullptr) {
            pq.push(temp->val);
            temp = temp->next;
        }
        temp = list2;
        while(temp != nullptr) {
            pq.push(temp->val);
            temp = temp->next;
        }
        ListNode *node = new ListNode();
        ListNode *temp2 = node;
        while(!pq.empty()) {
            int val = pq.top();
            pq.pop();
            temp2->next = new ListNode(val);
            temp2 = temp2->next;
        }
        
        return node->next;
    }
```