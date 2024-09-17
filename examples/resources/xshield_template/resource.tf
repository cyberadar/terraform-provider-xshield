resource "xshield_template" "my_template" {
  template_category    = "...my_template_category..."
  template_description = "...my_template_description..."
  template_name        = "...my_template_name..."
  template_paths = [
    {
      destination_asset_id = "...my_destination_asset_id..."
      destination_named_network = {
        named_network_id   = "...my_named_network_id..."
        named_network_name = "...my_named_network_name..."
      }
      destination_tag_based_policy = {
        criteria              = "...my_criteria..."
        criteria_as_params    = "...my_criteria_as_params..."
        tag_based_policy_id   = "...my_tag_based_policy_id..."
        tag_based_policy_name = "...my_tag_based_policy_name..."
      }
      direction = "...my_direction..."
      domain    = "...my_domain..."
      dst_ip = [
        "..."
      ]
      dst_process     = "...my_dst_process..."
      method          = "...my_method..."
      port            = "...my_port..."
      port_name       = "...my_port_name..."
      protocol        = "...my_protocol..."
      source_asset_id = "...my_source_asset_id..."
      source_named_network = {
        named_network_id   = "...my_named_network_id..."
        named_network_name = "...my_named_network_name..."
      }
      source_tag_based_policy = {
        criteria              = "...my_criteria..."
        criteria_as_params    = "...my_criteria_as_params..."
        tag_based_policy_id   = "...my_tag_based_policy_id..."
        tag_based_policy_name = "...my_tag_based_policy_name..."
      }
      src_ip      = "...my_src_ip..."
      src_process = "...my_src_process..."
      uri         = "...my_uri..."
    }
  ]
  template_ports = [
    {
      listen_port          = "...my_listen_port..."
      listen_port_name     = "...my_listen_port_name..."
      listen_port_protocol = "...my_listen_port_protocol..."
      listen_port_reviewed = "allow-any"
      listen_process_names = [
        "..."
      ]
    }
  ]
  template_type = "application-template"
}