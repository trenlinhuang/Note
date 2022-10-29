source: https://docs.raze.network/getting-started/technical-architecture

## Project Architecture

it will have three technical modules: mint, transfer and redeem.
- The mint module will convert a base token into its anonymized version
- while the redeem module will convert the anonymized token back to its native form.
- The transfer module is the one that enables the anonymous transfer of the anonymized token. This process will conceal the transaction amount and guarantee the anonymity for both the sender and receiver.

## Zero-knowledge Implementation
**Mint:**
The mint module will create a Raze account by running a mint contract and **deposit the anonymized token to the Raze account**. Each Raze account is identified by a public key pk and **the mint module will generate a ciphertext under the account public key which encrypts the amount of minted token**. If the account has already existed before the mint operation, the generated ciphertext will be **homomorphically added to the existing ciphertext to increase the encrypted amount**.

**Transfer:**
Since the balance and transaction **amount of each Raze account is encrypted** and therefore hidden, the remaining question for the transfer module is **how to hide the identities of the involved parties**. ... we hide the identities of the involved parties through the "one-out-of-many" proof.

**Redeem:** The redeem module converts the anonymized token back to its original form. The redeem module also needs to invoke a zero-knowledge proof functionality to prove that the user initiating the redeem module knows the secret key of the corresponding Raze account ...

Whenever a user deposits a certain amount of token through invoking the mint module, the token is no longer traceable.

### Mint
To invoke the Mint contract, the client-side runs a CreateMintTx algorithm,

which takes as input:
- the raze account pk
- and the amount of native token amt as inputs. 

The output of the CreateMintTx algorithm is:
- a ciphertext cp_1 encrypted under the public key pk.

If the raze account pk is already attached with a ciphertext cp_0, the newly created ciphertext will be **homomorphically added to the existing ciphertext to increase the encrypted amount**. The updated ciphertext will be formed as cp_1*cp_0. Otherwise, the new ciphertext will be attached to the raze account. **The native token will be stored in the mint contract**.

### Transfer
To invoke the transfer contract, the client-side runs a CreateTransferTx algorithm, which takes as inputs:
- the raze account secret key sk,
- and the amount of transferred token amt,
- the public keys of sender pk_s, receiver pk_r
- and the public keys of the anonymity set {pk_a}.

The output of the CreateTransferTx algorithm is:
- a zero-knowledge proof that the prover knows one of the secret keys of the aforementioned public key set,
- the payment consistency proof,❓
- and range proof.❓

### Redeem
The client-side runs a CreateRedeemTx algorithm to invoke the redeem contract. It takes:
- the account secret key sk,
- the withdrawal amount amt,
- and the public key pk

as input to generate a zero-knowledge proof showing that the user knows the secret key sk for the account public key pk and the account has enough balance for the redeem operation.

