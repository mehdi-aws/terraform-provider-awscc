---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_medialive_cloudwatch_alarm_template Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaLive::CloudWatchAlarmTemplate
---

# awscc_medialive_cloudwatch_alarm_template (Data Source)

Data Source schema for AWS::MediaLive::CloudWatchAlarmTemplate



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) A cloudwatch alarm template's ARN (Amazon Resource Name)
- `cloudwatch_alarm_template_id` (String) A cloudwatch alarm template's id. AWS provided templates have ids that start with `aws-`
- `comparison_operator` (String) The comparison operator used to compare the specified statistic and the threshold.
- `created_at` (String)
- `datapoints_to_alarm` (Number) The number of datapoints within the evaluation period that must be breaching to trigger the alarm.
- `description` (String) A resource's optional description.
- `evaluation_periods` (Number) The number of periods over which data is compared to the specified threshold.
- `group_id` (String) A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`
- `group_identifier` (String) A cloudwatch alarm template group's identifier. Can be either be its id or current name.
- `identifier` (String)
- `metric_name` (String) The name of the metric associated with the alarm. Must be compatible with targetResourceType.
- `modified_at` (String)
- `name` (String) A resource's name. Names must be unique within the scope of a resource type in a specific region.
- `period` (Number) The period, in seconds, over which the specified statistic is applied.
- `statistic` (String) The statistic to apply to the alarm's metric data.
- `tags` (Map of String) Represents the tags associated with a resource.
- `target_resource_type` (String) The resource type this template should dynamically generate cloudwatch metric alarms for.
- `threshold` (Number) The threshold value to compare with the specified statistic.
- `treat_missing_data` (String) Specifies how missing data points are treated when evaluating the alarm's condition.