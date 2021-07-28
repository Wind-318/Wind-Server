## vector创建和使用

- 创建容器，添加元素，遍历容器
	```C++
	#include <iostream>
	// 引入 vector 头文件
	#include <vector>
	using namespace std;

	int main() 
	{
		// 创建一个 int 类型的 vector
		vector<int> v;
		for (int i = 0; i < 100; i++) 
		{
			// 使用 emplace_back 方法将值添加到 vector 尾部
			v.emplace_back(i);
		}
		// 使用 for 循环遍历，auto 关键字自动推导类型
		for (const auto& val : v)
		{
			cout << val << endl;
		}

		return 0;
	}
	```
- vector 几种初始化方法
	```C++
		// 创建含有 100 个元素的 v1，默认值为 0
		vector<int> v1(100);
		// 创建含有 100 个元素的 v2，默认值为 2
		vector<int> v2(100, 2);
		// 将 v2 拷贝赋值给 v3
		vector<int> v3(v2);

		for (int i = 0; i < 100; i++)
		{
			cout << "v1 = " << v1[i] << ", " << "v2 = " << v2[i] << ", " << "v3 = " << v3[i] << endl;
		}
	```