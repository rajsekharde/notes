# Climbing stairs

It takes n steps to reach the top of a staircase. Each time, you can climb 1 or 2 steps.
In how many distinct ways can the stairs be climbed?

# Approach 1

Forward DP. Recursively build up total number of ways to climb stairs, while storing
calculated values. Starts from 0.

Time complexity: O(n)

```bash
    int climbStairs(int n) {
        vector<int> cache(n, -1);
        return climb(0, n, cache);
    }
    int climb(int i, int n, vector<int> &cache) {
        if(i == n) {
            return 1;
        }
        if(i > n) return 0;

        if(cache[i] != -1) return cache[i];

        int a = climb(i+1, n, cache) + climb(i+2, n, cache);
        cache[i] = a;
        return a;
    }
```

# Approach 2

Similar to 1, but starts from n. (Backward DP)

```bash
    int climbStairs(int n) {
        vector<int> cache(n+1, -1);
        return climb(n, cache);
    }
    int climb(int n, vector<int> &cache) {
        if(n == 0) {
            return 1;
        }
        if(n < 0) return 0;

        if(cache[n] != -1) return cache[n];

        int a = climb(n-1, cache) + climb(n-2, cache);
        cache[n] = a;
        return a;
    }
```

# Approach 3

Iterative DP. Build up the solution from bottom. Possible solutions for 0, 1 is 1 (1 step). And
solution for any n is sol(n-1) + sol(n-2). 

Time complexity: O(n)

Space complexity: O(n)

```bash
    int climbStairs(int n) {
        vector<int> dp(n+1, 0);
        dp[0] = 1;
        dp[1] = 1;

        for(int i=2; i<=n; i++) {
            dp[i] = dp[i-1] + dp[i-2];
        }

        return dp[n];
    }

Example: n = 5

dp[0] = 1
dp[1] = 1

dp[2] = dp[0] + dp[1] = 1 + 1 = 2

dp[3] = dp[2] + dp[1] = 2 + 1 = 3

dp[4] = dp[3] + dp[2] = 3 + 2 = 5

dp[5] = dp[4] + dp[3] = 5 + 3 = 8
```
