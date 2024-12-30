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
