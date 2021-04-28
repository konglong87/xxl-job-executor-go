package xxl

import (
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestFormatTimeToCronTab(t *testing.T) {
	t1 := time.Now()
	t.Log(t1)
	t.Log(FormatTimeToCronTab(t1))

	t2 := "2021-04-12 12:00:24.907769 +0800 CST"
	t2ExpectedCrontab := "24 00 12 12 04 ? 2021-2021"

	t2Val,_ := time.Parse("2006-01-02 15:04:05.999999999 +0800 CST", t2)
	t.Log(FormatTimeToCronTab(t2Val) == t2ExpectedCrontab)
	assert.Equal(t, t2ExpectedCrontab, FormatTimeToCronTab(t2Val))
}
