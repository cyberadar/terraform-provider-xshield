// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/colortokens/terraform-provider-xshield/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *TagRuleResourceModel) ToSharedTagRuleInput() *shared.TagRuleInput {
	onMatch := make(map[string]string)
	for onMatchKey, onMatchValue := range r.OnMatch {
		var onMatchInst string
		onMatchInst = onMatchValue.ValueString()

		onMatch[onMatchKey] = onMatchInst
	}
	var ruleCriteria string
	ruleCriteria = r.RuleCriteria.ValueString()

	ruleDescription := new(string)
	if !r.RuleDescription.IsUnknown() && !r.RuleDescription.IsNull() {
		*ruleDescription = r.RuleDescription.ValueString()
	} else {
		ruleDescription = nil
	}
	ruleEnabled := new(bool)
	if !r.RuleEnabled.IsUnknown() && !r.RuleEnabled.IsNull() {
		*ruleEnabled = r.RuleEnabled.ValueBool()
	} else {
		ruleEnabled = nil
	}
	ruleName := new(string)
	if !r.RuleName.IsUnknown() && !r.RuleName.IsNull() {
		*ruleName = r.RuleName.ValueString()
	} else {
		ruleName = nil
	}
	out := shared.TagRuleInput{
		OnMatch:         onMatch,
		RuleCriteria:    ruleCriteria,
		RuleDescription: ruleDescription,
		RuleEnabled:     ruleEnabled,
		RuleName:        ruleName,
	}
	return &out
}

func (r *TagRuleResourceModel) RefreshFromSharedTagRule(resp *shared.TagRule) {
	if resp != nil {
		r.ID = types.StringPointerValue(resp.ID)
		r.MatchingAssets = types.Int64PointerValue(resp.MatchingAssets)
		if len(resp.OnMatch) > 0 {
			r.OnMatch = make(map[string]types.String)
			for key, value := range resp.OnMatch {
				r.OnMatch[key] = types.StringValue(value)
			}
		}
		r.RuleCriteria = types.StringValue(resp.RuleCriteria)
		r.RuleDescription = types.StringPointerValue(resp.RuleDescription)
		r.RuleEnabled = types.BoolPointerValue(resp.RuleEnabled)
		r.RuleName = types.StringPointerValue(resp.RuleName)
	}
}

func (r *TagRuleResourceModel) ToSharedTagRule() *shared.TagRule {
	matchingAssets := new(int64)
	if !r.MatchingAssets.IsUnknown() && !r.MatchingAssets.IsNull() {
		*matchingAssets = r.MatchingAssets.ValueInt64()
	} else {
		matchingAssets = nil
	}
	onMatch := make(map[string]string)
	for onMatchKey, onMatchValue := range r.OnMatch {
		var onMatchInst string
		onMatchInst = onMatchValue.ValueString()

		onMatch[onMatchKey] = onMatchInst
	}
	var ruleCriteria string
	ruleCriteria = r.RuleCriteria.ValueString()

	ruleDescription := new(string)
	if !r.RuleDescription.IsUnknown() && !r.RuleDescription.IsNull() {
		*ruleDescription = r.RuleDescription.ValueString()
	} else {
		ruleDescription = nil
	}
	ruleEnabled := new(bool)
	if !r.RuleEnabled.IsUnknown() && !r.RuleEnabled.IsNull() {
		*ruleEnabled = r.RuleEnabled.ValueBool()
	} else {
		ruleEnabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	ruleName := new(string)
	if !r.RuleName.IsUnknown() && !r.RuleName.IsNull() {
		*ruleName = r.RuleName.ValueString()
	} else {
		ruleName = nil
	}
	out := shared.TagRule{
		MatchingAssets:  matchingAssets,
		OnMatch:         onMatch,
		RuleCriteria:    ruleCriteria,
		RuleDescription: ruleDescription,
		RuleEnabled:     ruleEnabled,
		ID:              id,
		RuleName:        ruleName,
	}
	return &out
}
