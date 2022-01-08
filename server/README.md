# Introduction

This folder contains reusable go module for diverse purposes

## Documentation

https://github.com/golang/go/wiki/Mobile#building-and-deploying-to-android
https://pkg.go.dev/golang.org/x/mobile/cmd/gobind

# Tests

## Acceptance tests

```bash
go test ./...
```

## End to end tests

e2e tests are meant to fully test the peeral lib.<br>
The test run will take several minutes.

```bash
godog
```

# Compile to aar

## Install gomobile

Inside a go module folder, install gomobile :
```bash
export PATH=$PATH:/home/$USER/go/bin
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
```

## build to aar

if you installed android with android studio, do not forget to setup $ANDROID_HOME and $ANDROID_NDK_HOME
```bash
export ANDROID_HOME=/home/$USER/Android/Sdk
export ANDROID_NDK_HOME=/home/$USER/Android/Sdk/ndk/<ndk_version>=19>
```

```bash
go get -d golang.org/x/mobile/cmd/gomobile
gomobile bind -target=android peeral.com/proxy-libp2p/libp2p peeral.com/proxy-libp2p/libp2p/interfaces
```
(you can't build "main" package)