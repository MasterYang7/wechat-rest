package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入目录

var Efs *embed.FS

// 版本信息

const Version = "39.4.2"
const BuildVersion = "250316"

// 应用描述

const AppName = "WeChat"
const AppSummary = "微信消息管家"
