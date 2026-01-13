package TPM

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"github.com/google/go-tpm/tpmutil"
	"io"

	"github.com/google/go-tpm-tools/client"
	"github.com/google/go-tpm-tools/simulator"
	"github.com/google/go-tpm/legacy/tpm2"
)

// 初始化 TPM 模拟器
func InitTPMSim() (io.ReadWriteCloser, error) {
	return simulator.Get()
}

// 生成 ECC 签名密钥
func GenerateKey(tpm io.ReadWriteCloser) (*client.Key, error) {
	template := tpm2.Public{
		Type:    tpm2.AlgECC,
		NameAlg: tpm2.AlgSHA256,
		Attributes: tpm2.FlagSign | tpm2.FlagFixedTPM |
			tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth,
		ECCParameters: &tpm2.ECCParams{
			CurveID: tpm2.CurveNISTP256,
			Sign: &tpm2.SigScheme{
				Alg:  tpm2.AlgECDSA,
				Hash: tpm2.AlgSHA256,
			},
		},
	}
	return client.NewKey(tpm, tpm2.HandleOwner, template)
}

// 签名消息并验证签名
func SignAndVerify(tpm io.ReadWriteCloser, key *client.Key, message []byte) error {
	hashAlg := crypto.SHA256
	hash := hashAlg.New()
	hash.Write(message)
	digest := hash.Sum(nil)

	sig, err := key.SignData(message)
	if err != nil {
		return fmt.Errorf("failed to sign data: %w", err)
	}
	fmt.Printf("Signature: %x\n", sig)

	pub := key.PublicKey().(*ecdsa.PublicKey)
	if !ecdsa.VerifyASN1(pub, digest, sig) {
		return fmt.Errorf("signature verification failed")
	}
	fmt.Println("Signature verified successfully.")
	return nil
}

// 执行 PCR 测量并返回测量值
func GetMeasurement(tpm io.ReadWriteCloser) (string, error) {

	// Read the system state
	// 测量的数据（例如：一段固件或配置）
	// 测量内核镜像, 例如：(考虑到具体部署操作系统未知，这里用字符串模拟内核文件内容)
	//kernelData, _ := os.ReadFile("/boot/vmlinuz")
	kerneldata := []byte("critical system state")
	hash := sha256.Sum256(kerneldata)
	//fmt.Println("kernal hash:", kernelhash)

	pcrIndex := 16
	pcrHandle := tpmutil.Handle(pcrIndex)

	if err := tpm2.PCRReset(tpm, pcrHandle); err != nil {
		return "", fmt.Errorf("PCR reset failed: %w", err)
	}
	if err := tpm2.PCRExtend(tpm, pcrHandle, tpm2.AlgSHA256, hash[:], ""); err != nil {
		return "", fmt.Errorf("PCR extend failed: %w", err)
	}
	pcrValue, err := tpm2.ReadPCR(tpm, pcrIndex, tpm2.AlgSHA256)
	if err != nil {
		return "", fmt.Errorf("read PCR failed: %w", err)
	}
	return fmt.Sprintf("%x", pcrValue), nil
}

// 从 TPM 获取 n 字节随机数
func GetRandomBytes(tpm io.ReadWriteCloser, n int) ([]byte, error) {

	randomValue, err := tpm2.GetRandom(tpm, uint16(n))

	if err != nil {
		fmt.Errorf("GetRandom failed: %w", err)
	}
	return randomValue, err
}
