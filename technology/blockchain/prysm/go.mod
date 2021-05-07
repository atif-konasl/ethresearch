module github.com/atif-konasl/eth-research/technology/blockchain/prysm

go 1.16

require (
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/prysmaticlabs/eth2-types v0.0.0-20210303084904-c9735a06829d // indirect
	github.com/prysmaticlabs/ethereumapis v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/grpc v1.37.0
)

replace github.com/prysmaticlabs/ethereumapis => github.com/lukso-network/vanguard-apis v0.0.0-20210506184759-c857a4b244df
