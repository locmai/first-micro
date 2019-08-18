## My First Micro Project

A demonstration using Go Micro framework

### Description

TBD

### Prerequisites

Required packages and tools

- Golang 1.12
- Make
- PostgreSQL server
- Consul registry
- NATS server

### Configurations

Using .env file for each environment.

### Project structure

```
.
└── root //general configurations and scripts
    ├── deployments  //continuous integrations, deployments
    ├── packages //shared packages,libs
    └── services //minimal services
```

### Quick start

At root directory

```
make example
```

At each service direcotry:

```
make proto
make build
./<service_name>
```

or

```
make proto
go run main.go
```

### To-do list

#### Features:

- [ ] Blog service
- [ ] JWT authentication
- [ ] Common libs
- [ ] API Gateway

#### CI/CD:

- [ ] Build pipeline
- [ ] Deploy pipeline
- [ ] App configuration script
