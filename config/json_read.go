package config

import (
	"encoding/json"

	"github.com/xiaoniudongman/go-utils/tools/errs"
)

// 从json string中获得map
func Json_parse(jsonText string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonText), &result)
	errs.CheckCommonErr(err)
	return
}
