package config
// 状态组
type stateGroup struct {
	Key  string  `json:"key"`
	Desc string  `json:"desc"`
	List []state `json:"list"`
}
type state struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
	Desc  string `json:"desc"`
}

// 所有状态配置
func GetAllStates() []stateGroup {
	return allStates
}
// allStates
// 所有用到的状态值，用于返回给前端, 前端只需要记key，不需要记value, 这样后端状态的变化不会导致前端调整很多，一旦有状态变化再这里重新配置
var allStates = []stateGroup{
	//---------user_status--start--
	{
		Key:  "user_status",
		Desc: "用户状态",
		List: []state{
			{
				Key:   "normal",
				Value: 1,
				Desc:  "正常",
			},
			{
				Key:   "forbidden",
				Value: 2,
				Desc:  "禁止登陆",
			},
			// more
		},
	},

	//--------user_status--end-----------
	{
		Key:  "order_status",
		Desc: "订单状态",
		List: []state{
			{
				Key:   "unpaid",
				Desc:  "未支付",
				Value: 1,
			},
			{
				Key:   "paid",
				Desc:  "已支付",
				Value: 2,
			},
			{
				Key:   "delivering",
				Desc:  "发货中",
				Value: 3,
			},
		},
	},
}
