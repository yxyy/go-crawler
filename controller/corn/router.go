package corn

import (
	"github.com/robfig/cron"
	"lhc.go.crawler/controller/corn/netbian"
	"lhc.go.crawler/logs"
)

func InitCornRouter()  {

	c := cron.New()

	if err := c.AddFunc("* * * 1 1 1", netbian.NewNetbian().Run);err!=nil{
		logs.Error.Info(err)
	}
}
