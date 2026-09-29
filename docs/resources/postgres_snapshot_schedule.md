---
subcategory: "Postgres"
---
# databricks_postgres_snapshot_schedule Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A snapshot schedule defines the cadences at which Lakebase automatically takes snapshots of a branch. There is exactly one snapshot schedule per branch (a singleton): it always exists — a branch with no cadences has an empty schedule — and is identified by the branch's resource name. You address it with `parent` (the branch) rather than a separate ID; Terraform manages the singleton in place, so applying the resource sets the branch's cadences.

Updating the resource replaces the entire set of cadences; an empty set disables automatic snapshots for the branch.

### Hierarchy Context

Snapshot schedules exist within the Lakebase Autoscaling resource hierarchy:
- A **snapshot schedule** belongs to a **branch** within a **project**
- Each branch has exactly one snapshot schedule

### Configuration

`schedule` is the set of cadences at which snapshots are taken. Each cadence sets exactly one recurrence pattern plus a `retention`:
- `daily_schedule` — a snapshot once per day, at `hour` (UTC, `0`–`23`).
- `weekly_schedule` — a snapshot once per week, on `day_of_week` (`MONDAY`–`SUNDAY`) at `hour` (UTC, `0`–`23`).
- `monthly_schedule` — a snapshot once per month, on `day` (`1`–`31`) at `hour` (UTC, `0`–`23`). In shorter months the snapshot is taken on the last day instead.

`retention` is how long snapshots from that cadence are kept before automatic deletion; it must be at least 1 hour, written as a duration string in hours/minutes/seconds — for example `168h0m0s` (7 days). When several cadences fire at the same time, one snapshot is taken and kept for the longest of their retentions.


## Example Usage
### Managing a Branch's Snapshot Schedule

A branch's snapshot schedule is a singleton addressed by the branch it belongs
to. Set `parent` to the branch's resource name and provide the desired cadences;
Terraform applies them in place. The example below takes a daily snapshot at
03:00 UTC and keeps it for 7 days.

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_snapshot_schedule" "this" {
  parent = "${databricks_postgres_project.this.name}/branches/production"
  schedule = [
    {
      daily_schedule = {
        hour = 3
      }
      retention = "168h0m0s" # keep daily snapshots for 7 days
    }
  ]
}
```

The `schedule` set can hold more than one cadence — for example, add a weekly
cadence alongside the daily one to keep some snapshots longer than others.

To disable automatic snapshots, set `schedule = []` and apply. Removing the
resource from your configuration only removes it from Terraform state; it does
not change the schedule on the branch.


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the parent
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