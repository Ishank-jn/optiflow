# OptiFlow
Electronic Data Interchange (EDI) system designed to streamline the exchange of documents between parties.

## Architecture
/edi-system
│
├── main.go
│
├── /internal
│   ├── /api
│   │   └── handlers.go
│   ├── /config
│   │   └── config.go
│   ├── /db
│   │   └── database.go
│   ├── /models
│   │   └── edi.go
│   ├── /services
│   │   └── edi_service.go
│   ├── /utils
│   │   └── utils.go
│   └── /oauth
│       └── oauth.go
│
├── /pkg
│   ├── /middleware
│   │   └── auth.go
│   └── /socket
│       └── socket.go
|
|── License
├── go.mod
├── go.sum
└── README.md
