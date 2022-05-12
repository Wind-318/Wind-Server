package data

// Algorithm 服务名
var AlgorithmServiceName = "AlgorithmService"

// 正态分布数据
type AlgorithmNormalDistributeInterface struct {
	// 偏移量
	Offset int
	// 最小值
	Minimum int
	// 最大值
	Maximum int
}

// randomtime 数据
type AlgorithmRandomTimeInterface struct {
	// 最少时间
	Start int
	// 随机再等待时间
	RangeNumber int
}

// zip 数据
type AlgorithmZipDataInterface struct {
	// 源文件夹地址
	Source string
	// 目标文件夹地址
	Target string
}

// match 数据
type AlgorithmMatchInterface struct {
	// 匹配字符串
	Str string
	// 待查询字符串
	Target string
}

// url 页面分析数据
type AlgorithmUrlHtmlInterface struct {
	// 页面 url 链接
	Url string
	// 正则匹配规则
	RegexpRule string
}

// string 分析数据
type AlgorithmHtmlInterface struct {
	// 要分析的字符串
	Html string
	// 正则匹配规则
	RegexpRule string
}
