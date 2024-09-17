resource "xshield_segment" "my_segment" {
  criteria           = "...my_criteria..."
  criteria_as_params = "...my_criteria_as_params..."
  description        = "...my_description..."
  namednetworks = [
    {
      named_network_id   = "...my_named_network_id..."
      named_network_name = "...my_named_network_name..."
    }
  ]
  tag_based_policy_name      = "...my_tag_based_policy_name..."
  target_breach_impact_score = 61
  templates = [
    {
      template_id   = "...my_template_id..."
      template_name = "...my_template_name..."
    }
  ]
  timeline = 10
}