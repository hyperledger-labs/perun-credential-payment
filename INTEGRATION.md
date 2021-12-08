# Integration within Hyperledger Aries

[Hyperledger Aries] is a framework that allows trusted online peer-to-peer interactions based on decentralized identities and verifiable credentials.

We describe how the Perun-based implementation of trustless credential payment can be integrated with [​​Aries RFC 0453: Issue Credential Protocol 2.0]. We will do that on the basis of the [Alice and Faber Demo].

## Alice and Faber Demo

The Alice and Faber Demo describes a setting where a former student, Alice, connects with her former College, Faber, and asks the college to issue her a digital verifiable credential for her degree. The protocol in the demo involves the following steps:

1. Alice establishes connection with Faber.
2. Alice requests the credential to be issued.
3. Faber issues the credential.

## Integration of Perun credential payment
We will now describe how Perun-based credential payment can be integrated into the credential issuance process between Alice and Faber.

| Alice and Faber Demo | Perun credential payment |
|-|-|
| 1. Alice establishes connection with Faber. | Alice opens a perun channel with the credential payment app installed. |
| 2. Alice requests the credential to be issued. | Alice proposes a channel update where the specifies the credential to be issued and the payment to be made. |
| 3. Faber issues the credential. | Faber proposes a channel update that reveals the credential to Alice and thereby releases the payment to Faber. |



[Alice and Faber Demo]: https://github.com/hyperledger/aries-cloudagent-python/blob/main/demo/README.md
[​​Aries RFC 0453: Issue Credential Protocol 2.0]: https://github.com/hyperledger/aries-rfcs/tree/b3a3942ef052039e73cd23d847f42947f8287da2/features/0453-issue-credential-v2
[Hyperledger Aries]: https://github.com/hyperledger/aries
