package dto

import (
	"encoding/json"
	"io"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// FromJSONCreateProductRequest parses the JSON data from the given io.Reader and creates a new CreateProductRequest object.
//
// Parameters:
// - body: The io.Reader containing the JSON data to be parsed.
//
// Returns:
// - *CreateProductRequest: A pointer to the newly created CreateProductRequest object.
// - error: An error object if there was an error parsing the JSON data.
func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	createProductRequest := CreateProductRequest{}
	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}

	return &createProductRequest, nil
}
