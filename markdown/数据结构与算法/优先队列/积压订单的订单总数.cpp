class Solution {
public:
    struct sell {
        int price;
        int number;
    };
    struct buy {
        int price;
        int number;
    };
    class com {
    public:
        bool operator()(const sell& a, const sell& b) {
            return a.price > b.price;
        }
        bool operator()(const buy& a, const buy& b) {
            return a.price < b.price;
        }
    };
    int getNumberOfBacklogOrders(vector<vector<int>>& orders) {
        priority_queue<sell, vector<sell>, com> sellQueue;
        priority_queue<buy, vector<buy>, com> buyQueue;
        for (const auto& order : orders) {
            if (order[2] == 1) {
                sell s = {order[0], order[1]};
                sellQueue.emplace(s);
            } else {
                buy b = {order[0], order[1]};
                buyQueue.emplace(b);
            }
            while (!sellQueue.empty() && !buyQueue.empty() && (buyQueue.top().price >= sellQueue.top().price)) {
                auto s = sellQueue.top();
                auto b = buyQueue.top();
                sellQueue.pop();
                buyQueue.pop();
                if (s.number > b.number) {
                    s.number -= b.number;
                    sellQueue.emplace(s);
                } else if (s.number < b.number) {
                    b.number -= s.number;
                    buyQueue.emplace(b);
                }
            }
        }
        int ans = 0;
        while (!sellQueue.empty()) {
            auto temp = sellQueue.top();
            sellQueue.pop();
            ans += temp.number % int(1e9 + 7);
            ans %= int(1e9 + 7);
        }
        while (!buyQueue.empty()) {
            auto temp = buyQueue.top();
            buyQueue.pop();
            ans += temp.number % int(1e9 + 7);
            ans %= int(1e9 + 7);
        }
        return ans;
    }
};