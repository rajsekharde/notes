# Reverse Linked List

Given: head of a singly linked list. Reverse the list and return the reversed list.

## Approach 1

Traverse through the list, push values to an array in order. Reverse the array and keep replacing the values of the list with the values in the array.

Time Complexity: O(n)

Space complexity: O(n)

```bash
    ListNode* reverseList(ListNode* head) {
        vector<int> temp;
        ListNode *tmp = head;
        while(tmp != nullptr) {
            temp.push_back(tmp->val);
            tmp = tmp->next;
        }
        int i=0;
        tmp = head;
        reverse(temp.begin(), temp.end());
        while(tmp != nullptr) {
            tmp->val = temp[i];
            tmp = tmp->next;
            i++;
        }
        return head;
    }
```

# Approach 2

Using recursion. Each recursive function call returns a reversed linked list. In each call, reverse the direction of link between current node and next node. next node will be automatically at the end of reversed list after recursive function call. and return the new head.

Time complexity: O(n)

Space complexity: O(1)

```bash
    ListNode* reverseList(ListNode* head) {
        if(!head || !head->next) return head;
        ListNode *newHead = reverseList(head->next);
        head->next->next = head;
        head->next = nullptr;
        return newHead;
    }
```

# Approach 3

Using iteration. Iterate throught the list and keep reversing the direction of connections. For each node curr, {next = curr->next, curr->next = prev, prev = curr}, then move to the next node using curr = next.

Time complexity: O(n)

Space complexity: O(1)

```bash
    ListNode* reverseList(ListNode* head) {
        ListNode *prev = nullptr, *next = nullptr, *curr;
        curr = head;
        while(curr != nullptr) {
            next = curr->next;
            curr->next = prev;
            prev = curr;
            curr = next;
        }
        head = prev;
        return head;
    }
```