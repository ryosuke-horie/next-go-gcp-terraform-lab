output "service_account_email" {
  description = "GitHub Actions用サービスアカウントのメールアドレス"
  value       = google_service_account.github_actions.email
}

output "workload_identity_pool_id" {
  description = "Workload Identity PoolのID"
  value       = google_iam_workload_identity_pool.github_actions.workload_identity_pool_id
}

output "workload_identity_pool_provider_id" {
  description = "Workload Identity Pool ProviderのID"
  value       = google_iam_workload_identity_pool_provider.github_actions.workload_identity_pool_provider_id
}
