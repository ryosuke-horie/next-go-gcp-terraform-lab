provider "google" {
  credentials = file("key.json")
  project     = var.project_id
  region      = var.default_region
}

# Artifact Registry APIを有効化
resource "google_project_service" "artifact_registry_api" {
  project = var.project_id
  service = "artifactregistry.googleapis.com"

  disable_on_destroy = false
}

# Cloud Resource Manager APIを有効化
resource "google_project_service" "cloud_resource_manager_api" {
  project = var.project_id
  service = "cloudresourcemanager.googleapis.com"

  disable_on_destroy = false
}

# 
# Arifact Registry
# 

# Artifact Registry Repositoryの作成
resource "google_artifact_registry_repository" "task-api-golang-repo" {
  location      = var.default_region
  repository_id = "task-api-golang-repo"
  description   = "タスク管理アプリケーションのGolang製APIイメージ格納用レジストリ"
  format        = "docker"
  #   kms_key_name           = "KEY"
  cleanup_policy_dry_run = false # クリーンアップポリシーを適用する
  cleanup_policies {
    id     = "delete-old-images"
    action = "DELETE"
    condition {
      older_than = "2592000s" # 30日を秒に換算
    }
  }

  depends_on = [
    google_project_service.artifact_registry_api,
    google_project_service.cloud_resource_manager_api
  ]
}

# IAMポリシーの設定
resource "google_project_iam_member" "artifact_registry_admin" {
  project = var.project_id
  role    = "roles/artifactregistry.admin"
  member  = "serviceAccount:gcp-terraform-sa@${var.project_id}.iam.gserviceaccount.com"

  depends_on = [
    google_project_service.cloud_resource_manager_api
  ]
}

# 
# CloudSQL
# 

resource "google_sql_database_instance" "next-go-gcp-terraform-k8s-lab-db-instance" {
  name             = "next-go-gcp-terraform-k8s-lab-db-instance"
  database_version = "POSTGRES_17"
  settings {
    edition = "ENTERPRISE" # v16以降は明示的に指定する
    tier    = "db-f1-micro"
  }

  deletion_protection = "true"
}

resource "google_sql_database" "next-go-gcp-terraform-k8s-lab-db" {
  name     = "next-go-gcp-terraform-k8s-lab-db"
  instance = google_sql_database_instance.next-go-gcp-terraform-k8s-lab-db-instance.name
}

resource "google_sql_user" "sql-user" {
  name     = "sql-user"
  instance = google_sql_database_instance.next-go-gcp-terraform-k8s-lab-db-instance.name
  password = var.db_password
}
