# 必要なプロジェクトサービスを有効化
resource "google_project_service" "iam_api" {
  service = "iam.googleapis.com"
  project = var.project_id
}

resource "google_project_service" "iamcredentials" {
  service = "iamcredentials.googleapis.com"
  project = var.project_id

  depends_on = [google_project_service.iam_api]
}

# Workload Identity Poolの作成
resource "google_iam_workload_identity_pool" "github_actions" {
  workload_identity_pool_id = var.workload_identity_pool_id
  display_name              = "GitHub Actions Pool"
  description               = "Workload Identity Pool for GitHub Actions"

  project = var.project_id
}

# Workload Identity Pool Providerの作成
resource "google_iam_workload_identity_pool_provider" "github_actions" {
  workload_identity_pool_provider_id = var.workload_identity_pool_provider_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_actions.workload_identity_pool_id

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }

  # OIDCトークンから取り込みたいclaimsとTerraform上でのキーをマッピング
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
    "attribute.ref"        = "assertion.ref"
  }

  # ここでGitHub側のclaimを用いた条件式を記述
  attribute_condition = "attribute.repository == \"ryosuke-horie/next-go-gcp-terraform-k8s-lab\" && attribute.ref == \"refs/heads/main\""
}


# GitHub Actions用のサービスアカウントを作成
resource "google_service_account" "github_actions" {
  account_id   = var.account_id
  display_name = "GitHub Actions Service Account"
  project      = var.project_id

  depends_on = [google_project_service.iamcredentials]
}

# Workload Identity Federationとのバインディング
resource "google_service_account_iam_binding" "github_actions_binding" {
  service_account_id = google_service_account.github_actions.name
  role               = "roles/iam.workloadIdentityUser"

  members = [
    "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_actions.name}/attribute.repository/${var.repository}",
  ]

  depends_on = [
    google_iam_workload_identity_pool_provider.github_actions,
    google_service_account.github_actions
  ]
}


# 必要なIAMロールを付与
resource "google_project_iam_member" "run_admin" {
  project = var.project_id
  role    = "roles/run.admin"
  member  = "serviceAccount:${google_service_account.github_actions.email}"
}

resource "google_project_iam_member" "artifactregistry_admin" {
  project = var.project_id
  role    = "roles/artifactregistry.admin"
  member  = "serviceAccount:${google_service_account.github_actions.email}"
}

resource "google_project_iam_member" "storage_admin" {
  project = var.project_id
  role    = "roles/storage.admin"
  member  = "serviceAccount:${google_service_account.github_actions.email}"
}

resource "google_project_iam_member" "cloudsql_admin" {
  project = var.project_id
  role    = "roles/cloudsql.admin"
  member  = "serviceAccount:${google_service_account.github_actions.email}"
}

resource "google_project_iam_member" "iam_service_account_user" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.github_actions.email}"
}
