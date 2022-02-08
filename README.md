# services.filter-group-management

Microservice allowing to manage filter groups

## Run locally

```
# Mount postgres dependency
make run.infra

# Add some fixtures
make run.docker.fixtures

# Build local binary
make build.local

# Run local binary with local config, on port 8080
make run.local

# Test a target
make run.group.create
```
