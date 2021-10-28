package xxl

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"io/ioutil"
)

var jobPathPrefix string = "/xxl-job-new/jobinfo"

func SetJobPathPrefix(pathDefault string) {
	jobPathPrefix = pathDefault
}

var (
	addJobPath   = jobPathPrefix + "/add"
	stopJobPath  = jobPathPrefix + "/stop"
	StartJobPath = jobPathPrefix + "/start"
)

type ExecutorRouteStrategyType string

const (
	//第一个
	FirstExecutorRouteStrategyType ExecutorRouteStrategyType = "FIRST"
	//最后一个
	LastExecutorRouteStrategyType ExecutorRouteStrategyType = "LAST"
	//轮询
	RoundExecutorRouteStrategyType ExecutorRouteStrategyType = "ROUND"
	//随机
	RANDOMExecutorRouteStrategyType ExecutorRouteStrategyType = "RANDOM"
	//一致性HASH
	CONSISTENTHASHExecutorRouteStrategyType ExecutorRouteStrategyType = "CONSISTENT_HASH"
	//最不经常使用
	LeastFrequentlyUsedExecutorRouteStrategyType ExecutorRouteStrategyType = "LEAST_FREQUENTLY_USED"
	//最近最久未使用
	LeastRecentlyUsedExecutorRouteStrategyType ExecutorRouteStrategyType = "LEAST_RECENTLY_USED"
	//故障转移
	FailOverExecutorRouteStrategyType ExecutorRouteStrategyType = "FAILOVER"
	//忙碌转移
	BusyOverExecutorRouteStrategyType ExecutorRouteStrategyType = "BUSYOVER"
	//分片广播
	ShardingBroadcastExecutorRouteStrategyType ExecutorRouteStrategyType = "SHARDING_BROADCAST"
)

//阻塞处理策略
type ExecutorBlockStrategy string

const (
	SerialExecutionBlockStrategy ExecutorBlockStrategy = "SERIAL_EXECUTION" //单机串行
	DiscardLaterBlockStrategy    ExecutorBlockStrategy = "DISCARD_LATER"    //丢弃后续调度
	CoverEarlyBlockStrategy      ExecutorBlockStrategy = "COVER_EARLY"      //覆盖之前调度
)

//{
//		"jobGroup": 45,
//		"jobDesc": "测试02addInfo03",
//		"executorRouteStrategy": "FIRST",
//		"cronGen_display": "*/20 * * * *  ?",
//		"jobCron": "*/20 * * * *  ?",
//		"glueType": "BEAN",
//		"executorHandler": "xsd-task.test3",
//		"executorBlockStrategy": "SERIAL_EXECUTION",
//		"childJobId": "",
//		"executorTimeout": 0,
//		"executorFailRetryCount": 0,
//		"author": "孔振龙",
//		"alarmEmail": "",
//		"executorParam": "{\"id\":99}",
//		"glueRemark": "GLUE代码初始化",
//		"glueSource": ""
//}
type AddJobInfo struct {
	JobGroupID             int                       `json:"jobGroup"`               //任务组id
	ExecutorTimeout        int64                     `json:"executorTimeout"`        // 任务超时时间，单位秒，大于零时生效
	ExecutorFailRetryCount int64                     `json:"executorFailRetryCount"` // 任务超时重试次数
	JobDesc                string                    `json:"jobDesc"`                //任务描述
	ExecutorRouteStrategy  ExecutorRouteStrategyType `json:"executorRouteStrategy"`  //执行策略
	ScheduleType           string                    `json:"scheduleType"`
	CronGenDisplay         string                    `json:"cronGen_display"`       //crontab表达式
	JobCron                string                    `json:"jobCron"`               //crontab表达式
	ChildJobId             string                    `json:"childJobId"`            //子任务id
	Author                 string                    `json:"author"`                //责任人
	AlarmEmail             string                    `json:"alarmEmail"`            //提醒邮件
	ExecutorHandler        string                    `json:"executorHandler"`       //任务标识
	ExecutorParams         string                    `json:"executorParam"`         // 任务参数
	ExecutorBlockStrategy  ExecutorBlockStrategy     `json:"executorBlockStrategy"` // 任务阻塞策略
	GlueType               string                    `json:"glueType"`              // 任务模式，可选值参考 com.xxl.job.core.glue.GlueTypeEnum
	GlueSource             string                    `json:"glueSource"`            // GLUE脚本代码
	GlueRemark             string                    `json:"glueRemark"`            // GLUE脚本标注
}

//	"code": 200,
//	"msg": null,
//	"content": null
type RespAddJob struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Content string `json:"content"`
}

func (e *executor) AddJob(taskInfo AddJobInfo) (respBody []byte, err error) {
	param, err := json.Marshal(taskInfo)
	if err != nil {
		e.log.Error("[err]AddJob:" + err.Error())
		return
	}
	res, err := e.post(addJobPath, string(param))
	//e.log.Info("任务增加ready:" + string(param))
	if err != nil {
		e.log.Error("[err]AddJob err : ", err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e.log.Error("[err]AddJob: ReadAll err : ", err.Error())
		return
	}
	defer res.Body.Close()
	e.log.Info("任务增加成功:" + string(body))
	return body, err
}

func (e *executor) StopJob(jobID int) {
	param, err := json.Marshal(map[string]interface{}{"id": jobID})
	if err != nil {
		e.log.Error("[err]StopJob param:" + err.Error())
		return
	}
	res, err := e.post(stopJobPath, string(param))
	if err != nil {
		e.log.Error("[err]StopJob err : ", err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e.log.Error("[err]StopJob: ReadAll err : ", err.Error())
	}
	e.log.Info("任务停止成功:" + string(body))
}

//启动一个任务
func (e *executor) StartJob(jobID string) (respBody []byte, err error) {
	param := map[string]interface{}{"id": fmt.Sprint(jobID)}
	res, err := e.postForm(StartJobPath, param)
	e.log.Info("任务增加 StartJob:", param)
	if err != nil {
		e.log.Error("[err]StartJob err : ", err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e.log.Error("[err]StartJob: ReadAll err : ", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	e.log.Info("任务启动成功:" + string(body))
	return body, nil
}

//AddJobByPostForm 动态增加任务
func (e *executor) AddJobByPostForm(taskInfo AddJobInfo) (respBody []byte, err error) {
	param := structs.Map(taskInfo)
	res, err := e.postForm(addJobPath, param)
	e.log.Info("任务增加 AddJobByPostForm:", param)

	if err != nil {
		e.log.Error("[err]AddJobByPostForm err : ", err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e.log.Error("[err]AddJobByPostForm: ReadAll err : ", err.Error())
		return
	}
	e.log.Info("任务增加成功:" + string(body))
	defer res.Body.Close()
	return body, err
}
