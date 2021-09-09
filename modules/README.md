# Introduction

This folder contains reusable go module for diverse purposes

## Documentation

https://github.com/golang/go/wiki/Mobile#building-and-deploying-to-android
https://pkg.go.dev/golang.org/x/mobile/cmd/gobind

# Compile to aar

## Install gomobile

Inside a go module folder, install gomobile :<br>
go install golang.org/x/mobile/cmd/gomobile@latest<br>
gomobile init

Do not forget to add /home/\<user>/go/bin to PATH

## build to aar

go get -d golang.org/x/mobile/cmd/gomobile
gomobile bind -o test.aar -target=android peeral.com/proxy-libp2p/proxy
(you can't build main package)