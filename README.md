# ssi-credential-payment

## Compile Solidity and generate Go bindings

```sh
abigen --pkg app --sol app/CredentialSwap.sol --out app/CredentialSwap.go --solc solc
```

## Run test
Prequisite: Have [ganache cli *TODO insert link*](TODO) installed.
```sh
go test ./... -v -p 1
```


## TODO

- README: Run test: Insert ganache link
- Check ganache executable default name
- Add CI
