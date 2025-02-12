package robot

import (
	"os"

	"github.com/yqchilde/wxbot/engine/pkg/log"
)

var configTemplate = `# 机器人WxId，修改为自己的机器人WxId
botWxId: "wxid_lbzw1d5b3pwl29"
# 机器人名字
botNickname: "YY Bot(个人助手)"
# 管理员WxId，多个管理员依次添加
superUsers:
  - "wxid_5cvtxuwufytd21"
# 管理员命令前缀
commandPrefix: "/"
# 本项目运行时的HTTP端口，用于在接入框架里配置回调地址
# 在所接入的VX框架中设置回调地址可写成 http://ip:9528/wxbot/callback，本地测试ip可写localhost
serverPort: 9528

# 接入框架配置
frameworks:
  # 框架选择，可选 千寻、VLW
  name: "千寻"
  # wxbot主动请求微信框架的地址，通常为VX框架设置的HTTP地址，由http://ip:port组成
  apiUrl: "http://192.168.31.8:9527"
  # VX框架HTTP鉴权Token (千寻目前没有，vlw需要)
  apiToken: ""
`

var version string

func init() {
	// 检查配置文件是否存在
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		log.Println("未发现配置文件，已为您生成配置文件，请修改后重新运行程序")
		if err := os.WriteFile("config.yaml", []byte(configTemplate), 0644); err != nil {
			log.Fatalf("生成配置文件失败: %v", err)
		}
		os.Exit(0)
	}

	// 打印版本
	log.Printf("当前运行版本: %s", version)
}
