package controllers

import (
	"dymall/config"
	"dymall/server/models"

	"net/http"
	"github.com/gin-gonic/gin"
)

func ProcessPayment(c *gin.Context) {
	var paymentRequest models.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// pay method
	/*
	if paymentRequest.PaymentMethod != "credit_card" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported payment method"})
		return
	}
	*/

	paymentResult := Pay(paymentRequest.Amount)
	if !paymentResult {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "payment failed"})
		return
	}

	// update order status
	var order models.Order
	result := config.DB.First(&order, paymentRequest.OrderID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	order.Status = "paid"
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "payment successful", "order_id": order.ID, "status": order.Status})
}

// pay the bill
func Pay(amount float64) bool {
	return true
}
/*
func Pay(amount float64) bool {
	// 创建预支付订单请求
	prepayRequest := jsapi.PrepayRequest{
		Appid:       "my_app_id", // 应用ID
		Mchid:       "my_mch_id", // 商户号
		Description: "Order Payment",
		OutTradeNo:  time.Now().Format("20060102150405"), // 商户订单号
		Amount: &jsapi.Amount{
			Total: int(amount * 00), // 金额单位为分
		},
		Payer: &jsapi.Payer{
			Openid: "user_openid", // 用户openid
		},
	}

	prepayResponse, err := config.WechatPayClient.Prepay(c, &prepayRequest)
	if err != nil {
		return nil, err
	}

	return prepayResponse, nil
}
*/
