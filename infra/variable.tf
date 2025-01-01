variable "project_id" {
  description = "project_id"
  type        = string
  default     = "plasma-renderer-446307-u5"
}

variable "use_key_file" {
  description = "key.jsonを利用するかどうか"
  type        = bool
  default     = true
}

variable "default_region" {
  description = "The default region for resources"
  default     = "asia-southeast1"
}

variable "db_name" {
  description = "データベースの名前"
  type        = string
  default     = "xodb"
}

variable "db_password" {
  description = "Databaseのパスワード"
  type        = string
  sensitive   = true # 機密情報として扱う
}
