//智能指针实现
class LRUCache {
public:
	//结点定义
    struct node {
		//指向下一个结点
        shared_ptr<node> next;
		//指向前一个结点
        weak_ptr<node> pre;
		//记录key
        int key;
		//记录value
        int val;

        node() : key(0), val(0) {}
        node(const int& _key, const int& _val) : key(_key), val(_val) {}
    };

	//一张记录key到链表结点的哈希表，用以实现O(1)复杂度查找
    unordered_map<int, shared_ptr<node>> rec;
	//头指针
    shared_ptr<node> head;
	//尾指针
    shared_ptr<node> tail;
	//容量
    int cap;
	//构造函数
    LRUCache(int capacity) {
        this->cap = capacity;
        this->head = make_shared<node>();
        this->tail = make_shared<node>();
        this->head->next = this->tail;
        this->tail->pre = this->head;
    }

    int get(int key) {
		//如果没有，返回-1
        if (!rec.count(key)) {
            return -1;
        }
		//如果有，先把这个结点放回头部位置，再返回找到的值
        toHead(rec[key]);
        return rec[key]->val;
    }

    void put(int key, int value) {
		//如果已经有这个值了，更新value，然后放回头部位置
        if (rec.count(key)) {
            rec[key]->val = value;
            toHead(rec[key]);
        } else if (rec.size() == cap) {   
			//如果已经满容量，删除哈希表记录的最后一个结点
            rec.erase(tail->pre.lock()->key);
			//同时在链表中也删除这个结点
            removeNode(tail->pre.lock());
			//再添加新节点到头部
            addToHead(make_shared<node>(key, value));
			//同时用哈希表记录位置
            rec[key] = head->next;
        } else {
			//容量没满，也没有这个key，直接添加，然后哈希表记录位置
            addToHead(make_shared<node>(key, value));
            rec[key] = head->next;
        }
    }

	//移动到头部的逻辑
    void toHead(shared_ptr<node>& ptr) {
		//让此结点的下一个结点指向前一个结点的指针指向此节点的前一个结点
        ptr->next->pre = ptr->pre;
		//让此节点的前一个结点的指向下一个结点的指针指向此节点的下一个结点
        ptr->pre.lock()->next = ptr->next;
		//让此节点的指向下一个结点的指针指向头节点的下一个指针
        ptr->next = head->next;
		//头节点的下一个结点的指向前一个结点的指针指向此节点
        head->next->pre = ptr;
		//头节点的指向下一个结点的指针指向此节点
        head->next = ptr;
		//此节点的指向前一个结点的指针指向头节点
        ptr->pre = head;
    }

	//添加到头部的逻辑
    void addToHead(shared_ptr<node>&& ptr) {
		//让此节点的指向下一个结点的指针指向头节点的下一个指针
        ptr->next = head->next;
		//头节点的下一个结点的指向前一个结点的指针指向此节点
        head->next->pre = ptr;
		//头节点的指向下一个结点的指针指向此节点
        head->next = ptr;
		//此节点的指向前一个结点的指针指向头节点
        ptr->pre = head;
    }

	//移除一个节点的逻辑
    void removeNode(shared_ptr<node>&& ptr) {
		//让此节点的前一个结点指向下一个结点的指针指向此节点的下一个结点
        ptr->pre.lock()->next = ptr->next;
		//让此节点的下一个节点的指向前一个结点的指针指向此节点的前一个结点
        ptr->next->pre = ptr->pre;
    }
};