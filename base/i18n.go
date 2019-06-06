package base

import (
	"github.com/Unknwon/i18n"
)

func init() {
	i18n.SetMessage("zh-CN", "conf/lang/lang_zh-CN.ini")
}

func Tr(lang, key string) string {
	return i18n.Tr(lang, key)
}
