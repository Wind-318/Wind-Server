class LFUCache {
private:
	struct node {
		int key;
		int fre;
		int val;
		node(int _key, int _fre, int _val) : key(_key), fre(_fre), val(_val) {}
	};
	int minfre;
	int cap;
	unordered_map<int, list<node>::iterator> keys;
	unordered_map<int, list<node>> freq;
public:
	LFUCache(int capacity) {
		this->cap = capacity;
		this->minfre = 0;
	}

	int get(int key) {
		if (this->cap == 0 || !keys.count(key)) {
			return -1;
		}

		auto temp = keys[key];
		int tempVal = temp->val;
		int tempFre = temp->fre;

		freq[tempFre].erase(temp);

		if (freq[tempFre].size() == 0) {
			freq.erase(tempFre);
			if (this->minfre == tempFre) {
				this->minfre++;
			}
		}

		freq[tempFre + 1].emplace_front(node(key, tempFre + 1, tempVal));
		keys[key] = freq[tempFre + 1].begin();
		return tempVal;
	}

	void put(int key, int value) {
		if (this->cap == 0) {
			return;
		}
		if (!keys.count(key)) {
			if (keys.size() == cap) {
				auto its = freq[minfre].back();
				keys.erase(its.key);
				freq[minfre].pop_back();
				if (freq[minfre].size() == 0) {
					freq.erase(minfre);
				}
			}
			freq[1].emplace_front(node(key, 1, value));
			keys[key] = freq[1].begin();
			minfre = 1;
		}
		else {
			auto temp = keys[key];
			int tempFre = temp->fre;

			freq[tempFre].erase(temp);
			if (freq[tempFre].size() == 0) {
				freq.erase(tempFre);
				if (minfre == tempFre) {
					minfre += 1;
				}
			}
			freq[tempFre + 1].emplace_front(node(key, tempFre + 1, value));
			keys[key] = freq[tempFre + 1].begin();
		}
	}
};