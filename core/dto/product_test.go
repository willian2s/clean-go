package dto_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"github.com/willian2s/clean-go/core/dto"
)

func TestFromJSONCreateProductRequest(t *testing.T) {
	fakeItem := dto.CreateProductRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONCreateProductRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, fakeItem.Name, itemRequest.Name)
	require.Equal(t, fakeItem.Price, itemRequest.Price)
	require.Equal(t, fakeItem.Description, itemRequest.Description)
}

func TestFromJSONCreateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateProductRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
