module github.com/lightningnetwork/lightning-onion

go 1.12

require (
	github.com/aead/chacha20 v0.0.0-20170221132220-d31a916ded42
	github.com/btcsuite/btcd v0.0.0-20180524035114-bc0944904505
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcutil v0.0.0-20190207003914-4c204d697803
	github.com/davecgh/go-spew v0.0.0-20171005155431-ecdeabc65495
)

replace github.com/btcsuite/btcd => ../btcsuite/btcd
