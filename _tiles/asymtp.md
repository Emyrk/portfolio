# TOTP 2 Factor Authentication

## Abstract

An asymmetric one time password solution that would replace the existing symmetric solutions, such as Google Authenticator. Symmetric solutions have a larger attack surface with there being a shared seed. This would allow users to control their keys, and mitigate risk in the case of a compromised server.

## Tech Stack

- Javascript library utilizing bitcoin HD wallet libraries

## Existing Solution

TOTP stands for "Time-based One Time Password" and is often used as a 2nd factor authentication for login. Google Authenticator is a very common TOTP app that generates these one time passwords that are valid for a certain period of time. This common method however is symmetric. This means that the private seed that is used on your phone, is also stored on the remote server. Any compromised server would render the 2nd factor authentication useless.

## Asymmetric Alternative

Asymmetric Cryptography has incredible advantages, and these advantages can be brought over to TOTP. If instead of using a shared key, the user can store the private key, and the server can store a public key that can verify the integrity of the OTP, then the risk of a compromised server will not expose all authentication factors.

# Asymmetric TOTP

I created a valid form of an asymmetric TOTP that could be used. Creating it was rather simple, as it leverages standard asymmetric signature methods and the bip39 HD wallet from bitcoin. This bip39 standard allows for a single private seed to generate an infinite number of deterministic keys. The keys are generated in a tree structure, and the user can generate a public seed for a subtree. This public seed allows for the derivation of the corrosponding public key component to the private key generation. In order to derive the same key, both sides need to use the same derivation path (location in the key tree).

The brilliance behind this, is that by sharing a public seed for a subtree, you have just allowed a 3rd party to verify any keys generated in the subtree. The method used to determine the generation in the demo, is based on time. So the derivation path changes every ~30 seconds, much like standard TOTP.

# Why I worked on this

I mainly started this thought process when thinking about another project of mine, Cryptid. Cryptid is a decentralized identity verification tool. It works leveraging physical security, biometrics, and blockchain. It promoted 3 factors authentication, however 2 of the 3 are physical security measures. For digital verification, Cryptid becomes weak. To create a digital verification side to Cryptid, the idea of using a 2nd factor authentication like Google Authenticator came to mine. A Cryptid ID could leverage a asymetric form of TOTP, and return 1 of the 2 physical verification methods.

If pairing each Cryptid ID with a private key component, much like used in credit cards today, Crypid could then be used for digital verification.

# Security

## Reusing a key

In the centralized TOTP, once a key is used, it is not allowed to be used again. This means anyone sniffing the network, could not reuse the key to gain access. In the decentralized model, one could not know if a key was used. To prevent a malicious party from tricking a user to gain access, a few methods could be used.

- Including (some math on) the domain in the derivation path, and informing a user when using the tool to verify integrity of the site.
- Using a diffe helmin exchange to derive a derivation path or part of. This would be a challenge and response model

## Key security

A user must be able to rotate their private key to ensure any compromised key can be replaced. It must also be included as it is more likely for a user to change phones. In the Cryptid application, no notification needs to be sent to any 3rd party using the public key side, as all key management would be done on chain.




