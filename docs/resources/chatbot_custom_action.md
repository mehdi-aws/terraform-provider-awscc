---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_chatbot_custom_action Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Chatbot::CustomAction Resource Type
---

# awscc_chatbot_custom_action (Resource)

Definition of AWS::Chatbot::CustomAction Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `action_name` (String)
- `definition` (Attributes) (see [below for nested schema](#nestedatt--definition))

### Optional

- `alias_name` (String)
- `attachments` (Attributes List) (see [below for nested schema](#nestedatt--attachments))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `custom_action_arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--definition"></a>
### Nested Schema for `definition`

Required:

- `command_text` (String)


<a id="nestedatt--attachments"></a>
### Nested Schema for `attachments`

Optional:

- `button_text` (String)
- `criteria` (Attributes List) (see [below for nested schema](#nestedatt--attachments--criteria))
- `notification_type` (String)
- `variables` (Map of String)

<a id="nestedatt--attachments--criteria"></a>
### Nested Schema for `attachments.criteria`

Optional:

- `operator` (String)
- `value` (String)
- `variable_name` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_chatbot_custom_action.example "custom_action_arn"
```
