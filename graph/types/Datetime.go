/*
@Time : 2020/10/14 15:17
@Author : Administrator
@Description :
@File : Datetime
@Software: GoLand
*/
package types

import (
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/json-iterator/go"
	"io"
	"strconv"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func appendQuote(v []byte) []byte {
	r := []byte("\"")
	r = append(r, v...)
	r = append(r, '"')
	return r
}

// Datetime 匹配 `Date` 的类型
type Datetime struct {
	t time.Time
}

const TimeLayout = "2006-01-02 15:04:05.000"

func NewDatetimeFromTime(t time.Time) *Datetime {
	return &Datetime{
		t: t,
	}
}

// GetTime 返回 time.Time
func (d *Datetime) GetTime() time.Time {
	return d.t
}

// UnmarshalGQL 反序列化方法，当 gqlgen 接受到请求后，会调用该方法，将该类型的数据进行反序列化
func (d *Datetime) UnmarshalGQL(vi interface{}) (err error) {
	v, ok := vi.(string)
	if !ok {
		return fmt.Errorf("unknown type of Datetime: `%+v`", vi)
	}
	if d.t, err = time.Parse(TimeLayout, v); err != nil {
		return err
	}

	return nil
}

// MarshalGQL 序列化
func (d Datetime) MarshalGQL(w io.Writer) {
	w.Write(appendQuote([]byte(d.t.Format(TimeLayout))))
}

//如果.gqlgen.yml引用的类型是返回编组器的函数，我们可以使用它来编码和解码到任意的go类型
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(TimeLayout)))
		//io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

//Unmarshal{Typename} 仅当标量作为输入的时候才需要，原始值已经从json 被解码成 int/float64/bool/nil/map[string]interface/[]interface
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}
