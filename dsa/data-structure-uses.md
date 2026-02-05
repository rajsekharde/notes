# When to use different data structures:

## Vector / Resizable Array
Time Complexity:
- Access by index- O(1)
- Access by value- O(n)

- Insertion at end- O(1)
- Insertion at position- O(n)

- Deletion from end- O(1)
- Deletion from position- O(n)

- Find(Linear search)- O(n)
- Binary search(If sorted)- O(log n)

- Sort- O(n log n)
- Traversal- O(n)

Used when:
- Maintaining order is important
- Frequent iterations (fast iteration is needed)
- Fast lookup not needed, just sequential access and stability
- Low memory overhead is important
- Most operations are push_back() or sequential additions
- Searching by index

## Hash Set / unordered_set
Time Complexity:
- Insertion- O(1) avg, O(n) worst case due to collisions
- Deletion- O(1) avg, O(n) worst case
- Find- O(1) avg, O(n) worst case
- Contains- O(1) avg, O(n) worst case
- Iteration- O(n)
- Rehashing(resizing)- O(n)

Used when:
- Order doesn't matter
- Fast lookups, insertion, deletion is needed
- No duplicates present
- Memory overhead is acceptable (Will use extra storage for hashes etc)
- Perfect when checking if an item is present quickly

## unordered_multiset
Hash-based set, allows duplicates

## Ordered Set / set
Time Complexity:
- Insertion- O(log n)
- Deletion- O(log n)
- Find- O(log n)
- Contains- O(log n)
- Iteration- O(n)
- Lower/Upper bound- O(log n)

Used when:
- Insertion and Deletion while mainting sorted order at all times
- No duplicates present
- Need to traverse in sorted order or perform range queries
- Ok with slow lookups and insertions
- Require automatic sorting after each modification

## multiset
Ordered set, allows duplicates

## Ordered Map / map
Time Complexity:
- Insertion- O(log n)
- Find- O(log n)
- Deletion- O(log n)
- Iteration- O(n), in sorted order

Used when:
- Sorted collection is needed where keys are always sorted
- Performing range queries(lower/upper bound)
- Guaranteed O(log n) time complexity needed for inserts, deletes, lookups

## multimap
Ordered map, allows duplicates

## Hash Map / unordered_map
Time Complexity:
- Insertion- O(1) avg, O(n) worst case due to collisions
- Find- O(1) avg, O(n) worst case
- Deletion- O(1) avg, O(n) worst case
- Iteration- O(n), order not guaranteed

Used when:
- Ordering is not required
- Fast lookups and insertions are needed
- Quickly checking if a key exists
- Frequent inserts, deletes

## unordered_multimap
Hash-based map, allows duplicates

## String (char vector)
Time Complexity:
- Index access- O(1)
- Concatenation- O(n)
- Substring- O(n)
- Comparison- O(n)

## Stack (LIFO)
Time Complexity:
- Push- O(1)
- Pop- O(1)
- Top- O(1)

## Queue / Deque
Time Complexity:
- Push- O(1)
- Pop- O(1)
- Front/Back- O(1)

Use Deque when-
- Frequent insertions/deletions from front and back

## Priority Queue / Heap
Time Complexity:
- Insertion- O(log n)
- Pop Max/Min- O(log n)
- Peek- O(1)

Used when:
- Frequentl take min/max value
- Duplicates present
- Top K elements
- Scheduling

## Binary Tree
Time Complexity:
- Traversal- O(n)
- Height- O(n) worst

## Binary Search Tree
Time Complexity:
- Insertion- O(log n) avg, O(n) worst
- Search- O(log n) avg, O(n) worsr

Used when:
- Sorted data
- Order-based queries

## Balanced BST
Time Complexity:
- Insertion-  O(log n)
- Search- O(log n)
- Deletion- O(log n)

Use when-
- Always-sorted data
- Range queries
- Order traversal

-> Underlying data structure for Ordered Set, Map

## Trie / Prefix Tree
Time Complexity:
- Insertion- O(L) # L - Word length
- Search- O(L)

Used when-
- Prefix search
- Autocomplete
- Dictionary problems
