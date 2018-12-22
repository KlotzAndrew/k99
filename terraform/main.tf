data "google_compute_network" "gke_network" {
  project = "${var.project_name}"
  name    = "my-gke-network"
}

resource "google_container_cluster" "gke-cluster" {
  name               = "k99-gke-cluster"
  network            = "default"
  zone               = "${var.zone}"
  initial_node_count = 0
}

resource "google_container_node_pool" "primary-pool" {
  name               = "primary-node-pool"
  zone               = "${var.zone}"
  cluster            = "${google_container_cluster.gke-cluster.name}"
  initial_node_count = 3

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"
  }

  autoscaling {
    min_node_count = "${var.min_node_count}"
    max_node_count = "${var.max_node_count}"
  }

  management {
    auto_repair = true
    auto_upgrade = true
  }
}
