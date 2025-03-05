package config

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"os"
	"log"
)

var WechatPayClient *jsapi.Client

func InitWechatPay() {
	// wechat payment configuration
	mchID := "your_mch_id"          // 商户号
	mchCertificateSerialNumber := "your_mch_certificate_serial_number" // 商户证书序列号
	apiV3Key := "your_api_v3_key"     // APIv3密钥
	privateKeyPath := "path/to/private_key.pem" // 商户私钥文件路径

	// 读取商户私钥
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
	}

	privateKey, err := core.LoadPrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("failed to load private key: %v", err)
	}

	// 读取商户证书
	certificateBytes, err := os.ReadFile("path/to/certificate.pem")
	if err != nil {
		log.Fatalf("failed to read certificate: %v", err)
	}

	certificate, err := core.LoadCertificate(certificateBytes)
	if err != nil {
		log.Fatalf("failed to load certificate: %v", err)
	}

	// 创建认证器
	authenticator := auth.NewWechatPayAuthenticator(
		mchID,
		mchCertificateSerialNumber,
		privateKey,
	)

	// 创建选项
	opts := []core.ClientOption{
		option.WithWechatPayAuthenticator(authenticator),
		option.WithWechatPayCertificates([]*core.Certificate{certificate}),
	}

	// 创建微信支付客户端
	WechatPayClient, err = jsapi.NewClient(opts...)
	if err != nil {
		log.Fatalf("failed to create wechat pay client: %v", err)
	}
}