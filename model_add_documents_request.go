package openapi

import (
	"bytes"
)

var _ MappedNullable = &AddDocumentsRequest{}

type AddDocumentsRequest struct {
	requests []*InsertDocumentRequest
}
type _AddDocumentsRequest AddDocumentsRequest

// AddDocumentsRequest represents a batch document insert request.
// The `requests` field is a slice of pointers to InsertDocumentRequest structs.
// Note: Since pointers are stored instead of copies of the structs, any modification to the data
func NewAddDocumentsRequest(requests []*InsertDocumentRequest) *AddDocumentsRequest {
	return &AddDocumentsRequest{requests: requests}
}

func NewAddDocumentsRequestWithDefaults() *AddDocumentsRequest {
	return &AddDocumentsRequest{
		requests: make([]*InsertDocumentRequest, 0),
	}
}

// Add a InsertDocumentRequest into the request
func (o *AddDocumentsRequest) AddInsertDocument(request *InsertDocumentRequest) *AddDocumentsRequest {
	o.requests = append(o.requests, request)
	return o
}
func (o *AddDocumentsRequest) GetRequests() []*InsertDocumentRequest {
	if o == nil {
		reqs := make([]*InsertDocumentRequest, 0)
		return reqs
	}
	return o.requests
}
func (o *AddDocumentsRequest) SetRequests(reqs []*InsertDocumentRequest) {
	o.requests = reqs
}

func (o AddDocumentsRequest) MarshalJSON() ([]byte, error) {
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
			"insert": m,
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
func (o AddDocumentsRequest) ToMap() (map[string]interface{}, error) {
	return nil, nil
}
