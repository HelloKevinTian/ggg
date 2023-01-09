package some

import (
	"fmt"
	"ggg/examples/helper"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

func SendRambler() {
	// testGet()
	testPost()
}

func testGet() {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")
	fmt.Println(resp, err, string(resp.Body()))
}

type ErrorData struct {
	State   string `json:"state"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RespData struct {
	ErrorData
	Data []Account `json:"data"`
}

type Account struct {
	GameId      interface{} `json:"game_id"`
	AppId       int         `json:"app_id"`
	AccountId   string      `json:"account_id"`
	AccountName string      `json:"account_name"`
	CompanyId   int         `json:"company_id"`
	AppName     string      `json:"app_name"`
}

type ReqData struct {
	AuthToken   string   `json:"auth_token"`
	UserId      int      `json:"user_id"`
	CompanyId   int      `json:"company_id"`
	ChildSystem string   `json:"child_system"`
	AccountId   []string `json:"account_id"`
}

func testPost() {
	authToken := helper.MD5("fotoable" + "/fb/listAccountByUser" + "fotoable2018")
	reqData := ReqData{
		AuthToken:   authToken,
		UserId:      100018,
		CompanyId:   100002,
		ChildSystem: "bi",
		AccountId:   []string{"320985608725280"},
	}

	client := resty.New()
	resp, _ := client.SetDebug(true).R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqData).
		Post("https://ramblertest.nuclearport.com/fb/listAccountByUser")
	// fmt.Println(resp.Body())

	var errRes RespData

	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := jsonIterator.Unmarshal(resp.Body(), &errRes); err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(errRes.State, errRes.Code, errRes.Data)
}
