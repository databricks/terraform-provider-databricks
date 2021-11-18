---
page_title: "Migration from 0.3.x to 0.4.x"
---
# Migration from 0.3.x to 0.4.x

Certain resources undergone changes in order to improve long-term maintainability. You can upgrade provider with `terraform init -upgrade`. If you're currently using v0.2.x of provider, please first complete the rewrites specified in [0.2.x to 0.3.x](migration-0.3.x.md) guide.

## provider

* Remove `azure_use_pat_for_spn`, `azure_use_pat_for_cli` attributes