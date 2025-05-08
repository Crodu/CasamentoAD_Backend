package payments

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func GeneratePayment(value float64, email string, firstName string, lastName string, product string, token string) (*payment.Response, error) {
	accessToken := token

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := payment.NewClient(cfg)

	request := payment.Request{
		TransactionAmount: value,
		PaymentMethodID:   "pix",
		Description:       product,
		Payer: &payment.PayerRequest{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resource, nil
}

func GetQRCode(paymentInfo *payment.Response) (string, error) {
	if paymentInfo.PointOfInteraction.TransactionData.QRCode == "" {
		return "", fmt.Errorf("QR code not found in payment info")
	}
	return paymentInfo.PointOfInteraction.TransactionData.QRCode, nil
}
