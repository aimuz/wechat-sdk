package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func NewRequest(method, url string, data []byte) (body []byte, err error) {

	if method == "GET" {
		url = fmt.Sprint(url, "?", string(data))
		data = nil
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return body, err
	}

	resp, err := client.Do(req)

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	return body, err
}

// Struct2Map
func Struct2Map(r interface{}) (s map[string]string, err error) {
	var temp map[string]interface{}
	var result = make(map[string]string)

	bin, err := json.Marshal(r)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(bin, &temp); err != nil {
		return nil, err
	}
	for k, v := range temp {
		switch v2 := v.(type) {
		case string:
			//fmt.Printf("%s=%s\n", k, v2)
			result[k] = v2
			break
		case int8, uint8, int, uint, int32, uint32, int64, uint64:
			fmt.Println("k2=", v2)
			result[k] = fmt.Sprint(v2)
			break
		case float32, float64:
			result[k] = fmt.Sprint(v2)
			break
		}
	}
	//fmt.Println(result)
	return result, nil
}

// 生成签名
func GenWeChatPaySign(m map[string]string, payKey string) (string, error) {
	delete(m, "sign")
	var signData []string
	for k, v := range m {
		if v != "" {
			signData = append(signData, fmt.Sprintf("%s=%s", k, v))
		}
	}

	sort.Strings(signData)
	signStr := strings.Join(signData, "&")
	signStr = signStr + "&key=" + payKey

	c := md5.New()
	_, err := c.Write([]byte(signStr))
	if err != nil {
		return "", err
	}
	signByte := c.Sum(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", signByte), nil
}

// 生成订单号，不推荐直接使用

func GetTradeNO(prefix string) string {
	now := time.Now()
	strTime := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(),
		now.Second())
	return prefix + strTime + RandomNumString(100000, 999999)
}

func RandomNum(min int64, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := min + r.Int63n(max-min+1)
	return num
}

func RandomNumString(min int64, max int64) string {
	num := RandomNum(min, max)
	return strconv.FormatInt(num, 10)
}
