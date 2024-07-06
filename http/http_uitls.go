package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	myConst "icode.baidu.com/baidu/personal-code/testGolang/const"
	"net/http"
	"time"
)

func SendRequest(url string, data *Request, traceID string) (*http.Response, error) {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)

	req, _ := http.NewRequest("POST", url, payloadBuf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("traceID", traceID)
	req.Header.Set("Authorization", myConst.Authorization)
	fmt.Printf("req : %+v\n", req)
	client := &http.Client{Timeout: time.Second * 10}

	return client.Do(req)
}
