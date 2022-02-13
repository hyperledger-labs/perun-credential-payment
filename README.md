# perun-credential-payment

This repository demonstrates how [go-perun] can be used to realize trustless credential payment.
The protocol guarantees that the credential issuer receives the payment if and only if the requested credential is correctly issued.
It is based on a smart contract deployed on a blockchain.

More details on the protocol can be found in [PROTOCOL](PROTOCOL.md).
More details on plans to integrate this implementation with the Hyperledger Aries Framework can be found in [INTEGRATION](INTEGRATION.md).

## Development

### Test
Ensure that [go] and [ganache-cli] are installed.
Start ganache.
```sh
DEPLOYER_SK=0x50b4713b4ba55b6fbcb826ae04e66c03a12fc62886a90ca57ab541959337e897
HOLDER_SK=0x1af2e950272dd403de7a5760d41c6e44d92b6d02797e51810795ff03cc2cda4f
ISSUER_SK=0xf63d7d8e930bccd74e93cf5662fde2c28fd8be95edb70c73f1bdd863d07f412e
BALANCE=10000000000000000000

ganache-cli --host 127.0.0.1 --port 8545 --account $DEPLOYER_SK,$BALANCE --account $HOLDER_SK,$BALANCE --account $ISSUER_SK,$BALANCE --blockTime=0
```
Run the tests.
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
