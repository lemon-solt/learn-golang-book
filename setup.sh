#!/bin/sh
GO_PLS_VERSION=golang.org/x/tools/gopls@v0.12.0
GO_OUTLINE_VERSION=github.com/ramya-rao-a/go-outline@latest
GO_CODE_VERSION=github.com/stamblerre/gocode@v1.0.0
/usr/local/go/bin/go install -v $GO_PLS_VERSION
/usr/local/go/bin/go install -v $GO_OUTLINE_VERSION
/usr/local/go/bin/go install -v $GO_CODE_VERSION

echo 'init終了、拡張機能をインストールしてください....'