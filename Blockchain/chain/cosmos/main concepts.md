# Main Concepts

## A Blockchain App Architecture

### Consensus in Tendermint Core and Cosmos
Only the top 150 nodes (as ranked by staked ATOM) participate in the transaction finalization process as validators. Voting power is calculated as all ATOM tokens staked by a validator and its delegates. For example, if a given validator's voting power is 15% of the total voting power of all validators, then the validator can expect to receive the block creation privilege 15% of the time.

The created block is broadcast to the other validators, who are expected to respond promptly and correctly:
* Validators confirm candidate blocks.
* Validators absorb penalties for failing to do so.
* Validators can and must reject invalid blocks.
* Validators accept the block by returning their signature.

When sufficient signatures have been collected by the block creator, the block is finalized and broadcast to the wider network.

### Upgradeability of Chains
In a Tendermint blockchain, transactions are irreversibly finalized upon block creation, and upgrades are themselves governed by the block creation and validation process. This leaves no room for uncertainty: either the nodes agree to simultaneously upgrade their protocol, or the upgrade proposal fails.

### Application Blockchain Interface (ABCI)
Tendermint BFT packages the networking and consensus layers of a blockchain and presents an interface to the application layer, the Application Blockchain Interface (ABCI). Developers can focus on higher-order concerns and delegate peer-discovery, validator selection, staking, upgrades, and consensus to the Tendermint BFT.

The Tendermint BFT engine is connected to the application by a socket protocol. ABCI provides a socket for applications written in other languages. If the application is written in the same language as the Tendermint implementation, the socket is not used.

The Tendermint BFT is not concerned with the interpretation of transactions. That occurs at the application layer. Tendermint presents confirmed, well-formed transactions and blocks of transactions agnostically. Tendermint is un-opinionated about the meaning any transactions have.

### The Cosmos SDK
* The Cosmos SDK provides a development head start and a pre-established framework for new applications.
* The Cosmos SDK provides a rich set of modules that address common concerns such as governance, tokens, other standards, and interactions with other blockchains through the **Inter-Blockchain Communication Protocol** (IBC).

### Additional details
#### CheckTx
Tendermint cannot determine the transaction interpretation because it is agnostic to it. To address this, the Application Blockchain Interface includes a CheckTx method. Tendermint uses this method to ask the application layer if a transaction is valid. Applications implement this function.

#### DeliverTx
Tendermint calls the DeliverTx method to pass block information to the application layer for interpretation and possible state machine transition.

#### BeginBlock and EndBlock
BeginBlock and EndBlock messages are sent through the ABCI even if blocks contain no transactions.