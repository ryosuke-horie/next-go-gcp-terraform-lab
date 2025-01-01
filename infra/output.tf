output "github_actions_service_account_email" {
  description = "GitHub Actions用サービスアカウントのメールアドレス"
  value       = module.github_actions_iam.service_account_email
}
