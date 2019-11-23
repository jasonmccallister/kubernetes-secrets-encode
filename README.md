# Secrets YAML Encode CLI
[![Go Report Card](https://goreportcard.com/badge/github.com/jasonmccallister/secretsyaml)](https://goreportcard.com/report/github.com/jasonmccallister/secretsyaml)

A Go program that takes a Kubernetes secrets YAML and encodes all the secrets in the file into `base64` encoded values.

## Installation

TODO

## Usage

```bash
./secretsyaml project-secrets.yaml
```

If passed the following file as the first argument:

```yaml
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  name: some-example-secrets
data:
  SECURITY_KEY: thisissupersecret
  DB_DRIVER: dontstoreinplaintext
  DB_SERVER: thisisok
  DB_USER: thisisalsook
  DB_PASSWORD: neverstoreinplaintextorencodedbybase64
  DB_DATABASE: thisisokaswell
  DB_SCHEMA: totallyfineaswell
  DB_TABLE_PREFIX: yesthisto
  DB_PORT: thisispublicbutmaybedontstore
```

This will output a the YAML with secrets encoded as `encoded.yaml` with the following contents:

```yaml
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  name: some-example-secrets
data:
  DB_DATABASE: dGhpc2lzb2thc3dlbGw=
  DB_DRIVER: ZG9udHN0b3JlaW5wbGFpbnRleHQ=
  DB_PASSWORD: bmV2ZXJzdG9yZWlucGxhaW50ZXh0b3JlbmNvZGVkYnliYXNlNjQ=
  DB_PORT: dGhpc2lzcHVibGljYnV0bWF5YmVkb250c3RvcmU=
  DB_SCHEMA: dG90YWxseWZpbmVhc3dlbGw=
  DB_SERVER: dGhpc2lzb2s=
  DB_TABLE_PREFIX: eWVzdGhpc3Rv
  DB_USER: dGhpc2lzYWxzb29r
  SECURITY_KEY: dGhpc2lzc3VwZXJzZWNyZXQ=
```
