#include <iostream>
//引入vector头文件
#include <vector>
using namespace std;

int main() 
{
	//创建一个int类型的vector
	vector<int> v;
	for (int i = 0; i < 100; i++) 
	{
		//使用emplace_back方法将值添加到vector尾部
		v.emplace_back(i);
	}
	//使用for循环遍历，auto关键字自动推导类型
	for (const auto& val : v)
	{
		cout << val << endl;
	}

    //创建含有100个元素的v1，默认值为0
    vector<int> v1(100);
    //创建含有100个元素的v2，默认值为2
	vector<int> v2(100, 2);
    //将v2拷贝赋值给v3
	vector<int> v3(v2);

	for (int i = 0; i < 100; i++)
	{
		cout << "v1 = " << v1[i] << ", " << "v2 = " << v2[i] << ", " << "v3 = " << v3[i] << endl;
	}

	return 0;
}