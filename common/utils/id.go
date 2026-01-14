package utils

import (
	"strconv"

	"github.com/yitter/idgenerator-go/idgen"
)

func InitIdGenerator() {
	// 1. 创建配置选项，传入WorkerId
	options := idgen.NewIdGeneratorOptions(1) // WorkerId 需要确保唯一

	// 2. （可选）自定义各项参数
	// 以下是默认值，你可以按需调整：
	options.WorkerIdBitLength = 6    // WorkerId 占用的位数（默认6，支持最多2^6=64个实例）
	options.SeqBitLength = 6         // 序列号占用的位数（默认6，每毫秒最多生成2^6=64个ID）
	options.BaseTime = 1582136402000 // 基准时间（毫秒），默认2020-02-20 02:20:02，从这个时间开始算

	// 3. 应用配置并初始化
	idgen.SetIdGenerator(options)
}

func NextIdStr() string {
	return strconv.FormatInt(idgen.NextId(), 10)
}

func NextId() int64 {
	return idgen.NextId()
}
