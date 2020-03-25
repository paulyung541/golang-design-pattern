package bridge

/*
	桥接模式：
	现在有两个维度分别会有自己的子类，这两个维度独立发展
	这个模式就是在两个维度的『抽象层』进行组合，
	比如「维度1」使用「维度2」的功能，或依赖「维度2」的操作

	例1：
		维度1：画图形，可以画「圆」「方」「椭圆」等等
		维度2：用什么工具画，「水彩笔」「蜡笔」

	例2：
		matchmaker 的候选客服，
		维度1：根据不同的规则筛选候选客服名单
		维度2：候选客服名单里面再根据不同规则

	例3：
		导出数据到文件：
		维度1：来源有不同的格式：string, byte
		维度2：输出文件格式有不同：json, Number(Mac)

	下面代码用 例3 进行，需要实现具体代码
	思路：string -> list<map> -> json
		 string -> list<map> -> Number

		 byte -> list<map> -> json
		 byte -> list<map> -> Number
*/

type DataForFile []map[string]string

type Resulter interface {
	ToFileData() DataForFile
}

type StringAdapter struct {
	data string
}

func (*StringAdapter) ToFileData() DataForFile {
	var dataList DataForFile
	// 这里处理数据
	return dataList
}

type ByteAdapter struct {
	data []byte
}

func (*ByteAdapter) ToFileData() DataForFile {
	var dataList DataForFile
	// 这里处理数据
	return dataList
}

type FileMaker interface {
	MakeFile() error
}

type JsonFileMaker struct {
	FileFrom Resulter
}

func (m *JsonFileMaker) MakeFile() error {
	fileData := m.FileFrom.ToFileData()
	//生成文件
	if fileData != nil {}
	return nil
}

type NumberFileMaker struct {
	FileFrom Resulter
}

func (m *NumberFileMaker) MakeFile() error {
	fileData := m.FileFrom.ToFileData()
	//生成文件
	if fileData != nil {}
	return nil
}
