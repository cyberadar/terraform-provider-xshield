---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xshield_named_network Data Source - terraform-provider-xshield"
subcategory: ""
description: |-
  NamedNetwork DataSource
---

# xshield_named_network (Data Source)

NamedNetwork DataSource

## Example Usage

```terraform
data "xshield_named_network" "my_namednetwork" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `assigned_by_tag_based_policy` (Boolean)
- `colortokens_managed` (Boolean)
- `id` (String) The ID of this resource.
- `ip_ranges` (Attributes List) (see [below for nested schema](#nestedatt--ip_ranges))
- `named_network_assignments` (Number)
- `named_network_description` (String)
- `named_network_name` (String)
- `namednetwork_tag_based_policy_assignments` (Number)
- `program_as_internet` (Boolean)
- `program_as_intranet` (Boolean)
- `region` (String)
- `service` (String)
- `total_comments` (Number)
- `total_count` (Number)
- `usergroup_named_network_assignments` (Number)

<a id="nestedatt--ip_ranges"></a>
### Nested Schema for `ip_ranges`

Read-Only:

- `id` (String)
- `ip_count` (Number)
- `ip_range` (String)
