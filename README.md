# Go Fermentation Monitor

Work in progress.

This repo is to be installed on a RaspberryPi to monitor the fermentation status of beer. It currently only works with Tilt fermentation devices. 

### Develop locally with:

```
go mod edit -replace github.com/alexhowarth/go-tilt=/<path to go-tilt>/go-tilt
```

### Start Go + Svelte applications:

```
➜  go-ferment-monitor git:(main) ✗ go run main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2021/06/08 15:10:23 Goroutines: 5
2021/06/08 15:10:23 Scanning for 2m0s
[GIN-debug] GET    /v1/mqtt/                 --> github.com/alexhowarth/go-ferment-monitor/controller.(*MQTTController).GetMQTTConfig-fm (5 handlers)
[GIN-debug] PUT    /v1/mqtt/                 --> github.com/alexhowarth/go-ferment-monitor/controller.(*MQTTController).UpdateMQTTConfig-fm (5 handlers)
[GIN-debug] GET    /v1/tilt/                 --> github.com/alexhowarth/go-ferment-monitor/controller.(*TiltController).GetTilts-fm (5 handlers)
[GIN-debug] PUT    /v1/tilt/                 --> github.com/alexhowarth/go-ferment-monitor/controller.(*TiltController).UpdateTilt-fm (5 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

```
➜  web git:(main) ✗ npm run dev

> svelte-app@1.0.0 dev /Users/alex/src/go-ferment-monitor/web
> rollup -c -w

rollup v2.23.0
bundles src/main.js → public/build/bundle.js...
Browserslist: caniuse-lite is outdated. Please run:
npx browserslist@latest --update-db
LiveReload enabled
...
  Your application is ready~! 🚀

  - Local:      http://localhost:5000
  - Network:    Add `--host` to expose
  ```