variable "project_id" {
  description = "project_id"
  type        = string
  default     = "plasma-renderer-446307-u5"
}

variable "default_region" {
  description = "The default region for resources"
  default     = "asia-southeast1"
}

variable "db_password" {
  description = "Databaseのパスワード"
  type        = string
  sensitive   = true # 機密情報として扱う
}
