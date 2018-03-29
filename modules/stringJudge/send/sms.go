package send

import (
	"github.com/golang/glog"
	"github.com/open-falcon/falcon-plus/modules/stringJudge/entity"
)

func SendSms(event *entity.HEvent) {
	glog.Infof("send sms ,info:%s", event)
}
