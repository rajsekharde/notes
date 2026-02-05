# Hash-Map, Hash-Set working

Internally, both are built on top of a plain array

1. Key is hashed into an integer using a hash function

2. The integer is converted to an array index using modulo

```bash
hash(key) -> integer
index = hash(key) % array-size
```

Arrays give O(1) access time using indexes. So lookup for each key in
a hash-map or a value in a hash-set also takes O(1) time

On calling a search function- map.get("cat"):
- hash("cat") is computed
- hash is converted to an index
- the slot in array is accessed directly
- value is returned

Time complexity:
- Insert- O(1)
- Lookup- O(1)
- Delete- O(1)

## Collisions

Two different keys can produce the same index
- hash("cat") % 8 = 3
- hash("dog") % 8 = 3

Solution- Separate Chaining:
- each array slot holds a bucket, usually a linked list or small array
- index 3 -> [{"cat", 5}, {"dog", "9"}]
- Lookup: go to index 3 and search only the small bucket

Other solutions: Open addressing

C++ uses separate chaining:
- an unordered_map is basically a vector<bucket>
- bucket: Linked list of elements
