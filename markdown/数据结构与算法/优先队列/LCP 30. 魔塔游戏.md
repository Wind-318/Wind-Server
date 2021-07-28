## [Leetcode.LCP 30. 魔塔游戏](https://leetcode-cn.com/problems/p0NxJO/)

```C++
class Solution {
public:
    int magicTower(vector<int>& nums) {
        if (accumulate(nums.begin(), nums.end(), 1ll) <= 0) {
            return -1;
        }
        int n = nums.size();
        int ans = 0;
        long long cur = 1;
        priority_queue<int, vector<int>, greater<int>> q;

        for (int i = 0; i < n; i++) {
            cur += nums[i];
            if (nums[i] < 0) {
                q.emplace(nums[i]);
            }
            if (cur <= 0) {
                cur -= q.top();
                q.pop();
                ans++;
            }
        }
        return ans;
    }
};
```