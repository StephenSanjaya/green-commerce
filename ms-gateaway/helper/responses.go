package helper

// HTTPError is a custom error response struct for swagger documentation
type HTTPError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
