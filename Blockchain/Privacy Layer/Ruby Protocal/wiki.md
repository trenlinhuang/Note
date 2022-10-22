# Welcome
- It proposes and implements a **privacy middle-layer** interacting with multi-chain.
- It is a fine-grained **private data access-control gateway** across different entities and organizations both in the decentralized and traditional financial world.

All the encrypted data will be stored in a **decentralized cloud** such as IPFS.

# Overview
### Bringing Privacy-as-a-Service to Web3
#### Privacy Layer-1 Blockchain
Ruby’s modular approach to data sharing and privacy protection makes it the ideal building block for privacy-compatible smart contract Dapps, while also acting as a privacy layer for protocols and Dapps.

#### Zkp Cryptographic Infrastructure
Ruby utilizes `Functional Encryption (FE)`, a leading-edge cryptographic solution that enables users to encrypt sensitive information on-chain, which can only be decrypted by holders of an approved private key.

#### Interoperability and Composability
Built to be the privacy protocol compatible with the multichain ecosystem, Ruby collators produce block candidates that are approved by Ruby nodes and validators. Once the block is accepted by the validators it will be added to the Ruby mainnet.❓没看懂什么意思

### Access Control with Privacy
#### Private KYC and Authentication
It focuses on performing KYC such as biometric authentication without leaking private identity information.

### Key Features
#### Privacy-Centric Web3 Dapps
Ruby’s FE smart contracts will serve as the foundation for privacy-centric smart contract Dapps built on the native Ruby blockchain. This will also act as a privacy layer for protocols and Web3 Dapps.

#### Privacy-as-a-Service
Our access control policy will be well-defined so that it is attuned to the specific application scenarios. A too general access control policy might render it unusable in reality given it might reveal too little information about the users. Therefore, the stakeholders of Ruby will periodically update the policy management standard.

#### Attribute and Policy Revocability
Since the attribute value or attribute policy for a secret key might change with time, we should provide a mechanism to revoke the respective secret key. For instance, consider the attribute age" a user's age constantly changes as time goes by.❓不理解

## Existing Problems
### Value Proposition
A natural requirement of institutional money is regulatory compliance. Due to the sensitive nature of know-your-client (KYC) information, users, especially newcomers to DeFi are usually reluctant to give out their **private information** to the DeFi apps mostly due to the fear of losing control of their private data.

In fact, `functional encryption` was originally proposed as a tool to enforce fine-grained access control over encrypted data [GPSW2006]. Therefore, it is a natural solution to all the aforementioned access control problems. ... functional encryption can also serve as a handy tool for the private data-sharing platform between TradFi and DeFi.