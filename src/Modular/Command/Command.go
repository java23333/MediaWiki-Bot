package Command

import (
	"strings"

	"github.com/nyancatda/ShionBot/src/Modular"
	"github.com/nyancatda/ShionBot/src/Struct"
)

func Command(SNSName string, Messagejson Struct.WebHookJson, CommandText string) (string, bool) {
	UserID := Modular.GetSNSUserID(SNSName, Messagejson)

	if find := strings.Contains(CommandText, "language"); find {
		return LanguageSettings(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "help"); find {
		return Help(SNSName, UserID)
	}
	if find := strings.Contains(CommandText, "wikiadd"); find {
		return WikiAdd(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "wikiupdate"); find {
		return WikiUpdate(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "wikidelete"); find {
		return WikiDelete(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "userinfo"); find {
		return UserInfo(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "importdata"); find {
		return ImportData(SNSName, UserID, CommandText)
	}
	return "", false
}
