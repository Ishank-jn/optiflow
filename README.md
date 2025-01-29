# OptiFlow
Electronic Data Interchange (EDI) system designed to streamline the exchange of documents between parties.

## Architecture
```
📦 
LICENSE
├─ README.md
cache
│  └─ cache.go
|
├─ configs
│  └─ config.yaml
|
├─ go.mod
├─ go.sum
|
├─ internal
|  ├─ api
|  |  └─ handlers.go
├  ├─ config
|  |  └─ config.go
├  ├─ db
|  |  └─ database.go
│  ├─ logger
│  │  └─ logger.go
│  ├─ metrics
│  │  └─ metrics.go
│  ├─ models
│  │  └─ edi.go
│  ├─ oauth
│  │  └─ oauth.go
│  ├─ services
│  │  └─ edi_service.go
│  └─ utils
│     └─ utils.go
|
├─ main.go
└─ pkg
   ├─ middleware
   │  ├─ auth.go
   │  └─ rate_limiter.go
   └─ socket
      └─ socket.go
```
