package command

import (
	"fmt"
	"strings"
)

const (
	DEFAULT_LARK_CHAT_FORMAT = "%s://applink.feishu.cn/client/chat/open?openId=%s"
	APP_LARK                 = "lark"
	LARK_CHAT_CMD            = "lark_chat_"
)

type LarkChatCMD struct {
	format    string
	commandID string

	attrs map[string]string
}

func NewLarkChatCMD(commandID string, attrs map[string]string) *LarkChatCMD {
	return &LarkChatCMD{
		format:    DEFAULT_LARK_CHAT_FORMAT,
		commandID: commandID,
		attrs:     attrs,
	}
}

func (oc *LarkChatCMD) GenURI() string {
	return fmt.Sprintf(oc.format, APP_LARK, oc.commandID)
}

func (oc *LarkChatCMD) IconApp() string {
	return APP_LARK
}

func (oc *LarkChatCMD) Filtered(keys []string) (string, string, bool) {
	for k, v := range oc.attrs {
		for _, query := range keys {
			if !strings.Contains(k, query) && !strings.Contains(v, query) {
				return "", "", false
			}
		}
		return k, v, true
	}
	return "", "", false
}
