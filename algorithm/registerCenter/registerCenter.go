package registerCenter

// Algorithm 服务名
var AlgorithmServiceName = "AlgorithmService"

// Algorithm 服务接口
type AlgorithmServiceInterface interface {
	// zip 算法
	ZipAlgorithm(request AlgorithmZipDataInterface, reply *interface{}) error
	// randomTime 算法
	RandomTimeAlgorithm(request AlgorithmRandomTimeInterface, reply *int) error
	// match 算法
	MatchAlgorithm(request AlgorithmMatchInterface, reply *bool) error
	// url 页面分析
	UrlRegexpAlgorithm(request AlgorithmUrlHtmlInterface, reply *[][]string) error
	// 字符串分析
	StringRegexpAlgorithm(request AlgorithmHtmlInterface, reply *[][]string) error
	// 判断密码合法性
	JudgePasswordAlgorithm(password string, reply *interface{}) error
	// 获取 url 页面内容
	GetBytesAlgorithm(url string, reply *[]byte) error
	// IP 访问次数加一
	AddIPAlgorithm(ip string, reply *interface{}) error
	// 判断是否频繁访问
	IfRestrictedAlgorithm(ip string, reply *bool) error
	// 获取随机 user-agent
	GetRandomUserAgentAlgorithm(request interface{}, reply *string) error
	// 加密
	EncryptionAlgorithm(password string, reply *string) error
	// 生成符合正态分布的整数，传入偏移量，允许的下限和上限
	GetNormalDistributionNumberAlgorithm(request AlgorithmNormalDistributeInterface, reply *int) error
	// 生成随机字符串
	GenerateRandStrAlgorithm(length int, reply *string) error
}

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
