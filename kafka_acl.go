package aiven

import (
	"encoding/json"
	"fmt"
)

type (
	// KafkaAcl represents a Kafka Topic ACL on Aiven
	KafkaAcl struct {
		Id         string `json:"id"`
		Permission string `json:"permission"`
		Topic      string `json:"topic"`
		Username   string `json:"username"`
	}

	// KafkaAclResponse is the response for listing Kafka ACLs
	KafkaAclResponse struct {
		APIResponse
		Acl KafkaAcl `json:"acl"`
	}

	// CreateKafkaAclRequest are the parameters used to create a Kafka ACL
	CreateKafkaAclRequest struct {
	}

	// KafkaAclHandler is the client which interacts with the kafka endpoints
	// on Aiven.
	KafkaAclHandler struct {
		client *Client
	}
)

// Create creates an ACL on a Kafka Topic
func (h *KafkaAclHandler) Create(project, service string, req CreateKafkaAclRequest) error {
	aclUrl := fmt.Sprintf("/project/%s/service/%s/acl", project, service)
	bts, err := h.client.doPostRequest(aclUrl)
	if err != nil {
		return err
	}

	rsp, err := handleAPIResponse(bts)
	if err != nil {
		return nil, err
	}

	if len(rsp.Errors != 0) {
		return nil, rsp.Errors[0]
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return errors.New(rsp.Message)
	}

	return nil
}
