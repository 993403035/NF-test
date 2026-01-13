package processor

import (
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/andy89923/nf-example/internal/sbi/Oracle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//func (p *Processor) FindSpyFamilyCharacterName(c *gin.Context, targetName string) {
//	if lastName, ok := p.Context().SpyFamilyData[targetName]; ok {  // 通过NfApp掉用NFContext的数据
//		c.String(http.StatusOK, fmt.Sprintf("Character: %s %s", targetName, lastName))
//		return
//	}
//	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not found in SPYxFAMILY", targetName))
//}

// 用processor的处理的意义是，通过NfApp统一调用NFContext中的数据
func (p *Processor) ChallengeHttp(c *gin.Context, targetUrl string) {

	// 初始化合约方法
	// 读取 config.toml 配置
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal("配置文件读取失败:", err)
	}
	config := &configs[0]

	// 建立链上连接
	cl, err := client.Dial(config)
	if err != nil {
		log.Fatal("连接链失败:", err)
	}

	// 加载合约实例
	contractAddress := common.HexToAddress("0x4d81524657269d51492e3ef7fdc565c2497be375")
	instance, err := Oracle.NewOracleClient(contractAddress, cl)
	if err != nil {
		log.Fatal("合约实例加载失败:", err)
	}
	OracleClientSession := &Oracle.OracleClientSession{
		Contract:     instance,
		CallOpts:     *cl.GetCallOpts(),
		TransactOpts: *cl.GetTransactOpts(),
	}

	//向目标域名发起http请求，获取响应

	//GetHttp := "res for " + targetUrl
	nodeaddr := common.HexToAddress("0x83309d045a19c44dc3722d15a6abd472f95866ac")
	_, _, err = OracleClientSession.RequestHttp(nodeaddr, targetUrl)
	if err != nil {
		log.Fatal(err)
	}
	GetHttp := "request oracle to challenge: " + targetUrl + "\n waiting for callback"
	c.String(http.StatusOK, GetHttp)
	return
}
