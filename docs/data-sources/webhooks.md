---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant-terraform-sdk_webhooks Data Source - terraform-provider-firehydrant-terraform-sdk"
subcategory: ""
description: |-
  Webhooks DataSource
---

# firehydrant-terraform-sdk_webhooks (Data Source)

Webhooks DataSource

## Example Usage

```terraform
data "firehydrant-terraform-sdk_webhooks" "my_webhooks" {
  page     = 2
  per_page = 9
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `page` (Number)
- `per_page` (Number)

### Read-Only

- `created_at` (String)
- `email` (String)
- `id` (String) The ID of this resource.
- `name` (String)
- `source` (String)
- `state` (String)
- `subscriptions` (String)
- `updated_at` (String)
- `url` (String)