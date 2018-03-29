// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/stringJudge/send"
	"github.com/open-falcon/falcon-plus/modules/stringJudge/mysql"
	"fmt"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/open-falcon/falcon-plus/modules/stringJudge/entity"
	"encoding/base64"
)

type StringJudge int

func (this *StringJudge) Ping(req model.NullRpcRequest, resp *model.SimpleRpcResponse) error {
	return nil
}

func (this *StringJudge) Send(items []*model.StringJudgeItem, resp *model.SimpleRpcResponse) error {
	fmt.Printf("receive items:%s", items)
	events := &[]entity.HEvent{}
	for _, item := range items {
		//解析报警信息
		eventString, err := base64.StdEncoding.DecodeString(item.Value)
		if err != nil {
			log.Errorf("decode event error:%s", err)
			continue
		}
		err = json.Unmarshal(eventString, &events)
		if err != nil {
			log.Printf("parse item err:%s", err)
			continue
		}
		//预警判断，发送报警信息
		go send.SendIfNecessary(*events)
		// 存储到mysql中，供dashboard使用
		go mysql.SaveItemToMysql(*events)
	}
	return nil

}
