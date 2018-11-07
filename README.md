<h1 align="center">Instrument Flight Rules</h1>

<div align="center">
<img src="https://user-images.githubusercontent.com/903488/47606510-3ecbd600-d9c9-11e8-91d4-6f13a813ee9e.png" alt="Instrument Flight Rules" />
</div>

An format for defining [concourse](https://concourse-ci.org/) resource behavior & requirements.

## Why?

I'm working on [Flightplans](https://github.com/waterborne-labs/flightplans), which is a service for developing pipelines via a web interface. During planning, I realized that I needed a repeatable way to understand the inputs/outputs of a resource. Those inputs and outputs are described in an Instrument Flight Rules file, `ifr.yaml`.

## Quickstart

Every resource needs to have an `ifr.yaml` (instrument flight rules) file defined.

_For now, these are defined in this repo, but if this project is successful, every resource should keep their instrument flight rules in their own repo._

Here's an _incomplete_ example `ifr.yaml` file for the [git resource](https://github.com/concourse/git-resource):

```yaml
version: 1
source:
- name: uri
  # whether this source parameter must be defined by the user
  required: true
  # Use `schema` to define the value type for this parameter. This uses JSON Schema.
  # https://json-schema.org/understanding-json-schema/reference/index.html
  # By default, we assume everything is a string. So this is optional.
  schema:
    type: string
  # an example value
  example: git@github.com:concourse/git-resource.git
- name: branch
  # `required` is optional, and defaults to `false`.
  required: false
  example: master
  # describe how this source parameter works, and even provide additional examples.
  description: |
    The branch to track. This is optional if the resource is only used in get steps; however, it is required when used in a put step. 
    If unset for get, the repository's default branch is used; usually master but could be different.
get:
- name: depth
  required: false
  example: 1
  description: If a positive integer is given, shallow clone the repository using the --depth option.
put:
- name: repository
  required: true
  # this parameter needs a path to something, most likely another resource or task output
  path_input: true
  example: source
  description: The path of the repository to push to the source.

outputs:
# this output is special, in that it should describe the contents of the directory created from a `get`
- path: /
  description: The cloned repo
# if you have specific files that always exist inside the root directory, add additional outputs, with the name being the path to that file. You should exclude the leading `/` in the path
- name: .git/committer
  example_contents: foo@bar.com
  description: |
    For committer notification on failed builds. This special file .git/committer which is populated with the email address of the author of the last commit. This can be used together with an email resource like mdomke/concourse-email-resource to notify the committer in an on_failure step.
```

## IFRS directory

Organized by git user/org directory, then `<repo>.yaml`. 
