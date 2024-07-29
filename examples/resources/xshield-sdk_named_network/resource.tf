resource "xshield-sdk_named_network" "my_namednetwork" {
  named_network_description = "...my_named_network_description..."
  named_network_id          = "...my_named_network_id..."
  named_network_name        = "...my_named_network_name..."
  program_as_intranet       = true
  region                    = "...my_region..."
  service                   = "...my_service..."
}