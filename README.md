# hcl2fmt

Recursively format HCL2 files.

## Usage

`hcl2fmt` formats all HCL files recursively from current working directory
`hcl2fmt -w dir` formats all HCL files recursively with `dir` as the base path

## Acknowledgement

It's a tiny tool, but being a Golang newbie I looked at both https://github.com/fatih/hclfmt and https://github.com/gruntwork-io/terragrunt `fmt` subcommand for tips and inspiration.
