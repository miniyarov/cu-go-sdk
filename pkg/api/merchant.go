package api

import (
	"net/http"
)

// MerchantAPI interacts with merchant API
type MerchantAPI struct {
	client      Client
	merchantKey string
	secretKey   string
}

// NewMerchantAPI constructs a new merchant API
func NewMerchantAPI(client Client, merchantKey string, secretKey string) *MerchantAPI {
	return &MerchantAPI{
		client:      client,
		merchantKey: merchantKey,
		secretKey:   secretKey,
	}
}

// InvoiceInfo retrieves an invoice information and its current status
func (m *MerchantAPI) InvoiceInfo(invoiceHash string) (interface{}, error) {
	query := map[string]string{"invoice-hash": invoiceHash}
	headers := map[string]string{"X-Merchant-Key": m.merchantKey, "X-Secret-Key": m.secretKey}
	req, err := m.client.NewRequest(http.MethodGet, "/merchant/invoice-inf", headers, query, nil)
	if err != nil {
		return nil, err
	}

	return m.client.DoRequest(req)
}
