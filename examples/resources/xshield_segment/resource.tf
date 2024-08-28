resource "xshield_segment" "my_segment" {
  criteria                   = "...my_criteria..."
  criteria_as_params         = "...my_criteria_as_params..."
  description                = "...my_description..."
  tagbasedpolicy_id          = "...my_tagbasedpolicy_id..."
  tag_based_policy_name      = "...my_tag_based_policy_name..."
  target_breach_impact_score = 3
  timeline                   = 9
}