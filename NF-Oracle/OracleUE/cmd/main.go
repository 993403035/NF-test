package main

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	. "github.com/andy89923/nf-example/OracleUE/TPM"
	"github.com/andy89923/nf-example/internal/sbi/Oracle"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init tpm
	tpm, err := InitTPMSim()
	if err != nil {
		log.Fatalf("Failed to initialize TPM simulator: %v", err)
	}
	defer tpm.Close()

	// 读取 config.toml 配置
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal("配置文件读取失败:", err)
	}
	config := &configs[0]

	// 建立链上连接
	c, err := client.Dial(config)
	if err != nil {
		log.Fatal("连接链失败:", err)
	}

	// 加载合约实例
	contractAddress := common.HexToAddress("0x4d81524657269d51492e3ef7fdc565c2497be375")
	instance, err := Oracle.NewOracleClient(contractAddress, c)
	if err != nil {
		log.Fatal("合约实例加载失败:", err)
	}

	OracleClientSession := &Oracle.OracleClientSession{
		Contract:     instance,
		CallOpts:     *c.GetCallOpts(),
		TransactOpts: *c.GetTransactOpts(),
	}
	// 从当前区块开始监听
	startBlock := uint64(0) // 你可以修改为最新区块，也可以记录上次监听区块高度

	fmt.Println("开始监听合约事件...")
	_, err = OracleClientSession.WatchRequestRandom(&startBlock, func(status int, logs []types.Log) {
		fmt.Println("事件状态码:", status)

		for idx, v := range logs {
			parsed, err := OracleClientSession.ParseRequestRandom(v)
			//fmt.Println(parsed.Oraclenode.String())
			if err != nil {
				fmt.Println("事件解析错误:", err)
				continue
			}
			fmt.Printf("第 %d 条日志: %v\n", idx, parsed)
			//Oracle查看是否请求的是自己的地址, if true, callback
			if common.HexToAddress(parsed.Oraclenode.String()) == common.HexToAddress("0x83309d045a19c44dc3722d15a6abd472f95866ac") {
				//use tpm generate random value
				randomBytes, err := GetRandomBytes(tpm, 6)
				if err != nil {
					log.Fatalf("Random generation error: %v", err)
				}
				fmt.Printf("Random bytes: %x\n", randomBytes)
				sigernum := big.NewInt(3)
				randomnumber := new(big.Int).SetBytes(randomBytes)
				ret, _, _, err := OracleClientSession.CallbackRandom(randomnumber, "2bc9403fc69cee7796d1489a84db7f6677419f7ba13f5cdb33017e69af31e3bbe7a92e2dd796b54fe32142d0692dd11df5aecfb5c8ab37d8401145531fea261006edaec31a0ab9ac27213fb88e1ebc58e3756f4df9a12362f1ac8b0bcb95527e23978db18e2d8d7e8ea52d43b7d18e28fbde097b1516070dc8cd28e0320bbbe2,2c89dd5e31dc4f475b35f85cae6ac4a8e07f55c0afde0ccce23a6c624e2c0e185f7e6c28b254b3ed55bba5c315594c404b64beacaed4598feedceb7dd9c3033e496a082b6bdc653703e97af0d45af751bcad586be6c48d67a8f46186ff27d0dd9d4c6984cd11b315417fda1e53b819762d0306619cafce6b6c167d92e98949cb,027e4f28b767aa14f65cc17b0c8557df2a55d8672a7d35fc1c7979096e330cbd5a66ee1e0c4fe3cb45316b355655444403b0c57f4ebc28312fd4ec73abd0069e5e8475b7aea8866b47e71e9079a70eeb1a528458ce75100e33ebb647bf379c15043c4c06c2b02c65619c72c9be90968b4f3d3167d73ee69e1241a3029b579945", "7296b3865f9b66196a1946a2dc0e30b949b575700e0a46cf7415ad6f0e198e5b3b6091443904f6cec1b3b1469e39f3c4135ab5494463ca412ca7032ee327925252954400d0870e7bccc45526ea2de72454293c13a5b273aa5678481c51cf029e706f74b883c903f078f606788548ba47c98b46824ae6732cef4608606354db4d", sigernum)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("call back result: %s\n", ret)
			}
		}
	})
	if err != nil {
		log.Fatal("监听器创建失败:", err)
	}

	// 添加对 RequestHTTP 的监听
	_, err = OracleClientSession.WatchRequestHttpevent(&startBlock, func(status int, logs []types.Log) {
		fmt.Println("HTTP请求事件状态码:", status)

		for idx, v := range logs {
			parsed, err := OracleClientSession.ParseRequestHttpevent(v)
			if err != nil {
				fmt.Println("HTTP事件解析错误:", err)
				continue
			}
			fmt.Printf("第 %d 条HTTP请求日志: %v\n", idx, parsed)

			// 检查是否是发给当前Oracle节点
			if common.HexToAddress(parsed.Oraclenode.String()) == common.HexToAddress("0x83309d045a19c44dc3722d15a6abd472f95866ac") {
				// 发起HTTP请求
				resp, err := http.Get(parsed.Domain)
				if err != nil {
					log.Printf("HTTP请求失败: %v", err)
					continue
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("读取HTTP响应失败: %v", err)
					continue
				}
				fmt.Printf("HTTP响应: %s\n", string(body))

				// 你可以在这里签名响应并发送回链上，下面为伪代码
				sigernum := big.NewInt(3) // 签名版本号
				ret, _, _, err := OracleClientSession.CallbackHttp(
					string(body), // 假设合约接受字符串结果
					"2bc9403fc69cee7796d1489a84db7f6677419f7ba13f5cdb33017e69af31e3bbe7a92e2dd796b54fe32142d0692dd11df5aecfb5c8ab37d8401145531fea261006edaec31a0ab9ac27213fb88e1ebc58e3756f4df9a12362f1ac8b0bcb95527e23978db18e2d8d7e8ea52d43b7d18e28fbde097b1516070dc8cd28e0320bbbe2,2c89dd5e31dc4f475b35f85cae6ac4a8e07f55c0afde0ccce23a6c624e2c0e185f7e6c28b254b3ed55bba5c315594c404b64beacaed4598feedceb7dd9c3033e496a082b6bdc653703e97af0d45af751bcad586be6c48d67a8f46186ff27d0dd9d4c6984cd11b315417fda1e53b819762d0306619cafce6b6c167d92e98949cb,027e4f28b767aa14f65cc17b0c8557df2a55d8672a7d35fc1c7979096e330cbd5a66ee1e0c4fe3cb45316b355655444403b0c57f4ebc28312fd4ec73abd0069e5e8475b7aea8866b47e71e9079a70eeb1a528458ce75100e33ebb647bf379c15043c4c06c2b02c65619c72c9be90968b4f3d3167d73ee69e1241a3029b579945",
					"7296b3865f9b66196a1946a2dc0e30b949b575700e0a46cf7415ad6f0e198e5b3b6091443904f6cec1b3b1469e39f3c4135ab5494463ca412ca7032ee327925252954400d0870e7bccc45526ea2de72454293c13a5b273aa5678481c51cf029e706f74b883c903f078f606788548ba47c98b46824ae6732cef4608606354db4d",
					sigernum,
				)
				if err != nil {
					log.Printf("回调HTTP响应失败: %v", err)
				} else {
					fmt.Printf("call back result: %s\n", ret)
					fmt.Println("HTTP响应已回传区块链")
				}
			}
		}
	})
	if err != nil {
		log.Fatal("HTTP监听器创建失败:", err)
	}

	// 阻塞主线程直到退出
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	fmt.Println("事件监听器已关闭。")
}
