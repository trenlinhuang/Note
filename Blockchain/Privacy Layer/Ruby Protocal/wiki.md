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

Compared with traditional public-key encryption schemes that only allow the message receiver to either decrypt the whole data set or nothing, functional encryption allows the sender to determine exactly which part of the message can be decrypted by exactly which kinds of receivers.（💡公钥加密模式解密全部信息，但functional加密可以定义不同的接收者能够解密的部分）

### Introduction
Function encryption was first proposed by Boneh, Sahai and Waters in 2011. This encryption al- gorithm broke the original ”All-or-Nothing” data query mode. **The computing results of functional encrypted data can be obtained without revealing the plaintext content**.（❓不揭露明文的情况下解密一些信息？）

A classic Functional Encryption Scheme contains four algorithms, namely:
- FE.Setup: (pk,msk) ← Setup(1<sup>λ</sup>). 生成主私钥msk和公钥pk
- FE.KeyGen: sk ← KeyGen(msk, f). 用主私钥来生成函数f的一个新私钥sk
- FE.Enc: c ← Encrypt(pk, x). 用公钥加密消息x
- FE.Dec: y ← Decrypt(sk, c). 用函数f的私钥解密密文得到消息y

... the data owner can specify who can access the encrypted data and has complete control over the data, *Ruby will use the Ciphertext-Policy ABE (CP-ABE) or Key-Policy ABE (KP-ABE) algorithm in the Functional Encryption algorithm to encrypt their personal data*

### ABE Scheme
> **The empty key ε:** The special key ε in K captures **all the information about the plaintext that intentionally leaks from the ciphertext**, such as the length of the encrypted plaintext. The secret key for ε is empty and also denoted by ε. Thus, anyone can run dec(ε, c) on a ciphertext c ← enc(pp, x) and obtain all the information about x that intentionally leaks from c.

> **Predicate encryption** [BW07, KSW08]. In many applications a plaintext x ∈ X is itself a pair (ind, m) ∈ I × M where ind is called an index and m is called the payload message. *For example, in an email system the index might be the sender's name while the payload is the email contents.*

#### Attribute-Based Encryption（补充）
source: https://link.springer.com/content/pdf/10.1007/978-3-642-19571-6_16.pdf

Sahai and Waters [SW05] proposed a notion of encryption, called Attribute-Based Encryption (ABE), where **one could express complex access policies**. Subsequently, Goyal, Pandey, Sahai and Waters [GPSW06] refined this concept into two different formulations of ABE: `Key Policy ABE` and `Ciphertext-Policy ABE`.

**Key-Policy ABE**

We first describe Key-Policy ABE for boolean formulas, as was realized by Goyal et. al. [GPSW06] 4 . A Key-Policy ABE system over n variables can be described as a predicate encryption scheme (with public index) for the predicate Pn : K × I → {0, 1} where:
1. The key space K is the set of all poly-sized boolean formulas φ in n variables z = (z1, . . . , zn) ∈ {0, 1}<sup>n</sup>. We let φ(z) denote the value of the formula φ at z.（💡键空间是多项式函数φ集合，它的参数为长度为n的布尔向量）
2. The plaintext is a pair (ind = z, m) where the index space is I := {0, 1}<sup>n</sup>, and where we interpret z as a bit vector representing the boolean values z1, . . . zn.（💡明文是一个元组(ind = z, m)，索引（index）的定义域为长度为n的比特向量空间）
3. The predicate Pn on K × I is defined as<br>![alt](Predicat%20Pn.png)

**Ciphertext-Policy ABE**

A dual concept of Attribute-Based Encryption is CiphertextPolicy Attribute-Based Encryption (CP-ABE), where **the roles of the ciphertext and key are essentially reversed**.
1. The key space K := {0, 1}n is the set of all n bit strings representing n boolean variables z = (z1, . . . , zn) ∈ {0, 1}<sup>n</sup>.（💡键空间是索引，定义域为比特串）
2. The plaintext is a pair (ind = φ, m) where the index space I is the set of all poly-sized boolean formulas φ over n variables.（💡明文是一个元组(ind = φ, m)）


Based on the ABE scheme, Ruby will support the revocation and restoration of user attributes. The execution includes 7 steps:
1. FE.Setup
2. FE.Setup.Cancel: 

