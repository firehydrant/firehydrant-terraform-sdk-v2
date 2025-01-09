---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant-terraform-sdk_task_list Resource - terraform-provider-firehydrant-terraform-sdk"
subcategory: ""
description: |-
  TaskList Resource
---

# firehydrant-terraform-sdk_task_list (Resource)

TaskList Resource

## Example Usage

```terraform
resource "firehydrant-terraform-sdk_task_list" "my_tasklist" {
  description = "...my_description..."
  name        = "...my_name..."
  task_list_items = [
    {
      description = "...my_description..."
      summary     = "...my_summary..."
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `task_list_items` (Attributes List) (see [below for nested schema](#nestedatt--task_list_items))

### Optional

- `description` (String)

### Read-Only

- `created_at` (String)
- `email` (String)
- `id` (String) The ID of this resource.
- `source` (String)
- `summary` (String)
- `updated_at` (String)

<a id="nestedatt--task_list_items"></a>
### Nested Schema for `task_list_items`

Required:

- `summary` (String) A summary of the task

Optional:

- `description` (String) A long-form description for the task if additional context is helpful

## Import

Import is supported using the following syntax:

```shell
terraform import firehydrant-terraform-sdk_task_list.my_firehydrant-terraform-sdk_task_list ""
```