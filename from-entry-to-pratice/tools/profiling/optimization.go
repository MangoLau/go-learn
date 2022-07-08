package profiling

import (
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}
	req := Request{"demo_transaction", payload}
	v, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func processRequest(reqs []string) []string {
	reps := []string{}
	
	for _, req := range reqs {
		reqObj := &Request{}
		json.Unmarshal([]byte(req), reqObj)		
		var ret strings.Builder
		for _, e := range reqObj.PayLoad {
			ret.WriteString(strconv.Itoa(e))
			ret.WriteRune(',')
		}
		repObj := &Response{reqObj.TransactionID, ret.String()}
		repJson, err := json.Marshal(&repObj)
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(repJson))
	}
	return reps
}

func processRequestOrigin(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		json.Unmarshal([]byte(req), reqObj)
		ret := ""
		for _, e := range reqObj.PayLoad {
			ret += strconv.Itoa(e) + ","
		}
		repObj := &Response{reqObj.TransactionID, ret}
		repJson, err := json.Marshal(&repObj)
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(repJson))
	}
	return reps
}
