resource "xshield-sdk_segment" "my_segment" {
  auto_synchronize_enabled                 = true
  criteria                                 = "...my_criteria..."
  criteria_as_params                       = "...my_criteria_as_params..."
  description                              = "...my_description..."
  id                                       = "f8974958-c4b6-4069-95c2-7ba88febfa1d"
  lowest_inbound_policy_status             = "...my_lowest_inbound_policy_status..."
  lowest_outbound_policy_status            = "...my_lowest_outbound_policy_status..."
  lowest_progressive_inbound_policy_status = "...my_lowest_progressive_inbound_policy_status..."
  matching_assets                          = 1
  policy_automation_configurable           = false
  policy_progressive_last_refreshed        = "2021-09-22T17:14:52.971Z"
  tagbasedpolicy_id                        = "...my_tagbasedpolicy_id..."
  tag_based_policy_name                    = "...my_tag_based_policy_name..."
}