package service

import (
	"encoding/json"
)

func Response(msg interface{}) []byte {
	resp := map[string]interface{}{
		"data":  msg,
		"error": "",
	}

	respRaw, _ := json.Marshal(resp)

	return respRaw
}

func ErrResponse(msg interface{}, err error) []byte {
	resp := map[string]interface{}{
		"data":  msg,
		"error": "",
	}

	if err != nil {
		resp["error"] = err.Error()
	}

	respRaw, _ := json.Marshal(resp)

	return respRaw
}
