# famed-submission-form

This repository contains a prototype of the famed submission form.
The submission form is intended to be used as a Web3 replacement for Google Forms.

## How it Works

### Edit Form

Login with Metamask with an account with an address equal to NEXT_PUBLIC_OWNER.
Click on the `Edit` tab and make changes to the submission form. Currently, only text-based questions can be added. The submission form changes in this prototypical implementation are not saved beyond a page reload.

### Submit Form

Fill in the answers to the question in the `Respond` tab. Login with Metamask and click `Send Form`. You will be requested to approve the signing of the to-be-submitted message by Metamask. The form will be submitted using [Waku](https://wakuconnect.dev/). The current state of the Waku network does not guarantee instant submission. The backend will decrypt, verify the signature and log the transmitted message as soon as the Waku network forwards it.

## How to Use

Set up a .local.env variable in `/frontend` containing:

```
NEXT_PUBLIC_PUBLIC_ENCRYPTION_KEY=<Public encryption key for metamask based encryption, derived from ETH_PRIVATE_KEY used in backend https://docs.metamask.io/guide/rpc-api.html#unrestricted-methods -> eth_getEncryptionPublicKey>
NEXT_PUBLIC_OWNER=<Ethereum Address of the owner of the submission form>
```

**Note:** eth_getEncryptionPublicKey is [depreacted](https://medium.com/metamask/metamask-api-method-deprecation-2b0564a84686) an will be replaced as described in the [backlog](#Backlog).

Install frontend libraries:

```
npm i
```

Start frontend:

```
npm run dev
```

Open a new terminal and navigate to `/backend`:

Start backend:

```
ETH_PRIVATE_KEY=<Ethereum private key for decryption> go run main.go
```

### Backlog

- Implement a submission filter based on Soulbound Tokens
- Implement interfaces to forward submissions
- Persist submission form changes
- Extend submission form elements to feature parity with Google Forms
- Replace Ethereum-based encryption with Waku's native encryption
- Use forward secrecy enabled encryption schema 
