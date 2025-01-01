provider "google" {
  project = var.project_id
  region  = var.default_region
}

terraform {
  backend "gcs" {
    bucket      = "gs-state-terraform-plasma-renderer-446307-u5"
    prefix      = "terraform/state"
    credentials = var.use_key_file ? file("key.json") : null
  }
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

  deletion_protection = false
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

# 
# Cloud Run
# 
resource "google_cloud_run_v2_service" "default" {
  name                = "cloudrun-service"
  location            = var.default_region
  deletion_protection = false
  ingress             = "INGRESS_TRAFFIC_ALL"

  template {
    scaling {
      max_instance_count = 1
    }

    volumes {
      name = "cloudsql"
      cloud_sql_instance {
        instances = [google_sql_database_instance.next-go-gcp-terraform-k8s-lab-db-instance.connection_name]
      }
    }

    containers {
      image = "asia-southeast1-docker.pkg.dev/plasma-renderer-446307-u5/task-api-repositry/gotodo:latest"

      resources {
        limits = {
          cpu    = "1"
          memory = "2Gi"
        }
      }

      ports {
        container_port = 8080
      }

      volume_mounts {
        name       = "cloudsql"
        mount_path = "/cloudsql"
      }

      env {
        name  = "DB_USER"
        value = google_sql_user.sql-user.name
      }

      env {
        name  = "DB_PASSWORD"
        value = var.db_password
      }

      env {
        name  = "DB_NAME"
        value = google_sql_database.next-go-gcp-terraform-k8s-lab-db.name
      }

      # DB_HOST は Cloud SQL の Unix ソケットパス
      env {
        name  = "DB_HOST"
        value = "/cloudsql/plasma-renderer-446307-u5:asia-southeast1:next-go-gcp-terraform-k8s-lab-db-instance"
      }
    }
  }

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }
}

# Cloud SQL Client ロールの付与
resource "google_project_iam_member" "cloud_run_cloudsql_client" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:gcp-terraform-sa@${var.project_id}.iam.gserviceaccount.com"

  depends_on = [
    google_project_service.cloud_resource_manager_api
  ]
}

# Cloud Run サービスへのパブリックアクセスを許可
resource "google_cloud_run_service_iam_member" "allow_public" {
  project  = var.project_id
  location = google_cloud_run_v2_service.default.location
  service  = google_cloud_run_v2_service.default.name

  role   = "roles/run.invoker"
  member = "allUsers"

  depends_on = [
    google_cloud_run_v2_service.default
  ]
}
