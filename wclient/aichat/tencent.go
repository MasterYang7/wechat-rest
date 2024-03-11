package aichat

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/liudding/go-llm-api/tencent"
)

func TencentText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	keys := strings.Split(llmc.Secret, ",")
	if len(keys) != 3 {
		return "", errors.New("密钥格式错误")
	}

	appId, err := strconv.ParseInt(keys[0], 10, 64)
	if err != nil {
		return "", errors.New("AppID错误")
	}

	// 初始化模型

	config := tencent.DefaultConfig(appId, keys[1], keys[2])

	if len(llmc.Endpoint) > 1 {
		config.BaseURL = llmc.Endpoint
	}

	client := tencent.NewClientWithConfig(config)

	req := tencent.ChatCompletionRequest{
		Messages: []tencent.ChatCompletionMessage{},
	}

	// 设置上下文

	if llmc.RoleContext != "" {
		req.Messages = []tencent.ChatCompletionMessage{
			{Content: llmc.RoleContext, Role: tencent.ChatMessageRoleUser},
			{Content: "OK", Role: tencent.ChatMessageRoleAssistant},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		if role == "user" {
			role = tencent.ChatMessageRoleUser
		}
		if role == "model" {
			role = tencent.ChatMessageRoleAssistant
		}
		req.Messages = append(req.Messages, tencent.ChatCompletionMessage{
			Content: msg.Content, Role: role,
		})
	}

	req.Messages = append(req.Messages, tencent.ChatCompletionMessage{
		Content: ask, Role: tencent.ChatMessageRoleUser,
	})

	// 请求模型接口

	res, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	if len(res.Choices) <= 0 {
		return "", errors.New("未得到预期的结果")
	}

	reply := ""

	reply += res.Choices[0].Messages.Content

	// 更新历史记录
	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: reply, Role: "model"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
