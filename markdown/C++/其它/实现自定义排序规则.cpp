//示例
#include <iostream>
#include <vector>
#include <functional>
using namespace std;

class Sort {
public:
	template<class T>
	void quickSort(vector<T>& arr, int left, int right, function<bool(T, T)> const& f = [](const auto& a, const auto& b)->bool {return a <= b; });
};

class myCom {
public:
	bool operator()(const int& a, const int& b) {
		return a >= b;
	}
};

int main() {
	Sort s;
	vector<int> a{3, 1, 5, 6, 0, 3, 1, -35, -23, 132, 61, -320};

    // 通过自定义仿函数定义排序规则
	s.quickSort<int>(a, 0, a.size() - 1, myCom());

	for (const auto& i : a) {
		cout << i << ",";
	}

	cout << endl;

    // 通过匿名函数自定义排序规则
	s.quickSort<int>(a, 0, a.size() - 1, [](const auto& a, const auto& b)->bool {return a <= b; });

	for (const auto& i : a) {
		cout << i << ",";
	}
	return 0;
}

// 支持自定义排序的快排，默认排序规则从小到大排序
template<class T>
void Sort::quickSort(vector<T>& arr, int left, int right, function<bool(T, T)> const& f) {
	if (left >= right) {
		return;
	}
	int low = left;
	int high = right;
	auto key = arr[low];

	while (low < high) {
		while (low < high && f(key, arr[high])) {
			high--;
		}
		arr[low] = arr[high];
		while (low < high && f(arr[low], key)) {
			low++;
		}
		arr[high] = arr[low];
	}

	arr[low] = key;
	quickSort(arr, left, low - 1, f);
	quickSort(arr, high + 1, right, f);
}