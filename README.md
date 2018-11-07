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

Check out [ifr for concourse/git-resource](./ifrs/concourse/git-resource.yaml) for an example.

The tool supports both yaml and json resource definitions. Json is assumed by default unless you pass the `-y` flag.

## IFRS directory

Organized similar to that of go imports, e.g `github.com/<user>/<repo>.yaml`.
