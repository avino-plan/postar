// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/13 00:42:12

package http

import (
	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// pingHandler tells you if postar is ready or not.
func pingHandler(ctx iris.Context) {
	ctx.Write([]byte(`<h1 style="text-align: center;">Pong!</h1><h3 style="text-align: center;">- Postar is ready! -</h3>`))
}

// sendHandler handles the service of sending emails.
func sendHandler(ctx context.Context) {

	// 获取邮件发送任务信息
	sendTask := models.NewEmptySendTask()
	err := ctx.ReadJSON(&sendTask)
	if err != nil {
		core.Logger().Errorf("The error is %s.", err.Error())
		ctx.StatusCode(400)
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Write(models.WrongRequestBodyResponse())
		return
	}

	// 发送邮件
	email := core.NewEmail(sendTask.To, sendTask.Subject, sendTask.ContentType, sendTask.Body)
	err = core.Send(email)
	if err != nil {
		core.Logger().Errorf("The error is %s. The information of this email is %+v.", err.Error(), sendTask)
		ctx.StatusCode(500)
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Write(models.FailedToSendEmailResponse())
		return
	}

	core.Logger().Debugf("Email %+v successfully sent.", email)
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Write(models.EmailSuccessfullySentResponse())
}
