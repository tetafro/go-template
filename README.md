# go-template

A template for a my Go projects.

## Project settings

1. Add [secrets](https://github.com/tetafro/fake-news/settings/secrets/actions)
   for SSH access and Ansible Vault.
2. [Setup repo for Codecov](https://app.codecov.io/gh/tetafro?repoDisplay=Inactive),
   add `CODECOV_TOKEN` env.

## Build and run

Copy and populate config
```sh
cp config.example.yaml config.yaml
```

Start
```sh
make build run
```
