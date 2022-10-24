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

The major technical contributions of this project will be:
- We build a private **data management framework** by leveraging the power of **functional encryption** and **smart contracts**. The proposed system will enable:
  - A fine-grained `data management layer` for the multi-chain world and Web 3.0, including a private data sharing platform between TradFi and DeFi, NFT-gated access, etc.
  - A Fine-grained `access control policy layer` is proposed so that it can be built into the smart contracts to be deployed in multi-chain.
- We devise an access control mechanism to ensure that legitimate buyers, on receiving the data owner’s permission, can get access to the target data
- We introduce a payment system, which will not only enforce payment privacy when needed but allow the accrue of value transferred among different parties in the system.

# Architecture
### Layers Structure
#### Policy Management Layer
The policy management layer defines the **access control** policy of the functional encryption scheme in the specific application scenarios.
> For instance, in the case of private data sharing between TradFi and DeFi, it will be the regulatory authorities and users who will be responsible for defining exactly what kind of private KYC information is allowed to leak to what kind of third-party DeFi apps.

#### Storage Layer
The storage layer is responsible for storing the encrypted message under functional encryption scheme. The *digital watermarking technology❓* will add an additional authenticity guarantee to the underlying message.

#### Credential Issuance Layer
It will be responsible for setting up all the key authorities in the system.

### Application Layer
...

## Functional Encryption
### Background
There are three relevant projects to Ruby:
1. The first one is perhaps the Enigma project, a privacy protocol that **enables the creation of decentralized applications (DApps) that guarantee privacy**. The protocol Enigma is based on secure `multi-party computation` (MPC). 
2. The second one, Insights Network, is a **data exchange protocol** based on combining blockchain technology, Substrate module, and MPC. *It is based on the EOS blockchain and a custom MPC system.❓*
3. The third one, NuCypher, is a **cryptographic infrastructure** for privacy-preserving applications. Its main technology is threshold proxy re-encryption and fully homomorphic encryption.

Compared with traditional public-key encryption schemes that only allow the message receiver to either decrypt the whole data set or nothing, functional encryption allows the sender to determine exactly which part of the message can be decrypted by exactly which kinds of receivers.（公钥加密模式解密全部信息，但functional加密可以定义不同的接收者能够解密的部分）