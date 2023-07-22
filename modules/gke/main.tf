resource "google_container_cluster" "primary" {
  name     = var.cluster_name
  location = "us-central1-a"

  
  initial_node_count       = 1
  node_config {
    preemptible  = true
    machine_type = "g1-small"
    disk_size_gb = 10
    
    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}