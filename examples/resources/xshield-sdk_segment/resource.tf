resource "xshield-sdk_segment" "my_segment" {
  auto_synchronize_enabled                 = false
  criteria                                 = "...my_criteria..."
  description                              = "...my_description..."
  lowest_inbound_policy_status             = "...my_lowest_inbound_policy_status..."
  lowest_outbound_policy_status            = "...my_lowest_outbound_policy_status..."
  lowest_progressive_inbound_policy_status = "...my_lowest_progressive_inbound_policy_status..."
  policy_automation_configurable           = false
  tagbasedpolicy_id                        = "...my_tagbasedpolicy_id..."
  tag_based_policy_name                    = "...my_tag_based_policy_name..."
}