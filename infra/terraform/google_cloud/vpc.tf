# resource "google_project_service" "vpc_api" {
#   service = "vpcaccess.googleapis.com"
#   project = var.project_id
# }
#
# # VPCネットワーク作成
# resource "google_compute_network" "vpc_network" {
#   depends_on = [google_project_service.vpc_api]
#
#   name                    = "${var.service_name}-vpc-network"
#   auto_create_subnetworks = false
# }
#
# # サブネット作成
# resource "google_compute_subnetwork" "vpc_subnet" {
#   depends_on = [google_project_service.vpc_api]
#
#   name          = "${var.service_name}-vpc-subnet"
#   ip_cidr_range = "10.0.0.0/24"
#   region        = var.region
#   network       = google_compute_network.vpc_network.id
# }
#
# # VPCアクセスコネクタ作成
# resource "google_vpc_access_connector" "vpc_connector" {
#   depends_on = [google_project_service.vpc_api]
#
#   name   = "${var.service_name}-vpc-connector"
#   region = var.region
#   network = google_compute_network.vpc_network.id
#
#   ip_cidr_range = "10.8.0.0/28"
# }