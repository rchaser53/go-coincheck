package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	host := "https://coincheck.com/api"
	path := "/rate/xrp_jpy"
	url := host + path

	ts := strconv.FormatInt(getUnixMilli(), 10)
	text := ts + "GET" + "/exchange/orders/transactions"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("ACCESS-SIGNATURE", createSign(text))
	req.Header.Set("ACCESS-NONCE", ts)
	req.Header.Set("ACCESS-KEY", os.Getenv("CCKey"))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(byteArray)
	println(bodyStr)
}

func getUnixMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func createSign(text string) string {
	hash := hmac.New(sha256.New, []byte(os.Getenv("CCKeySecret")))
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
