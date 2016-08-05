package vars

import (
	"pandora/modules/conf"
)
//使用vars.Conf获取配置读取对象
//conf
var (
	TemplatePath string //对应配置文件的 template_path
	Listen string //对应配置文件的 listen
	WebRoot string //对应配置文件的 web_root
	Conf         conf.Configuration

)
