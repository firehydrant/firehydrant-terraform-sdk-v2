---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant_functionality Data Source - terraform-provider-firehydrant"
subcategory: ""
description: |-
  Functionality DataSource
---

# firehydrant_functionality (Data Source)

Functionality DataSource

## Example Usage

```terraform
data "firehydrant_functionality" "my_functionality" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `active_incidents` (List of String) List of active incident guids
- `alert_on_add` (Boolean)
- `auto_add_responding_team` (Boolean)
- `created_at` (String)
- `description` (String)
- `external_resources` (Attributes List) Information about known linkages to representations of services outside of FireHydrant. (see [below for nested schema](#nestedatt--external_resources))
- `id` (String) The ID of this resource.
- `labels` (Map of String) An object of label key and values
- `links` (Attributes List) List of links attached to this functionality. (see [below for nested schema](#nestedatt--links))
- `name` (String)
- `owner` (Attributes) TeamEntity model (see [below for nested schema](#nestedatt--owner))
- `slug` (String)
- `updated_at` (String)
- `updated_by` (Attributes) (see [below for nested schema](#nestedatt--updated_by))

<a id="nestedatt--external_resources"></a>
### Nested Schema for `external_resources`

Read-Only:

- `connection_id` (String)
- `connection_name` (String)
- `connection_type` (String)
- `created_at` (String)
- `name` (String)
- `remote_id` (String)
- `remote_url` (String)
- `updated_at` (String)


<a id="nestedatt--links"></a>
### Nested Schema for `links`

Read-Only:

- `href_url` (String)
- `icon_url` (String)
- `id` (String)
- `name` (String)


<a id="nestedatt--owner"></a>
### Nested Schema for `owner`


<a id="nestedatt--updated_by"></a>
### Nested Schema for `updated_by`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)
- `source` (String)
