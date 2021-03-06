package e

import "errors"

var (
	ErrNilParam     = errors.New("check param")
	ErrNilAppID     = errors.New("check param: appid")
	ErrNilAppSecret = errors.New("check param: appSecret")
	ErrorNoOpenId   = errors.New("check param: openid")
	ErrorNoToken    = errors.New("check param: access_token")
	ErrorInitClient = errors.New("call NewClientWithParam first")

	//统一下单必须参数
	UnifiedOrderMustParam = []string{"appid", "mch_id", "nonce_str", "sign", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type"}
	//统一下单可选参数
	UnifiedOrderOptionalParam = []string{"device_info", "sign_type", "detail", "attach", "fee_type", "time_start", "time_expire", "goods_tag", "limit_pay", "receipt", "openid"}

	//订单查询必需参数
	QueryOrderMustParam = []string{"appid", "mch_id", "nonce_str", "sign"}
	//订单查询可选参数
	QueryOrderOptionalParam = []string{"sign_type"}

	////
	///
	///
	//
	/*dev Branch*/
	ErrSignType    = errors.New("unknown sign type")
	ErrParams      = errors.New("param is empty")
	ErrTradeType   = errors.New("need trade_type")
	ErrOpenId      = errors.New("JSAPI need openid")
	ErrNoSign      = errors.New("result not have sign")
	ErrCheckSign   = errors.New("check sign fail")
	ErrNoWatermark = errors.New("got no watermark")
	ErrAppId       = errors.New("different appid")
	ErrMchId       = errors.New("different mch_id")
)
