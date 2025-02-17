package openapi

import (
	"bytes"
	"encoding/json"
)

var _ MappedNullable = &DeleteDocumentsRequest{}

type DeleteDocumentsRequest struct {
	requests []*DeleteDocumentRequest
}
type _DeleteDocumentsRequest DeleteDocumentsRequest

// DeleteDocumentsRequest represents a batch document delete request.
// The `requests` field is a slice of pointers to DeleteDocumentRequest structs.
// Note: Since pointers are stored instead of copies of the structs, any modification to the data
func NewDeleteDocumentsRequest(requests []*DeleteDocumentRequest) *DeleteDocumentsRequest {
	return &DeleteDocumentsRequest{requests: requests}
}

func NewDeleteDocumentsRequestWithDefaults() *DeleteDocumentsRequest {
	return &DeleteDocumentsRequest{
		requests: make([]*DeleteDocumentRequest, 0),
	}
}

// Add a DeleteDocumentRequest into the request
func (o *DeleteDocumentsRequest) AddDeleteDocument(request *DeleteDocumentRequest) *DeleteDocumentsRequest {
	o.requests = append(o.requests, request)
	return o
}
func (o *DeleteDocumentsRequest) GetRequests() []*DeleteDocumentRequest {
	if o == nil {
		reqs := make([]*DeleteDocumentRequest, 0)
		return reqs
	}
	return o.requests
}
func (o *DeleteDocumentsRequest) SetRequests(reqs []*DeleteDocumentRequest) {
	o.requests = reqs
}

func (o DeleteDocumentsRequest) MarshalJSON() ([]byte, error) {
	if IsNil(o.requests) {
		return nil, nil
	}
	var toSerialize bytes.Buffer
	for _, v := range o.requests {

		m, err := (*v).ToMap()
		if err != nil {
			return nil, err
		}
		obj := map[string]interface{}{
			"delete": m,
		}
		objJson, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		_, err = toSerialize.Write(objJson)
		if err != nil {
			return nil, err
		}
		_, err = toSerialize.Write([]byte("\n"))
		if err != nil {
			return nil, err
		}

	}
	return toSerialize.Bytes(), nil
}
func (o DeleteDocumentsRequest) ToMap() (map[string]interface{}, error) {
	return nil, nil
}
