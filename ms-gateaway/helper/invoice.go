package helper

import (
	"bytes"
	"encoding/json"
	pb "ms-gateaway/pb/order"
	"net/http"
	"os"
)

type Invoice struct {
	ID         string `json:"id"`
	InvoiceUrl string `json:"invoice_url"`
}

func CreateInvoiceCheckout(res *pb.CheckoutOrderResponse) (string, error) {
	apiKey := os.Getenv("XENDIT_API_KEY")
	apiUrl := "https://api.xendit.co/v2/invoices"

	data := []map[string]interface{}{}
	for _, p := range res.Products {
		d := map[string]interface{}{}
		d["name"] = p.ProductName
		d["quantity"] = p.Quantity
		d["price"] = p.Price
		data = append(data, d)
	}

	bodyRequest := map[string]interface{}{
		"external_id":      "1",
		"amount":           res.TotalPrice,
		"description":      "Dummy Invoice Final Project",
		"invoice_duration": 86400,
		"currency":         "IDR",
		"items":            data,
	}

	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	request.SetBasicAuth(apiKey, "")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	resInvoice := Invoice{}
	if err := json.NewDecoder(response.Body).Decode(&resInvoice); err != nil {
		return "", err
	}

	return resInvoice.InvoiceUrl, nil
}

func CreateInvoiceTopUp(amount float64) (string, error) {
	apiKey := os.Getenv("XENDIT_API_KEY")
	apiUrl := "https://api.xendit.co/v2/invoices"

	bodyRequest := map[string]interface{}{
		"external_id":      "1",
		"amount":           amount,
		"description":      "Dummy Invoice Mini Project",
		"invoice_duration": 86400,
		"currency":         "IDR",
	}

	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	request.SetBasicAuth(apiKey, "")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	resInvoice := Invoice{}
	if err := json.NewDecoder(response.Body).Decode(&resInvoice); err != nil {
		return "", err
	}

	return resInvoice.InvoiceUrl, nil
}
