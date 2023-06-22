GOCMD=go build .

debug:
	$(GOCMD)
	dlv exec ./simple-contractapi-cc
	# set peer address on shim.go:L40: *peerAddress = "peer0.org1.example.com:7051"
	# main.go:L73 -> contract_chaincode.go:L119 -> shim.go:L68 -> shim.go:L40

debug-with-peer-address:
	$(GOCMD)
	dlv exec ./simple-contractapi-cc -- -peer.address localhost:7051

run:
	go run main.go
