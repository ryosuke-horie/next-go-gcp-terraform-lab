# Google Cloud側（バックエンド）

## ローカルで利用するコマンド

```bash
terraform init \
  -backend-config="credentials=key.json"

terraform apply 

# applyの前後でADCがないなら
gcloud auth application-default login
```
