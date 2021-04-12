package xxl

import (
	"fmt"
	"time"
)

func main(){
	t1 := time.Now()
	fmt.Println(t1.Format("2013-2013 01 02 15 04 05"))
	fmt.Println(t1.Format("2013 01 02 15 04 05"))
	fmt.Println(t1.Format("2006  01 02 15 04 05"))
	fmt.Println(t1.Format("2006 ? 01 02 15 04 05"))
	fmt.Println(t1.Format("2006-2006 ? 01 02 15 04 05"))
	fmt.Println(t1.Format("05 04 15 02 01 ? 2006-2006"))
}

const(
	//单次指定时间执行, crontab表达式
	cronTabFormatSingleTime = "05 04 15 02 01 ? 2006-2006"
)
func FormatTimeToCronTab(t time.Time)(cronExpr string){
	return t.Format(cronTabFormatSingleTime)
}
