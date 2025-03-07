---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xshield_tag_rule Resource - terraform-provider-xshield"
subcategory: ""
description: |-
  TagRule Resource
---

# xshield_tag_rule (Resource)

TagRule Resource

## Example Usage

```terraform
resource "xshield_tag_rule" "my_tagrule" {
  on_match = {
    key = "value"
  }
  rule_criteria    = "...my_rule_criteria..."
  rule_description = "...my_rule_description..."
  rule_enabled     = true
  rule_name        = "...my_rule_name..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `rule_criteria` (String)

### Optional

- `on_match` (Map of String)
- `rule_description` (String)
- `rule_enabled` (Boolean)
- `rule_name` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `matching_assets` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import xshield_tag_rule.my_xshield_tag_rule ""
```
