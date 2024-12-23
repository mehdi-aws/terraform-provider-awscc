---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_resource_gateway Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Creates a resource gateway for a service.
---

# awscc_vpclattice_resource_gateway (Resource)

Creates a resource gateway for a service.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ip_address_type` (String)
- `name` (String)
- `security_group_ids` (Set of String) The ID of one or more security groups to associate with the endpoint network interface.
- `subnet_ids` (Set of String) The ID of one or more subnets in which to create an endpoint network interface.
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `vpc_identifier` (String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `resource_gateway_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_vpclattice_resource_gateway.example "arn"
```