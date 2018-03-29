package mysql

import (
	"github.com/open-falcon/falcon-plus/modules/stringJudge/g"
	log "github.com/Sirupsen/logrus"
	"time"
	"github.com/open-falcon/falcon-plus/modules/stringJudge/entity"
)

func SaveItemToMysql(events []entity.HEvent) {
	for _, event := range events {
		//解析
		dbConn, err := g.GetDbConn("stringJudge")
		if err != nil {
			log.Errorf("%s make dbConn fail", err)
			return
		}

		// endpoint表
		sqlStr := `INSERT INTO string_events(type, reason, message,involved_object,source,count,first_time,last_time)
		VALUES (?,?,?,?,?,?,?,?)`
		//解析时间  2018-03-23T02:13:53Z
		layout := "2006-01-02T15:04:05Z"
		firstTime, err := time.Parse(layout, event.FirstTimestamp)
		if err != nil {
			log.Errorf("parse time err:%s", err)
			return
		}
		lastTime, err := time.Parse(layout, event.LastTimestamp)
		if err != nil {
			log.Errorf("parse time err:%s", err)
			return
		}

		_, err = dbConn.Exec(sqlStr, event.Type, event.Reason, event.Message, event.InvolvedObject.Name, event.Source.Component, event.Count, firstTime, lastTime)
		if err != nil {
			log.Errorf("exec sql error:%s", err)
			return
		}
	}
}
