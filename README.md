# kubenorm

Kubenorm is a tool to normalize Kubernetes resources. It is designed to be used as a pre-commit hook to ensure that all resources are formatted in a consistent way.

## Installation

```bash
go install github.com/Avielyo10/kubenorm
```

## Usage

```bash
echo 24Gi | kubenorm [flags]
```

## Flags

| Flag | Description |
|------|-------------|
| `-n` | nano        |
| `-u` | micro       |
| `-m` | milli       |
| `-k` | kilo        |
| `-M` | mega // default |
| `-g` | giga        |
| `-t` | tera        |
| `-p` | peta        |
| `-e` | exa         |
| `-h` | help        |
| `-v` | version     |
