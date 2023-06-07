#!/bin/sh
# GO_PLS_VERSION=golang.org/x/tools/gopls@v0.12.0
# GO_OUTLINE_VERSION=github.com/ramya-rao-a/go-outline@latest
# GO_CODE_VERSION=github.com/stamblerre/gocode@v1.0.0
# /usr/local/go/bin/go install -v $GO_PLS_VERSION
# /usr/local/go/bin/go install -v $GO_OUTLINE_VERSION
# /usr/local/go/bin/go install -v $GO_CODE_VERSION

GO_PLS_VERSION=golang.org/x/tools/gopls@v0.12.0
GO_OUTLINE_VERSION=github.com/ramya-rao-a/go-outline@latest
GO_CODE_VERSION=github.com/stamblerre/gocode@v1.0.0

go version
go install -v github.com/onsi/ginkgo/ginkgo@latest
go install -v github.com/onsi/gomega@latest
go install -v $GO_PLS_VERSION
go install -v $GO_OUTLINE_VERSION
go install -v $GO_CODE_VERSION


code --install-extension golang.go

echo 'init終了、拡張機能をインストールしてください....'