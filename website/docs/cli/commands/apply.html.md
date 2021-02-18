---
layout: "docs"
page_title: "Command: apply"
sidebar_current: "docs-commands-apply"
description: |-
  The `terraform apply` command is used to apply the changes required to reach the desired state of the configuration, or the pre-determined set of actions generated by a `terraform plan` execution plan.
---

# Command: apply

> **Hands-on:** Try the [Terraform: Get Started](https://learn.hashicorp.com/collections/terraform/aws-get-started?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) collection on HashiCorp Learn.

The `terraform apply` command is used to apply the changes required
to reach the desired state of the configuration, or the pre-determined
set of actions generated by a `terraform plan` execution plan.

## Usage

Usage: `terraform apply [options] [plan]`

By default, `apply` scans the current directory for the configuration
and applies the changes appropriately. However, you can optionally give the
path to a saved plan file that was previously created with
[`terraform plan`](plan.html).

If you don't give a plan file on the command line, `terraform apply` will
create a new plan automatically and then prompt for approval to apply it. If the
created plan does not include any changes to resources or to root module
output values then `terraform apply` will exit immediately, without prompting.

The command-line flags are all optional. The list of available flags are:

* `-backup=path` - Path to the backup file. Defaults to `-state-out` with
  the ".backup" extension. Disabled by setting to "-".

* `-compact-warnings` - If Terraform produces any warnings that are not
  accompanied by errors, show them in a more compact form that includes only
  the summary messages.

* `-lock=true` - Lock the state file when locking is supported.

* `-lock-timeout=0s` - Duration to retry a state lock.

* `-input=true` - Ask for input for variables if not directly set.

* `-auto-approve` - Skip interactive approval of plan before applying.

* `-no-color` - Disables output with coloring.

* `-parallelism=n` - Limit the number of concurrent operation as Terraform
  [walks the graph](/docs/internals/graph.html#walking-the-graph). Defaults to
  10.

* `-refresh=true` - Update the state for each resource prior to planning
  and applying. This has no effect if a plan file is given directly to
  apply.

* `-state=path` - Path to the state file. Defaults to "terraform.tfstate".
  Ignored when [remote state](/docs/language/state/remote.html) is used. This setting
  does not persist and other commands, such as init, may not be aware of the
  alternate statefile. To configure an alternate statefile path which is
  available to all terraform commands, use the [local backend](/docs/language/settings/backends/local.html).

* `-state-out=path` - Path to write updated state file. By default, the
  `-state` path will be used. Ignored when
  [remote state](/docs/language/state/remote.html) is used.

* `-target=resource` - A [Resource
  Address](/docs/cli/state/resource-addressing.html) to target. For more
  information, see
  [the targeting docs from `terraform plan`](/docs/cli/commands/plan.html#resource-targeting).

* `-var 'foo=bar'` - Set a variable in the Terraform configuration. This flag
  can be set multiple times. Variable values are interpreted as
  [literal expressions](/docs/language/expressions/types.html) in the
  Terraform language, so list and map values can be specified via this flag.

* `-var-file=foo` - Set variables in the Terraform configuration from
  a [variable file](/docs/language/values/variables.html#variable-definitions-tfvars-files). If
  a `terraform.tfvars` or any `.auto.tfvars` files are present in the current
  directory, they will be automatically loaded. `terraform.tfvars` is loaded
  first and the `.auto.tfvars` files after in alphabetical order. Any files
  specified by `-var-file` override any values set automatically from files in
  the working directory. This flag can be used multiple times.

## Passing a Different Configuration Directory

Terraform v0.13 and earlier also accepted a directory path in place of the
plan file argument to `terraform apply`, in which case Terraform would use
that directory as the root module instead of the current working directory.

That usage is still supported in Terraform v0.14, but is now deprecated and we
plan to remove it in Terraform v0.15. If your workflow relies on overriding
the root module directory, use
[the `-chdir` global option](./#switching-working-directory-with-chdir)
instead, which works across all commands and makes Terraform consistently look
in the given directory for all files it would normaly read or write in the
current working directory.

If your previous use of this legacy pattern was also relying on Terraform
writing the `.terraform` subdirectory into the current working directory even
though the root module directory was overridden, use
[the `TF_DATA_DIR` environment variable](/docs/cli/config/environment-variables.html#tf_data_dir)
to direct Terraform to write the `.terraform` directory to a location other
than the current working directory.