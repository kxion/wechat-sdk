package wechat_sdk

import (
	"errors"
	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	订单查询
*/

func (m *myPayer) UnifiedQuery(param Param) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		signType           = e.SignTypeMD5 //此处默认MD5
		queryMustParam     = []string{"appid", "mch_id", "nonce_str", "sign"}
		queryOneParam      = []string{"transaction_id", "out_trade_no"}
		queryOptionalParam = []string{"sign_type"}
	)
	//校验订单号
	var count = 0
	for _, k := range queryOneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: transaction_id or out_trade_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: transaction_id or out_trade_no")
	}

	//校验其他参数
	for _, k := range queryMustParam {
		if k == "sign" {
			continue
		}
		if param.Get(k) == nil {
			return nil, errors.New("need param: " + k)
		}
	}

	for k := range param {
		if k == "sign" {
			continue
		}
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !util.HaveInArray(queryMustParam, k) && !util.HaveInArray(queryOptionalParam, k) && !util.HaveInArray(queryOneParam, k) {
			return nil, errors.New("no need param: " + k)
		}
	}
	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/orderquery",
		ContentType: e.PostContentType,
	}

	response, err := postToWx(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}

	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, e.ErrCheckSign
	}

	return result, nil
}
