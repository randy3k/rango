build:
	xgo --out dist/rango --targets=windows-6.0/*,darwin-10.9/amd64,linux/* ./cmd/rango
