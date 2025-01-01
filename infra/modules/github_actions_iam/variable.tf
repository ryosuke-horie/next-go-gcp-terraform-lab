variable "workload_identity_pool_id" {
  description = "Workload Identity PoolのID"
  type        = string
}

variable "workload_identity_pool_provider_id" {
  description = "Workload Identity Pool ProviderのID"
  type        = string
}

variable "project_id" {
  description = "GCPプロジェクトID"
  type        = string
}

variable "account_id" {
  description = "GitHub Actions用サービスアカウントのID"
  type        = string
  default     = "github-actions"
}

variable "repository" {
  description = "GitHub リポジトリ名"
  type        = string
  default     = "ryosuke-horie/next-go-gcp-terraform-k8s-lab"
}
