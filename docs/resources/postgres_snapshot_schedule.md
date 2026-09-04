---
subcategory: "Postgres"
---
# databricks_postgres_snapshot_schedule Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)



## Example Usage


## Arguments
The following arguments are supported:
* `schedule` (list of ScheduleCadence, optional) - The cadences at which automatic snapshots are taken. Update replaces the
  whole set; an empty set disables automatic snapshots. Order is not
  significant. When several cadences fire together, one snapshot is taken,
  retained for the longest of their retentions
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### DailySchedule
* `hour` (integer, optional) - The hour of the day, in UTC, at which to take the snapshot, in [0, 23]

### MonthlySchedule
* `day` (integer, required) - The day of the month on which to take the snapshot, in [1, 31]. In shorter
  months the snapshot is taken on the last day instead (day 31 runs on Feb 28
  or 29, and on Apr 30), so every month gets exactly one snapshot
* `hour` (integer, optional) - The hour of the day, in UTC, at which to take the snapshot, in [0, 23]

### ScheduleCadence
* `retention` (string, required) - How long snapshots from this cadence are kept before automatic deletion.
  Must be at least 1 hour. Applied when a snapshot is taken; not retroactive,
  so changing it affects only later snapshots
* `daily_schedule` (DailySchedule, optional) - Take a snapshot once per day
* `monthly_schedule` (MonthlySchedule, optional) - Take a snapshot once per month
* `weekly_schedule` (WeeklySchedule, optional) - Take a snapshot once per week

### WeeklySchedule
* `day_of_week` (string, required) - The day of the week on which to take the snapshot. Possible values are: `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`
* `hour` (integer, optional) - The hour of the day, in UTC, at which to take the snapshot, in [0, 23]

## Attributes
In addition to the above arguments, the following attributes are exported:
* `name` (string) - The resource name of the branch's snapshot schedule.
  Format: projects/{project_id}/branches/{branch_id}/snapshot-schedule

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_snapshot_schedule.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_snapshot_schedule.this "name"
```