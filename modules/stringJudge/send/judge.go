package send

import (
	log "github.com/Sirupsen/logrus"
	"strings"
	"github.com/open-falcon/falcon-plus/modules/stringJudge/entity"
)

func SendIfNecessary(events []entity.HEvent) {
	for _, event := range events {
		//解析json
		log.Infof("get event:%s", event)
		//如果是Warning就发送
		if strings.EqualFold("Warning", event.Type) {
			log.Infof("send warning message")
			return
		}
		log.Infof("Normal event")
	}

}
