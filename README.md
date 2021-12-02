# verifiable-credential-payment

This repository demonstrates how [go-perun] can be used to realize secure trustless credential payment.
The protocol guarantees that the credential issuer receives the payment if and only if the requested credential is correctly issued.
It is based on a smart contract deployed on a blockchain.

## Run test
**Prequisite:** Have [go] and [ganache-cli] installed.
```sh
go test ./... -v
```

## Development

### Compile smart contract

This step is only necessary if you want to make changes to the smart contract.
It requires [abigen] and [solc].

```sh
abigen --pkg app --sol app/CredentialSwap.sol --out app/CredentialSwap.go --solc solc
```

[abigen]: https://github.com/ethereum/go-ethereum
[ganache-cli]: https://github.com/trufflesuite/ganache
[go]: https://go.dev
[go-perun]: https://github.com/hyperledger-labs/go-perun
[solc]: https://docs.soliditylang.org/en/v0.8.10/installing-solidity.html



## TODO

- Check balance before transition into offer state.