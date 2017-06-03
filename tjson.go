package tjson

import (
	"encoding/json"
	"strconv"
	"fmt"
	"log"
)

type Json struct {
	Data interface{}
}

func NewJsonFormStr(str string) *Json {
	rtjson := &Json{}
	json.Unmarshal([]byte(str), &rtjson.Data)
	return rtjson
}

func (j *Json) Get(keys ...interface{}) *Json {
	if j == nil || len(keys) == 0 {
		return j
	}
	
	var currp interface{} = j.Data
	for _, kv := range keys {
		switch key := kv.(type) {
		case int:
			cpv, isarr := currp.([]interface{})
			if !isarr {
				log.Printf(`unknow path key type %T %T`, key, currp)
				return nil
			}
			if key < 0 || key >= len(cpv) {
				log.Printf(`out of index idx:%d len:%d`, key, len(cpv))
				return nil
			}
			currp = cpv[key]
		case string:
			cpv, isobj := currp.(map[string]interface{})
			if !isobj {
				log.Printf(`unknow path key type %T %T`, key, currp)
				return nil
			}
			sigv, ishas := cpv[key]
			if !ishas {
				log.Printf(`unknow path key:%s`, key)
				return nil
			}
			currp = sigv
		default:
			log.Printf(`unknow key type %T`, kv)
			return nil
		}
	}
	return &Json{Data: currp}
}

func (j *Json) Int(defaultV ...int) int {
	defaultv := 0
	if len(defaultV) >= 1 {
		defaultv = defaultV[0]
	}
	if j==nil{
		return defaultv
	}
	switch vv := j.Data.(type) {
	case int:
		return vv
	case float64:
		return int(vv)
	case string:
		ri, e := strconv.ParseFloat(vv, 64)
		if e != nil {
			return defaultv
		} else {
			return int(ri)
		}
	default:
		log.Println(fmt.Errorf(`unknowType:%T %v`, j.Data,j.Data))
		return defaultv
	}
	return defaultv
}

func (j *Json) String() string {
	return j.StringWithDefault(``)
}

func (j *Json) StringWithDefault(defaultV ...string) string {
	defaultv := ``
	if len(defaultV) >= 1 {
		defaultv = defaultV[0]
	}
	if j==nil{
		return defaultv
	}
	switch vv := j.Data.(type) {
	case string:
		return vv
	case int:
		return strconv.Itoa(vv)
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case []interface{},map[string]interface{}:
		str, _ := json.Marshal(j.Data)
		return string(str)
	default:
		log.Println(fmt.Errorf(`unknowType:%T %v`, j.Data,j.Data))
		return defaultv
	}
	return defaultv
}