package bloomfilter

import (
	"testing"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/xiaoniudongman/go-utils/tools/filter/bitmap"

	"github.com/xiaoniudongman/go-utils/tools/types/nums"

	cuckoo "github.com/seiflotfy/cuckoofilter"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/dbs/rediss"
	"github.com/xiaoniudongman/go-utils/tools/converter"
	"github.com/xiaoniudongman/go-utils/tools/maths"
)

const lenth = 300000

func TestCuckoo(t *testing.T) {
	lenth := 30000
	can := maths.GenNRandInt(lenth, 100)
	cf := cuckoo.NewFilter(1000)
	var tmp []int
	for j := 0; j < lenth; j++ {
		if cf.InsertUnique([]byte(converter.IntToStr(can[j]))) {
			tmp = append(tmp, can[j])
		}
	}

	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		So(nums.UniqueInt(can)[11], ShouldEqual, tmp[11])
	})
}

func BenchmarkUniqueInt(b *testing.B) {
	b.StopTimer()
	lenth1 := 500
	can := maths.GenNRandInt(lenth1, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		nums.UniqueInt(can)
	}
}

func BenchmarkCuckoo(b *testing.B) {
	b.StopTimer()
	can := maths.GenNRandInt(lenth, 100)
	cf := cuckoo.NewFilter(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var tmp []int
		for j := 0; j < lenth; j++ {
			if cf.InsertUnique([]byte(converter.IntToStr(can[j]))) {
				tmp = append(tmp, can[j])
			}
		}
	}
	cf.Reset()
}

func BenchmarkUniqueRedis(b *testing.B) {
	b.StopTimer()
	can := maths.GenNRandInt(lenth, 100)
	canstr := converter.IntsToStrs(can)
	c := config.ConfigEngine{}
	var err error
	err = c.Load("../../../config/test.yaml")
	errs.CheckCommonErr(err)
	redisItem := new(rediss.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	redisdb := new(rediss.RedisDbInfo)
	redisdb.GetRedisConnFromConf(&c, "Redis")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bitmap.UniqueByRedis(redisdb, redisItem, canstr)
	}
}
