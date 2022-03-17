package api_test

import (
	"testing"

	"cryptounifier-go-sdk/pkg/api"
)

func TestMerchantAPI_InvoiceInfo(t *testing.T) {
	// Arrange
	response := `{"message":{"hash":"CFEFKEMTOK","title":"New Invoice","description":"Pay for a product or service","currency":"usd","target_value":"15.000000","status":0}}`
	client := newMockClient(response)
	merchantKey := "MERCHANT_KEY"
	secretKey := "SECRET_KEY"
	merchantAPI := api.NewMerchantAPI(client, merchantKey, secretKey)
	invoiceHash := "CFEFKEMTOK"

	// Act
	resp, err := merchantAPI.InvoiceInfo(invoiceHash)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, ok := resp.(map[string]interface{})["message"]; !ok {
		t.Errorf("invalid response")
	}
}
