# verifiable-credential-payment

This repository demonstrates how [go-perun] can be used to realize trustless credential payment.
The protocol guarantees that the credential issuer receives the payment if and only if the requested credential is correctly issued.
It is based on a smart contract deployed on a blockchain.

More details on the protocol can be found in [PROTOCOL](PROTOCOL.md).
More details on plans to integrate this implementation with the Hyperledger Aries Framework can be found in [INTEGRATION](INTEGRATION.md).

## Development

### Test
Ensure that [go] and [ganache-cli] are installed.
```sh
go test ./... -v
```

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

- Rename to 'perun-credential-payment'?