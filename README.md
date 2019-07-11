# secret-yaml

[![Maintainability](https://api.codeclimate.com/v1/badges/e047b9311147b1e8b419/maintainability)](https://codeclimate.com/github/antonmarin/secret-yaml/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e047b9311147b1e8b419/test_coverage)](https://codeclimate.com/github/antonmarin/secret-yaml/test_coverage)
[![Build Status](https://travis-ci.org/antonmarin/secret-yaml.svg?branch=master)](https://travis-ci.org/antonmarin/secret-yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonmarin/secret-yaml)](https://goreportcard.com/report/github.com/antonmarin/secret-yaml)

## Goal

Lightweight tool to encrypt yaml values,
which you can quickly install right inside your pipeline.
Usable when you don't need centralized secrets management.

## Install

```
export OS=$(uname | tr '[:upper:]' '[:lower:]')
curl -LsSo /usr/local/bin/syml https://github.com/antonmarin/secret-yaml/releases/latest/download/syml-$OS
chmod +x /usr/local/bin/syml
```

## Usage

- `export SYML_SECRET=$(syml generateSecretKey)`
  generate secret and store inside env variable
- `syml encrypt --secret=${SYML_SECRET}
  ~/decryptedSecrets/secret.yaml > ~/encryptedSecrets/secret.yaml`
  encrypt values inside yaml-file and save to new file
- `syml decrypt --secret=${SYML_SECRET}
  ~/encryptedSecrets/secret.yaml > ~/decryptedSecrets/secret.yaml`
  decrypt values inside yaml-file and save to new file

[![asciicast](https://asciinema.org/a/256378.svg)](https://asciinema.org/a/256378)

`syml help` to get more about usage
