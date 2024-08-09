package my_type

import "encoding/json"

type GoodsDataType int

const (
	Grab     = 1
	Simulate = 2
	Input    = 3
)

func (p GoodsDataType) String() string {
	var str string
	switch p {
	case Grab:
		str = "爬虫"
	case Simulate:
		str = "模拟"
	case Input:
		str = "输入"
	}
	return str
}

func (p GoodsDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
