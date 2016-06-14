package loggan_test

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/naoina/loggan"
)

func TestRawFormatter_Format(t *testing.T) {
	now := time.Now()
	for _, v := range []struct {
		entry  *loggan.Entry
		expect string
	}{
		{&loggan.Entry{
			Level:   loggan.DEBUG,
			Time:    now,
			Message: "test_raw_log1",
			Fields: loggan.Fields{
				"first":  1,
				"second": "2",
				"third":  "san",
			},
		}, "test_raw_log1"},
		{&loggan.Entry{
			Message: "test_raw_log2",
			Level:   loggan.INFO,
		}, "test_raw_log2"},
	} {
		var buf bytes.Buffer
		formatter := &loggan.RawFormatter{}
		if err := formatter.Format(&buf, v.entry); err != nil {
			t.Errorf(`RawFormatter.Format(&buf, %#v) => %#v; want %#v`, v.entry, err, nil)
		}
		actual := buf.String()
		expect := v.expect
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf(`RawFormatter.Format(&buf, %#v) => %#v; want %#v`, v.entry, actual, expect)
		}
	}
}

func TestLTSVFormatter_Format(t *testing.T) {
	now := time.Now()
	for _, v := range []struct {
		entry    *loggan.Entry
		expected string
	}{
		{&loggan.Entry{
			Level:   loggan.DEBUG,
			Time:    now,
			Message: "test_ltsv_log1",
			Fields: loggan.Fields{
				"first":  1,
				"second": "2",
				"third":  "san",
			},
		}, "level:DEBUG\ttime:" + now.Format(time.RFC3339Nano) + "\tmessage:test_ltsv_log1\tfirst:1\tsecond:2\tthird:san"},
		{&loggan.Entry{
			Level:   loggan.INFO,
			Time:    now,
			Message: "test_ltsv_log2",
		}, "level:INFO\ttime:" + now.Format(time.RFC3339Nano) + "\tmessage:test_ltsv_log2"},
		{&loggan.Entry{
			Level: loggan.WARN,
			Time:  now,
		}, "level:WARN\ttime:" + now.Format(time.RFC3339Nano)},
		{&loggan.Entry{
			Level: loggan.ERROR,
			Time:  now,
		}, "level:ERROR\ttime:" + now.Format(time.RFC3339Nano)},
		{&loggan.Entry{
			Level: loggan.FATAL,
		}, "level:FATAL"},
		{&loggan.Entry{}, "level:NONE"},
	} {
		var buf bytes.Buffer
		formatter := &loggan.LTSVFormatter{}
		if err := formatter.Format(&buf, v.entry); err != nil {
			t.Errorf(`LTSVFormatter.Format(&buf, %#v) => %#v; want %#v`, v.entry, err, nil)
		}
		actual := buf.String()
		expected := v.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(`LTSVFormatter.Format(&buf, %#v); buf => %#v; want %#v`, v.entry, actual, expected)
		}
	}
}

func TestJSONFormatter_Format(t *testing.T) {
	now := time.Now()
	for _, v := range []struct {
		entry    *loggan.Entry
		expected string
	}{
		{&loggan.Entry{
			Level:   loggan.DEBUG,
			Time:    now,
			Message: "test_json_log1",
			Fields: loggan.Fields{
				"first":  1,
				"second": "2",
				"third":  "san",
			},
		}, `{"level":"DEBUG","time":"` + now.Format(time.RFC3339Nano) + `","message":"test_json_log1","first":1,"second":"2","third":"san"}`},
		{&loggan.Entry{
			Level:   loggan.INFO,
			Time:    now,
			Message: "test_json_log2",
		}, `{"level":"INFO","time":"` + now.Format(time.RFC3339Nano) + `","message":"test_json_log2"}`},
		{&loggan.Entry{
			Level: loggan.WARN,
			Time:  now,
		}, `{"level":"WARN","time":"` + now.Format(time.RFC3339Nano) + `"}`},
		{&loggan.Entry{
			Level: loggan.ERROR,
			Time:  now,
		}, `{"level":"ERROR","time":"` + now.Format(time.RFC3339Nano) + `"}`},
		{&loggan.Entry{
			Level: loggan.FATAL,
		}, `{"level":"FATAL"}`},
		{&loggan.Entry{}, `{"level":"NONE"}`},
	} {
		var buf bytes.Buffer
		formatter := &loggan.JSONFormatter{}
		if err := formatter.Format(&buf, v.entry); err != nil {
			t.Errorf(`JSONFormatter.Format(&buf, %#v) => %#v; want %#v`, v.entry, err, nil)
		}
		actual := buf.String()
		expected := v.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(`JSONFormatter.Format(&buf, %#v); buf => %#v; want %#v`, v.entry, actual, expected)
		}
	}
}
