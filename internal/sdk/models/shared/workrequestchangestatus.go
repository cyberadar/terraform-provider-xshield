// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type WorkrequestChangeStatus string

const (
	WorkrequestChangeStatusPending    WorkrequestChangeStatus = "Pending"
	WorkrequestChangeStatusInProgress WorkrequestChangeStatus = "InProgress"
	WorkrequestChangeStatusCancelled  WorkrequestChangeStatus = "Cancelled"
	WorkrequestChangeStatusCompleted  WorkrequestChangeStatus = "Completed"
	WorkrequestChangeStatusSuperseded WorkrequestChangeStatus = "Superseded"
	WorkrequestChangeStatusRetry      WorkrequestChangeStatus = "Retry"
)

func (e WorkrequestChangeStatus) ToPointer() *WorkrequestChangeStatus {
	return &e
}
func (e *WorkrequestChangeStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "InProgress":
		fallthrough
	case "Cancelled":
		fallthrough
	case "Completed":
		fallthrough
	case "Superseded":
		fallthrough
	case "Retry":
		*e = WorkrequestChangeStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WorkrequestChangeStatus: %v", v)
	}
}
