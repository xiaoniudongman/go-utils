package config

/*
有关ini配置文件的封装
*/
import (
	"github.com/xiaoniudongman/go-utils/tools/errs"
	"gopkg.in/ini.v1"
)

type IniParser struct {
	conf_reader *ini.File // config reader
}

func (this *IniParser) Load(c string) error {
	conf, err := ini.Load(c)
	if err != nil {
		this.conf_reader = nil
		return err
	}
	this.conf_reader = conf
	return nil
}

func (this *IniParser) GetBool(section string, key string) bool {
	if this.conf_reader == nil {
		return false
	}
	s := this.conf_reader.Section(section)
	if nil == s {
		return false
	}
	_bool, err := s.Key(key).Bool()
	errs.CheckFatalErr(err)
	return _bool
}

func (this *IniParser) GetString(section string, key string) string {
	if this.conf_reader == nil {
		return ""
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return ""
	}
	return s.Key(key).String()
}

func (this *IniParser) GetInt(section string, key string) int {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_int, err := s.Key(key).Int()
	errs.CheckFatalErr(err)
	return value_int
}

func (this *IniParser) GetInt32(section string, key string) int32 {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_int, err := s.Key(key).Int()
	errs.CheckFatalErr(err)
	return int32(value_int)
}

func (this *IniParser) GetUint32(section string, key string) uint32 {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_int, err := s.Key(key).Uint()
	errs.CheckFatalErr(err)
	return uint32(value_int)
}

func (this *IniParser) GetInt64(section string, key string) int64 {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_int, err := s.Key(key).Int64()
	errs.CheckFatalErr(err)
	return value_int
}

func (this *IniParser) GetUint64(section string, key string) uint64 {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_int, err := s.Key(key).Uint64()
	errs.CheckFatalErr(err)
	return value_int
}

func (this *IniParser) GetFloat32(section string, key string) float32 {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_float, err := s.Key(key).Float64()
	errs.CheckFatalErr(err)
	return float32(value_float)
}

func (this *IniParser) GetFloat64(section string, key string) float64 {
	if this.conf_reader == nil {
		return 0
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}
	value_float, err := s.Key(key).Float64()
	errs.CheckFatalErr(err)
	return value_float
}

func (this *IniParser) GetSectionMap(section string) map[string]string {
	var value_map = make(map[string]string)
	s := this.conf_reader.Section(section)
	for _, v := range s.KeyStrings() {
		val := s.Key(v).String()
		value_map[v] = val
	}
	return value_map
}
