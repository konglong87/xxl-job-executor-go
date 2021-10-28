package xxl

import (
	"fmt"
	"testing"
	"time"
)

func TestExecutor_AddJobByPostForm(t *testing.T) {
	taskInfo := AddJobInfo{
		JobGroupID:             2,
		JobDesc:                "oaa-service" + "_" + fmt.Sprintf("task_id_%d_%d", 87, time.Now().Unix()),
		ExecutorRouteStrategy:  FirstExecutorRouteStrategyType,
		CronGenDisplay:         FormatTimeToCronTab(time.Now().Add(20 * time.Minute)),
		JobCron:                FormatTimeToCronTab(time.Now().Add(20 * time.Minute)),
		ChildJobId:             "",
		Author:                 "yyg",
		AlarmEmail:             "",
		ExecutorHandler:        "runTaskHandler",
		ExecutorParams:         "{\"id\":12345}",
		ExecutorBlockStrategy:  SerialExecutionBlockStrategy,
		ExecutorTimeout:        0,
		ExecutorFailRetryCount: 0,
		GlueType:               "BEAN",
		ScheduleType:           "CRON",
		ScheduleConf:           FormatTimeToCronTab(time.Now().Add(20 * time.Minute)),
		MisfireStrategy:        MisfireStrategyNothing,
	}
	ne := NewExecutor(
		ServerAddr("https://tms.testing.pipacoding.com/xxl-job-new"),
		RegistryKey("runTaskHandler"),
		AccessToken("xxl-job-new-testing"),
	)
	ne.Init()
	ex, err := ne.AddJobByPostForm(taskInfo)
	t.Log(err)
	t.Log(string(ex))
}
