resource "xshield_named_network" "my_namednetwork" {
  ip_ranges = [
    {
      ip_range = "...my_ip_range..."
    }
  ]
  named_network_description = "...my_named_network_description..."
  named_network_name        = "...my_named_network_name..."
  program_as_intranet       = false
  region                    = "...my_region..."
  service                   = "...my_service..."
}