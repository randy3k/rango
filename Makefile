all:

xgo:
	xgo --out dist/rango --targets=windows-6.0/*,darwin-10.9/amd64,linux/* ./cmd/rango

goreleaser:
	docker run --rm --privileged -v $$PWD:/go/src/github.com/randy3k/rango -w /go/src/github.com/randy3k/rango bepsays/ci-goreleaser goreleaser --snapshot --rm-dist
