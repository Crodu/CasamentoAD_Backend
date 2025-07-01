package payments

import (
	"context"
	"fmt"

	"github.com/Crodu/CasamentoBackend/internal/models"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

type ProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func CreatePreference(accessToken string, gift models.Gift) (*preference.Response, error) {
	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Create a new client for the preference service
	client := preference.NewClient(cfg)

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:       gift.Name,
				Quantity:    1,
				UnitPrice:   gift.Price,
				Description: gift.Description,
			},
		},
		BackURLs: &preference.BackURLsRequest{
			Success: "https://www.example.com/success",
			Failure: "https://www.example.com/failure",
			Pending: "https://www.example.com/pending",
		},
		AutoReturn: "approved",
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(resource)

	return resource, nil
}
