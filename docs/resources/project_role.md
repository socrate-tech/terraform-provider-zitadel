---
page_title: "zitadel_project_role Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing the project roles, which can be given as authorizations to users.
---

# zitadel_project_role (Resource)

Resource representing the project roles, which can be given as authorizations to users.

## Example Usage

```terraform
resource "zitadel_project_role" "default" {
  org_id       = data.zitadel_org.default.id
  project_id   = data.zitadel_project.default.id
  role_key     = "super-user"
  display_name = "display_name2"
  group        = "role_group"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) Name used for project role
- `project_id` (String) ID of the project
- `role_key` (String) Key used for project role

### Optional

- `group` (String) Group used for project role
- `org_id` (String) ID of the organization

### Read-Only

- `id` (String) The ID of this resource.

## Import

```terraform
# The resource can be imported using the ID format `<project_id:role_key[:org_id]>`, e.g.
terraform import project_role.imported '123456789012345678:my-role-key:123456789012345678'
```
