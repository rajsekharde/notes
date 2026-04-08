# Contains duplicate

Given: integer array nums. Return true if any element appears at least twice. return false if all elements are unique

## Approach 1

Maintain a hash set for elements already seen

Traverse the array. if current element belongs to set, return false. return true after traversing

```bash
unordered_map<int> seen;
for(int i: nums)
{
    if(seen.count(i)) return true
    else seen.insert(i);
}
return false;
```

Time complexity: Traversal - O(n), hashing/searching: O(1). Overall: O(n)

Space complexity: Hash set: O(n)

## Approach 2

Sort the array, and check if consecutive elements are equal. if duplicates are present, they should be conseccutive in the sorted array.

```bash
sort(nums.begin(), nums.end());
for(int i=0; i<nums.size()-1; i++)
{
    if(nums[i] == nums[i+1]) return true;
}
return false;
```

Time complexity: Sorting: O(nlogn), traversal: O(n). Overall: O(nlogn)

Space complexity: O(1)