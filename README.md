## 🚀 **Project Kit with Golang clean architecture**

**API Project Kit** This service is used for handle all endpoint and data about Module Users.

### **`Technology Stacks 🍔`**

```
  -- Golang
  -- Gin
  -- Gorm
  -- Multi DB (Mysql or Postgresql)
  -- Migration Database
```

### **`Project Structures 🏢`**

```
.
└── README.md
└── main.go
└── go.sum
└── go.mod
└── env-example       (this will be the environment file)
└── Dockerfile
└── entity/
└── helper/
└── services/
    └── handler/
        └── ...[.go]
    └── repo/
        └── ...[.go]
    └── usecase/
        └── ...[.go]
└── utils/
    └── config/
        └── ...[.go]
    └── database/
        └── ...[.go]
    └── middleware/
        └── ...[.go]
    └── public/


```

### Flow Development

During the development cycle, a variety of supporting branches are used :

- **\*feature/\*\*** -- feature branches are used to development new features for the upcoming releases. May branch off from development and must merge into development.
- **\*hotfix/\*\*** -- hotfix branches are necessary to act immediately upon an undesired status of master. May branch off from master and must merge into master, staging, and development.

Creating a new **_feature_**

1. create new branch from development. ex: `feature/name-of-feature`.
1. write your code.
1. commit & push your work to the same named branch on the server.
1. create PR into development branch for testing in dev server.

Creating a new **_hotfix_**

1. create new branch from master. ex: `hotfix/name-of-hotfix`.
1. write your code.
1. commit & push your work to the same named branch on the server.
1. create PR into master branch.

### How to install

1. clone this repo [restapps-service](#) into `go/path/`
1. go to api-on-premise `cd go/path/api-on-premise`
1. copy env-example into .env `cp env-example .env`
1. ajust config in .env
1. run project `go run main.go`

### Deployment

This flow of deployment using Git Flow with 3 main branches

- master -- this branch contains production code. All development code is merged into master in sometime.
- staging -- this branch is a nearly exact replica of a production environment for software testing.
- development -- this branch contains pre-production code. When the features are finished then they are merged into develop.

### Command

- run program `go run main.go`

### Build the binary

Just run `go build -o restapps-service main.go` it build the binary named `restapps-service`

### Migration

- create migration `migrate create -ext sql -dir utils/database/migrations create_(table)_table`
- up migration `migrate -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" -path src/database/migrations up`
- down migration `migrate -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" -path src/database/migrations down`
