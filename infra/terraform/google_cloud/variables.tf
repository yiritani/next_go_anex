variable "project_id" {
  type        = string
}

variable "region" {
  description = "GCPのリージョン"
  type        = string
  default     = "us-central1"
}

variable "dockerfile_backend" {
  description = "バックエンドのDockerfileパス"
  type        = string
  default     = "apps/backend/Dockerfile"
}

variable "dockerfile_frontend" {
  description = "フロントエンドのDockerfileパス"
  type        = string
  default     = "apps/frontend/Dockerfile"
}

variable "dockerfile_job" {
  description = "jobのDockerfileパス"
  type        = string
  default     = "apps/job/Dockerfile"
}

variable "dockerfile_backend_grpc" {
  description = "バックエンド_grpcのDockerfileパス"
  type        = string
  default     = "apps/backend_grpc/Dockerfile"
}

variable "dockerfile_frontend_grpc" {
  description = "フロントエンド_grpcのDockerfileパス"
  type        = string
  default     = "apps/frontend_grpc/Dockerfile"
}

variable "dockerfile_migration" {
  description = "migrationのDockerfileパス"
  type        = string
  default     = "apps/backend/Dockerfile.migration"
}

variable "service_name" {
  description = "サービス名"
  type        = string
  default     = "next-go"
}

variable "image_repo" {
  description = "イメージリポジトリ名"
  type        = string
  default     = "next-go"
}

variable "vpc_network" {
  description = "The VPC network name"
  type        = string
  default     = "next-go-vpc"
}

variable "domain_name" {
  description = "The domain name for Cloud DNS"
  type        = string
  default     = "next-go.com"
}

variable "managed_zone" {
  description = "The name of the Cloud DNS managed zone"
  type        = string
  default     = "next-go-zone"
}

variable "support_email" {
  description = "The support email for IAP"
  type        = string
  default     = "yusuke.iritani0909@gmail.com"
}

variable "access_user_email" {
  description = "The email address of the user to grant access to"
  type        = string
  default     = "yusuke.iritani0909@gmail.com"
}

variable "bucket_name" {
    description = "The name of the GCS bucket"
    type        = string
    default     = "next-go-bucket"
}