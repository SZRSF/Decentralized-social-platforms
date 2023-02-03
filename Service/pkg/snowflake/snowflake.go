package snowflake

//雪花算法生成ID

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

// 全局node 节点
var node *sf.Node

// Init startTime 时间因子, machineID 机器标识
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	sf.Epoch = st.UnixNano() / 1e6
	node, err = sf.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// GenID 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}
