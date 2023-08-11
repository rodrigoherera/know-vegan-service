# Know vegan service

Know vegan service is the backend service for React Native Application [know-vegan-app](https://github.com/rodrigoherera/know-vegan-app).
The backend is used to enable the CURL for Categories, Products and Ingredients, using Go, Gin and GORM to control the api access and the way to comunicate with the SQL.

## Setup

You need to have running a MySQL server

```
go mod download
go run main.go
```

## Usage
Normally to access to the backend from the app is neccessary to expose a public URL, with the localhost is not enought, here you can use [ngrok](https://ngrok.com/) to attach your localhost endpoint to the public URL that ngrok gives you and with that you can access publicly to the resources.