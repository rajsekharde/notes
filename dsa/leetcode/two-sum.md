# Two Sum

Given, an integer array nums and int target. Return indices of two numbers such that they add up to target.

## Brute force solution

Check every possible pair of numbers. Time complexity = O(n^2)

## Optimised solution

nums = [2,7,11,15], target = 9

store {nums: indices} in a hash map

iterate through nums. For each number, check if target-number is present in set. If yes, return indices

```bash
vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> mp;

    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];

        if (mp.count(complement)) {
            return {mp[complement], i};
        }

        mp[nums[i]] = i;
    }

    return {};
}
```

return {}
Time complexity = O(n)