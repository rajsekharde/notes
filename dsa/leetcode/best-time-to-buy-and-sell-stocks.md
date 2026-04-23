# Best time to buy and sell stocks

Given: An integer array prices with prices[i] = price of a stock on ith day. Return maximum profit that can be achieved by choosing a single day to buy the stock and another day in the future to sell it. Return 0 for no profit.

## Approach 1

Maintain two pointers: l(buying day) and r(selling day). with l=0, r=1. Keep shifting r to the right till array limit is reached. For each r, check if profit is max wrt l. if prices[r] < prices[l], set l=r.

Time complexity: O(n)

```bash
    int maxProfit(vector<int>& prices) {
        int l=0, r=1, mp=0, n=prices.size();
        while(r<n) {
            if(prices[r] > prices[l]) mp = max(mp, (prices[r]-prices[l]));
            else l = r;
            r++;
        }
        return mp;
```

Topics: Sliding window, greedy