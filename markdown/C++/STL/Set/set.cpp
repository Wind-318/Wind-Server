#include <iostream>
#include <string>
#include <unordered_set>
using namespace std;

class MyData {
private:
    friend class hashfunc;
    int val;
    string name;
public:
    MyData(const int& _val, const string& _name) : val(_val), name(_name) {}

    const string& getName() const {
        return this->name;
    }

    const int& getVal() const {
        return this->val;
    }

    void setName(const string& s) {
        this->name = s;
    }

    void setVal(const int& v) {
        this->val = v;
    }

    // 实现比较
    bool operator==(const MyData& d) const {
        return this->val == d.val && this->name == d.name;
    }

    bool operator<(const MyData& a) const {
        return this->getVal() > a.getVal();
    }

};

// 实现哈希
class hashfunc {
public:
    size_t operator()(const MyData& d) const {
        return hash<string>()(d.name) ^ hash<int>()(d.val);
    }
};

class com {
public:
    bool operator()(const MyData& a, const MyData& b) const {
        return a.getName() == b.getName() && a.getVal() == b.getVal();
    }
};


int main() {
    unordered_set<MyData, hashfunc> s;
    s.emplace(MyData(15, "张三"));
    for (const auto& val : s) {
        cout << val.getName() << "   " << val.getVal() << endl;
    }

    set<MyData> s;
    s.emplace(MyData(22, "老七"));
    s.emplace(MyData(22, "老七"));
    s.emplace(MyData(25, "老八"));
    for (const auto& val : s) {
        cout << val.getName() << " " << val.getVal() << endl;
    }

    return 0;
}