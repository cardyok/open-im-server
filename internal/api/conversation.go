// Copyright © 2023 OpenIM. All rights reserved.
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

package api

import (
	"github.com/OpenIMSDK/tools/apiresp"
	"github.com/OpenIMSDK/tools/checker"
	"github.com/OpenIMSDK/tools/errs"
	"github.com/OpenIMSDK/tools/log"
	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/protocol/conversation"
	"github.com/OpenIMSDK/tools/a2r"

	"github.com/openimsdk/open-im-server/v3/pkg/rpcclient"
)

type ConversationApi rpcclient.Conversation

func NewConversationApi(client rpcclient.Conversation) ConversationApi {
	return ConversationApi(client)
}

func (o *ConversationApi) GetAllConversations(c *gin.Context) {
	a2r.Call(conversation.ConversationClient.GetAllConversations, o.Client, c)
}

func (o *ConversationApi) GetSortedConversationList(c *gin.Context) {
	a2r.Call(conversation.ConversationClient.GetSortedConversationList, o.Client, c)
}

func (o *ConversationApi) GetConversation(c *gin.Context) {
	a2r.Call(conversation.ConversationClient.GetConversation, o.Client, c)
}

func (o *ConversationApi) GetConversations(c *gin.Context) {
	a2r.Call(conversation.ConversationClient.GetConversations, o.Client, c)
}

func (o *ConversationApi) SetConversations(c *gin.Context) {
	a2r.Call(conversation.ConversationClient.SetConversations, o.Client, c)
}
func (o *ConversationApi) DelConversations(c *gin.Context) {
	var req conversation.GetConversationReq
	if err := c.BindJSON(&req); err != nil {
		log.ZWarn(c, "gin bind json error", err, "req", req)
		apiresp.GinError(c, errs.ErrArgs.WithDetail(err.Error()).Wrap()) // 参数错误
		return
	}
	if err := checker.Validate(&req); err != nil {
		apiresp.GinError(c, err) // 参数校验失败
		return
	}
	if req.ConversationID != "" {
		req.ConversationID = "&" + req.ConversationID
	}
	data, err := o.Client.GetConversation(c, &req)
	if err != nil {
		apiresp.GinError(c, err) // RPC调用失败
		return
	}
	apiresp.GinSuccess(c, data) // 成功
}
func (o *ConversationApi) GetConversationOfflinePushUserIDs(c *gin.Context) {
	a2r.Call(conversation.ConversationClient.GetConversationOfflinePushUserIDs, o.Client, c)
}
