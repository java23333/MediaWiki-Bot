package MessagePushAPI

import (
	"strconv"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI/KaiHeiLaAPI"
	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI/LineAPI"
	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI/QQAPI"
	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI/TelegramAPI"
	"github.com/nyancatda/ShionBot/src/utils/Language"
	"github.com/nyancatda/ShionBot/src/utils/ViolationWordFilter"
)

//发送消息
//SNSName 聊天软件
//ChatType 聊天类型(Default,Friend,Group,GroupAt,Temp)
//UserID 接收消息的用户ID(仅用于判断用户)
//target 接受的聊天的ID
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为空即可)
//AtID (可选)需要@的人的ID
//group (可选)临时会话群号
func SendMessage(SNSName string, ChatType string, UserID string, target string, text string, quote bool, quoteID string, AtID string, group int) {
	if !ViolationWordFilter.ViolationWordFilter(text) {
		text = Language.Message(SNSName, UserID).ContainsIllegalWords
	}

	switch SNSName {
	case "QQ":
		targets, _ := strconv.Atoi(target)
		switch ChatType {
		case "Friend":
			quoteIDs, _ := strconv.Atoi(quoteID)
			QQAPI.SendFriendMessage(targets, text, quote, quoteIDs)
		case "Group":
			quoteIDs, _ := strconv.Atoi(quoteID)
			QQAPI.SendGroupMessage(targets, text, quote, quoteIDs)
		case "GroupAt":
			AtID, _ := strconv.Atoi(AtID)
			QQAPI.SendGroupAtMessage(targets, text, AtID)
		case "Temp":
			quoteIDs, _ := strconv.Atoi(quoteID)
			QQAPI.SendTempMessage(targets, group, text, quote, quoteIDs)
		}
	case "Telegram":
		targets, _ := strconv.Atoi(target)
		switch ChatType {
		case "GroupAt":
			text = "@" + AtID + " " + text
			TelegramAPI.SendMessage("Group", targets, text, true, false, 0, false)
		case "Group":
			quoteIDs, _ := strconv.Atoi(quoteID)
			TelegramAPI.SendMessage("Group", targets, text, true, false, quoteIDs, quote)
		default:
			quoteIDs, _ := strconv.Atoi(quoteID)
			TelegramAPI.SendMessage("Friend", targets, text, true, false, quoteIDs, quote)
		}
	case "Line":
		switch ChatType {
		case "GroupAt":
			text = "@" + AtID + " " + text
			LineAPI.SendPushMessage(ChatType, target, text, false)
		case "Group":
			if quote {
				LineAPI.SendReplyMessage(ChatType, quoteID, text, false)
			} else {
				LineAPI.SendPushMessage(ChatType, target, text, false)
			}
		default:
			LineAPI.SendPushMessage(ChatType, target, text, false)
		}
	case "KaiHeiLa":
		switch ChatType {
		case "Group":
			KaiHeiLaAPI.SendChannelMessage("Channel", 1, target, text, quote, quoteID)
		case "Friend":
			KaiHeiLaAPI.SendDirectMessage(ChatType, 1, target, text, quote, quoteID)
		}
	}
}

//(QQ)发送头像戳一戳
//target 目标QQ号
//subject 消息接受主体，为群号或QQ号
//kind 上下文类型,可选值 Friend,Group,Stranger
func SendNudge(target int, subject int, kind string) {
	QQAPI.SendNudge(target, subject, kind)
}
