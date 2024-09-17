resource "xshield_tag_rule" "my_tagrule" {
  on_match = {
    "see" : "documentation",
  }
  rule_criteria           = "...my_rule_criteria..."
  rule_criteria_as_params = "...my_rule_criteria_as_params..."
  rule_description        = "...my_rule_description..."
  rule_enabled            = false
  rule_name               = "...my_rule_name..."
}