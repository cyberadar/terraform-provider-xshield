// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type WorkrequestWorkItemAction string

const (
	WorkrequestWorkItemActionAssetAdded                              WorkrequestWorkItemAction = "AssetAdded"
	WorkrequestWorkItemActionAssetUpdated                            WorkrequestWorkItemAction = "AssetUpdated"
	WorkrequestWorkItemActionTagRuleAdded                            WorkrequestWorkItemAction = "TagRuleAdded"
	WorkrequestWorkItemActionTagRuleModified                         WorkrequestWorkItemAction = "TagRuleModified"
	WorkrequestWorkItemActionTagBasedPolicyAdded                     WorkrequestWorkItemAction = "TagBasedPolicyAdded"
	WorkrequestWorkItemActionTagBasedPolicyAssignmentsModified       WorkrequestWorkItemAction = "TagBasedPolicyAssignmentsModified"
	WorkrequestWorkItemActionTagBasedPolicyDeleted                   WorkrequestWorkItemAction = "TagBasedPolicyDeleted"
	WorkrequestWorkItemActionTagUpdated                              WorkrequestWorkItemAction = "TagUpdated"
	WorkrequestWorkItemActionTagDeleted                              WorkrequestWorkItemAction = "TagDeleted"
	WorkrequestWorkItemActionNamedNetworkAssigned                    WorkrequestWorkItemAction = "NamedNetworkAssigned"
	WorkrequestWorkItemActionNamedNetworkUnassigned                  WorkrequestWorkItemAction = "NamedNetworkUnassigned"
	WorkrequestWorkItemActionNamedNetworkDeleted                     WorkrequestWorkItemAction = "NamedNetworkDeleted"
	WorkrequestWorkItemActionNamedNetworkRangeAdd                    WorkrequestWorkItemAction = "NamedNetworkRangeAdd"
	WorkrequestWorkItemActionNamedNetworkRangeRemove                 WorkrequestWorkItemAction = "NamedNetworkRangeRemove"
	WorkrequestWorkItemActionNamedNetworkProgramAsIntranet           WorkrequestWorkItemAction = "NamedNetworkProgramAsIntranet"
	WorkrequestWorkItemActionScimPullData                            WorkrequestWorkItemAction = "ScimPullData"
	WorkrequestWorkItemActionTagBasedPolicyMemberChange              WorkrequestWorkItemAction = "TagBasedPolicyMemberChange"
	WorkrequestWorkItemActionEvaluateTemplates                       WorkrequestWorkItemAction = "EvaluateTemplates"
	WorkrequestWorkItemActionTemplateEdit                            WorkrequestWorkItemAction = "TemplateEdit"
	WorkrequestWorkItemActionUpdateNetworkMap                        WorkrequestWorkItemAction = "UpdateNetworkMap"
	WorkrequestWorkItemActionEnableUserGroups                        WorkrequestWorkItemAction = "EnableUserGroups"
	WorkrequestWorkItemActionDisableUserGroups                       WorkrequestWorkItemAction = "DisableUserGroups"
	WorkrequestWorkItemActionAgentDeleted                            WorkrequestWorkItemAction = "AgentDeleted"
	WorkrequestWorkItemActionTagBasedPolicyProgressiveInboundRefresh WorkrequestWorkItemAction = "TagBasedPolicyProgressiveInboundRefresh"
	WorkrequestWorkItemActionEnableCrowdstrikeHostGroup              WorkrequestWorkItemAction = "EnableCrowdstrikeHostGroup"
	WorkrequestWorkItemActionDisableCrowdstrikeHostGroup             WorkrequestWorkItemAction = "DisableCrowdstrikeHostGroup"
)

func (e WorkrequestWorkItemAction) ToPointer() *WorkrequestWorkItemAction {
	return &e
}
func (e *WorkrequestWorkItemAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AssetAdded":
		fallthrough
	case "AssetUpdated":
		fallthrough
	case "TagRuleAdded":
		fallthrough
	case "TagRuleModified":
		fallthrough
	case "TagBasedPolicyAdded":
		fallthrough
	case "TagBasedPolicyAssignmentsModified":
		fallthrough
	case "TagBasedPolicyDeleted":
		fallthrough
	case "TagUpdated":
		fallthrough
	case "TagDeleted":
		fallthrough
	case "NamedNetworkAssigned":
		fallthrough
	case "NamedNetworkUnassigned":
		fallthrough
	case "NamedNetworkDeleted":
		fallthrough
	case "NamedNetworkRangeAdd":
		fallthrough
	case "NamedNetworkRangeRemove":
		fallthrough
	case "NamedNetworkProgramAsIntranet":
		fallthrough
	case "ScimPullData":
		fallthrough
	case "TagBasedPolicyMemberChange":
		fallthrough
	case "EvaluateTemplates":
		fallthrough
	case "TemplateEdit":
		fallthrough
	case "UpdateNetworkMap":
		fallthrough
	case "EnableUserGroups":
		fallthrough
	case "DisableUserGroups":
		fallthrough
	case "AgentDeleted":
		fallthrough
	case "TagBasedPolicyProgressiveInboundRefresh":
		fallthrough
	case "EnableCrowdstrikeHostGroup":
		fallthrough
	case "DisableCrowdstrikeHostGroup":
		*e = WorkrequestWorkItemAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WorkrequestWorkItemAction: %v", v)
	}
}
