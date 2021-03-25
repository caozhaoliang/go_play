package demoArchive

type RspStrategyDetail struct {
	Strategy     StrategyInfo   `json:"strategy"`
	Subscribe    SubscribeInfo  `json:"subscribe"`
	Auth         AuthInfo       `json:"auth"`
	SuccessCases []SuccessStock `json:"success_cases"`
	Stocks       []CurHoldInfo  `json:"stocks"` // 当前持仓
}
type StrategyInfo struct {
	Id             int        `json:"strategy_id"`
	Version        int        `json:"strategy_version"`
	Type           int        `json:"strategy_type"`
	Name           string     `json:"strategy_name"`
	Tags           []string   `json:"tags"`
	Instruction    string     `json:"instruction"`
	InstructionUrl string     `json:"instruction_url"`
	Senior         int        `json:"senior"`
	AllowFollow    int        `json:"allow_follow"`
	ViewPermission int        `json:"view_permission"`
	Returns        ReturnInfo `json:"returns"`
}

type ReturnInfo struct {
	AccuReturnRate  int64 `json:"accureturn_rate"`   //累计收益率
	AverageRateYear int64 `json:"average_rate_year"` // 年化收益率
	ReturnRateN     int   `json:"return_rate_N"`     // N
	ReturnRateNday  int64 `json:"return_rate_Nday"`  // N 天收益率
}
type SubscribeInfo struct {
	SubscribeStatus int    `json:"subscribe_status"`
	SuccessTips     string `json:"success_tips"`
	CancelTips      string `json:"cancel_tips"`
}

type AuthInfo struct {
	StrategyGroup int    `json:"strategy_group"`
	IsAuth        int    `json:"is_auth"`
	Tips          string `json:"tips"`
	BuyAccount    int    `json:"buy_account"`
	Worth         int    `json:"worth"`
	Unit          int    `json:"unit"`
}
type SuccessStock struct {
	SecuCode       string `json:"secu_code"`
	SecuName       string `json:"secu_name"`
	SelDate        string `json:"sel_date"`
	AccureturnRate int64  `json:"accureturn_rate"`
}
type CurHoldInfo struct {
	SecuCode       string `json:"secu_code"`
	SecuName       string `json:"secu_name"`
	SelPrice       int64  `json:"sel_price"`
	CurPrice       int64  `json:"cur_price"`
	SelDate        string `json:"sel_date"`
	ReturnRate     int    `json:"return_rate"`
	ReturnRateNDay int64  `json:"return_rate_Nday"` // N天收益率 n为0时表示当天收益率
}
