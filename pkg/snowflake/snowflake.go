package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineId int64) (err error) {
	// startTime 雪花算法开始计算的时间
	// machineId 机器唯一标识
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
