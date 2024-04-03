---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_networkmanager_link Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::NetworkManager::Link
---

# awscc_networkmanager_link (Data Source)

Data Source schema for AWS::NetworkManager::Link



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `bandwidth` (Attributes) The Bandwidth for the link. (see [below for nested schema](#nestedatt--bandwidth))
- `created_at` (String) The date and time that the device was created.
- `description` (String) The description of the link.
- `global_network_id` (String) The ID of the global network.
- `link_arn` (String) The Amazon Resource Name (ARN) of the link.
- `link_id` (String) The ID of the link.
- `provider_name` (String) The provider of the link.
- `site_id` (String) The ID of the site
- `state` (String) The state of the link.
- `tags` (Attributes Set) The tags for the link. (see [below for nested schema](#nestedatt--tags))
- `type` (String) The type of the link.

<a id="nestedatt--bandwidth"></a>
### Nested Schema for `bandwidth`

Read-Only:

- `download_speed` (Number) Download speed in Mbps.
- `upload_speed` (Number) Upload speed in Mbps.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.