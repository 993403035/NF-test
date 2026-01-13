package context

import (
	"os"

	"github.com/andy89923/nf-example/internal/logger"
	"github.com/andy89923/nf-example/pkg/factory"
	"github.com/google/uuid"

	"github.com/free5gc/openapi/models"
)

// context里面定义的是整个NF在全生命周期里都会存储和使用的数据
type NFContext struct {
	NfId        string
	Name        string
	UriScheme   models.UriScheme
	BindingIPv4 string
	SBIPort     int
	//DRoT里用到的数据
	OracleNodeID string
	RandomNumber int
	HttpResponse string
}

var nfContext = NFContext{}

func InitNfContext() {
	cfg := factory.NfConfig

	nfContext.NfId = uuid.New().String()
	nfContext.Name = "DRoT" //

	nfContext.UriScheme = cfg.Configuration.Sbi.Scheme
	nfContext.SBIPort = cfg.Configuration.Sbi.Port
	nfContext.BindingIPv4 = os.Getenv(cfg.Configuration.Sbi.BindingIPv4)
	if nfContext.BindingIPv4 != "" {
		logger.CtxLog.Info("Parsing ServerIPv4 address from ENV Variable.")
	} else {
		nfContext.BindingIPv4 = cfg.Configuration.Sbi.BindingIPv4
		if nfContext.BindingIPv4 == "" {
			logger.CtxLog.Warn("Error parsing ServerIPv4 address as string. Using the 0.0.0.0 address as default.")
			nfContext.BindingIPv4 = "0.0.0.0"
		}
	}
	//用于响应外部请求的数据
	nfContext.OracleNodeID = "OracleNode 1"
	nfContext.RandomNumber = 123
	nfContext.HttpResponse = "Hello DRoT HttpResponse"

}

func GetSelf() *NFContext {
	return &nfContext
}
