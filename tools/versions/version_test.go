package versions

import (
	"fmt"
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/xiaoniudongman/go-utils/config"
)

const (
	Config_path = "../config/test.yaml"
)

func TestMysqlDbInfo_QueryStruct(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	dbinfo := Version{}
	dbinfo.GeVersionFromConf(&c, "Version")
	fmt.Println(dbinfo.GetAlgoVersion())
}
