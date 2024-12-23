---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudtrail_dashboard Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.
---

# awscc_cloudtrail_dashboard (Resource)

The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String) The name of the dashboard.
- `refresh_schedule` (Attributes) Configures the automatic refresh schedule for the dashboard. Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule. (see [below for nested schema](#nestedatt--refresh_schedule))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `termination_protection_enabled` (Boolean) Indicates whether the dashboard is protected from termination.
- `widgets` (Attributes List) List of widgets on the dashboard (see [below for nested schema](#nestedatt--widgets))

### Read-Only

- `created_timestamp` (String) The timestamp of the dashboard creation.
- `dashboard_arn` (String) The ARN of the dashboard.
- `id` (String) Uniquely identifies the resource.
- `status` (String) The status of the dashboard. Values are CREATING, CREATED, UPDATING, UPDATED and DELETING.
- `type` (String) The type of the dashboard. Values are CUSTOM and MANAGED.
- `updated_timestamp` (String) The timestamp showing when the dashboard was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.

<a id="nestedatt--refresh_schedule"></a>
### Nested Schema for `refresh_schedule`

Optional:

- `frequency` (Attributes) (see [below for nested schema](#nestedatt--refresh_schedule--frequency))
- `status` (String) The status of the schedule. Supported values are ENABLED and DISABLED.
- `time_of_day` (String) StartTime of the automatic schedule refresh.

<a id="nestedatt--refresh_schedule--frequency"></a>
### Nested Schema for `refresh_schedule.frequency`

Optional:

- `unit` (String) The frequency unit. Supported values are HOURS and DAYS.
- `value` (Number) The frequency value.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--widgets"></a>
### Nested Schema for `widgets`

Optional:

- `query_parameters` (List of String) The placeholder keys in the QueryStatement. For example: $StartTime$, $EndTime$, $Period$.
- `query_statement` (String) The SQL query statement on one or more event data stores.
- `view_properties` (Map of String) The view properties of the widget.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cloudtrail_dashboard.example "dashboard_arn"
```