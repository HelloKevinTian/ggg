package helper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// GetJSON ...
func GetJSON(url string, vo interface{}) bool {
	b, _ := GetJSONStatusCode(url, vo)
	return b
}

// GetJSONStatusCode ...
func GetJSONStatusCode(url string, vo interface{}) (bool, int) {
	fmt.Printf("func GetJSON start.URL: [%s]", url)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	start := time.Now()
	resp, err := client.Get(url)
	fmt.Printf("fun GetJSON spend time(%.2f s). url: %s", time.Since(start).Seconds(), url)
	if err != nil {
		fmt.Println(fmt.Errorf("send get error. %+v", err))
		return false, 0
	}
	defer resp.Body.Close()
	if vo == nil {
		return true, 0
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("read data error. %+v", err))
		return false, resp.StatusCode
	}
	fmt.Println("read date.", string(body))
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := jsonIterator.Unmarshal(body, vo); err != nil {
		fmt.Println(fmt.Errorf("format object error. %+v", err))
		return false, resp.StatusCode
	}
	return true, resp.StatusCode
}

//PostJSON 发送POST请求传入接受都为JSON格式 vo为返回数据的struct的指针。
func PostJSON(url, data string, vo interface{}) bool {
	fmt.Printf("func PostJSON start.URL: [%s] DATA: [%s]", url, data)
	reader := bytes.NewReader([]byte(data))
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	start := time.Now()
	resp, err := client.Post(url, "application/json", reader)
	fmt.Printf("fun PostJSON spend time(%.2f s). url: %s", time.Since(start).Seconds(), url)
	if err != nil {
		fmt.Println(fmt.Errorf("send post error. %+v", err))
		return false
	}
	defer resp.Body.Close()
	if vo == nil {
		return true
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("read data error. %+v", err))
		return false
	}
	//utils.InfoLog("read date.", string(body))
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := jsonIterator.Unmarshal(body, vo); err != nil {
		fmt.Println(fmt.Errorf("format object error. %+v", err))
		return false
	}
	return true
}

// PostForm 发送POST请求传入接受都为JSON格式 vo为返回数据的struct的指针。
func PostForm(postURL string, data map[string][]string, vo interface{}) bool {
	fmt.Printf("func PostForm start.URL: [%s] DATA: [%s]", postURL, data)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.PostForm(postURL, data)
	start := time.Now()
	fmt.Printf("fun PostJSON spend time(%.2f s). url: %s", time.Since(start).Seconds(), postURL)
	if err != nil {
		fmt.Println(fmt.Errorf("send post error. %+v", err))
		return false
	}
	defer resp.Body.Close()
	if vo == nil {
		return true
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("read data error. %+v", err))
		return false
	}
	//utils.InfoLog("read date.", string(body))
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := jsonIterator.Unmarshal(body, vo); err != nil {
		fmt.Println(fmt.Errorf("format object error. %+v", err))
		return false
	}
	return true
}
