## Set
- ### C++ 中有四种 set 类型，分别是：
  - **set** ：底层由红黑树实现，只存储键值，并且是**有序**存储，不允许有重复值存在，插入查找复杂度为$O(logn)$，需要包含头文件 **#include \<set>**
  - **multiset** ：底层由红黑树实现，只存储键值，并且是**有序**存储，允许有重复值存在，插入查找复杂度为$O(logn)$，需要包含头文件 **#include \<set>**
  - **unordered_set** ：底层由哈希表实现，只存储键值，并且是**无序**存储，不允许有重复值存在，插入查找复杂度为$O(1)$，需要包含头文件 **#include \<unordered_set>**
  - **unordered_multiset** ：底层由哈希表实现，只存储键值，并且是**无序**存储，允许有重复值存在，插入查找复杂度为$O(1)$，需要包含头文件 **#include \<unordered_set>**

- ### set Demo
	```C++
	#include <iostream>
	#include <ctime>
	#include <set>
	using namespace std;

	int main() {
		srand((size_t)time(nullptr));
		set<int> s;
		for (int i = 0; i < 100; i++) {
			s.emplace(rand());
		}

		for (const auto& val : s) {
			cout << val << " ";
		}

		return 0;
	}
	/*
	示例输出为：
	345 480 566 731 825 1970 2171 2225 2417 2426 2910 3109 3378 3817 4220 4452 
	4702 5011 7182 8039 8062 8344 8419 9224 9483 9487 9505 9790 11129 12084 
	12622 12875 13034 13350 13682 13779 13833 14463 14567 14682 14909 
	15561 15569 15747 16064 16522 16646 16862 16869 16904 17241 17336 17459 
	17592 17721 17736 18076 18481 18926 19511 19709 19747 20048 20115 20392 
	20997 21045 21370 21431 21676 22220 22492 22627 22958 23140 23303 23395 
	23773 24129 24294 25032 26045 26069 26096 26101 26414 26813 26911 27052 
	27491 27559 27584 27598 28588 29033 29513 29798 31194 31533 31809
	*/
	```
- ### unordered_set Demo
	```C++
	// 使用方法同 set，剩下两种 set 使用方法同上
	#include <iostream>
	#include <ctime>
	#include <unordered_set>
	using namespace std;

	int main() {
		srand((size_t)time(nullptr));
		unordered_set<int> s;
		for (int i = 0; i < 100; i++) {
			s.emplace(rand());
		}
		for (const auto& val : s) {
			cout << val << " ";
		}

		return 0;
	}
	```
- ### 方法
| 方法 | 作用 | 函数原型 |
| :--: | :--: | :--: |
| count(Keyval) | 查找 Keyval 个数 | size_t count(const int& _Keyval) const |
| contains(Keyval) | 查找 Keyval 是否存在 | size_t contains(const int& _Keyval) const |
| begin() | 返回指向容器中第一个元素的迭代器(还有一个 const 版本) | iterator begin() noexcept |
| end() | 返回指向容器中最后一个位置的**后一个位置**的迭代器(还有 const 版本) | iterator end() noexcept |
| rbegin() | 返回指向容器中第一个元素的**前一个位置**的**反向**迭代器(还有一个 const 版本) | reverse_iterator rbegin() noexcept |
| rend() | 返回指向容器中最后一个位置的**反向**迭代器(还有 const 版本) | reverse_iterator rend() noexcept |
| find(Keyval) | 查找 Keyval 值，若存在则返回该元素迭代器，没找到返回 end() 迭代器 | iterator find(const _Other& _Keyval) |
| lower_bound(Keyval) | 返回找到的第一个大于等于 Keyval 的元素的迭代器 | iterator lower_bound(const _Other& _Keyval) |
| upper_bound(Keyval) | 返回找到的第一个大于 Keyval 的元素的迭代器 | iterator upper_bound(const key_type& _Keyval) |
| empty() | 判断容器是否为空 | bool empty() const noexcept |
| size() | 判断元素个数 | size_t size() const noexcept |
| max_size() | 返回容器能容纳的最多元素个数 | size_t max_size() const noexcept |
| erase(Keyval) | 删除 Keyval 元素 | iterator erase(iterator _Where) noexcept |
| swap(set& right) | 交换这两个 set 中的元素 | void swap(set& _Right) noexcept(noexcept(_Mybase::swap(_Right))) |
| clear() | 清空容器中的元素 | void clear() noexcept |
| emplace(Keyval) | 直接构造添加 Keyval 元素到容器中，并返回一个迭代器和是否插入成功 | pair<iterator, bool> emplace(_Valtys&&... _Vals) |
- ### unordered_set 自定义键值
  - 首先看一下 unordered_set 的定义：
	```C++
	template <
	class _Kty, 				// 键值
	class _Hasher = hash<_Kty>, 		// 哈希函数
	class _Keyeq = equal_to<_Kty>, 		// 比较函数
	class _Alloc = allocator<_Kty>		// 分配器
	>
	```
	想要存储自定义键值，需要自己实现哈希函数和比较函数
  - #### 通过重载运算符实现
	```C++
	#include <iostream>
	#include <string>
	#include <unordered_set>
	using namespace std;

	// 自定义类型
	class MyData {
	private:
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

		// 实现比较，后面要加 const
		bool operator==(const MyData& d) const {
			return this->val == d.val && this->name == d.name;
		}

	};

	// 实现哈希，后面要加 const
	class hashfunc {
	public:
		size_t operator()(const MyData& d) const {
			return hash<string>()(d.getName()) ^ hash<int>()(d.getVal());
		}
	};


	int main() {
		unordered_set<MyData, hashfunc> s;
		s.emplace(MyData(15, "张三"));
		for (const auto& val : s) {
			cout << val.getName() << "   " << val.getVal() << endl;
		}

		return 0;
	}
	```
  - #### 通过自定义比较函数实现
    ```C++
	#include <iostream>
	#include <string>
	#include <unordered_set>
	using namespace std;

	// 自定义类型
	class MyData {
	private:
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
	};

	// 实现哈希
	class hashfunc {
	public:
		size_t operator()(const MyData& d) const {
			return hash<string>()(d.getName()) ^ hash<int>()(d.getVal());
		}
	};


	class com {
	public:
		bool operator()(const MyData& a, const MyData& b) const {
			return a.getName() == b.getName() && a.getVal() == b.getVal();
		}
	};

	int main() {
		unordered_set<MyData, hashfunc, com> s;
		s.emplace(MyData(15, "张三"));
		for (const auto& val : s) {
			cout << val.getName() << "   " << val.getVal() << endl;
		}

		return 0;
	}
    ```
- ### set 自定义键值
  - #### set 定义
    ```C++
	template <
	class _Kty, 
	class _Pr = less<_Kty>, 		
	class _Alloc = allocator<_Kty>
	>
	```
  - #### set 存储自定义键值可以通过重载小于运算符实现
    ```C++
	#include <iostream>
	#include <string>
	#include <set>
	using namespace std;

	// 自定义类型
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

		bool operator<(const MyData& a) const {
			return this->getVal() > a.getVal();
		}
	};

	int main() {
		set<MyData> s;
		s.emplace(MyData(22, "老七"));
		s.emplace(MyData(22, "老七"));
		s.emplace(MyData(25, "老八"));
		for (const auto& val : s) {
			cout << val.getName() << " " << val.getVal() << endl;
		}

		return 0;
	}
	```
***
# 参考
- #### [C++ STL set容器完全攻略](http://c.biancheng.net/view/7192.html)