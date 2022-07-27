# famed-submission-form

This respository contains a prototype of the famed submission form.
The submission form is intended to be used as a Web3 replacement for Google form.

## How it Works

### Edit Form

Login with Metamask with account with address equal to NEXT_PUBLIC_OWNER.
Click on edit tab and make changes to the submission form. Currently only text based questions can be added. The submission form changes in this prototipical implementation are not saved beyond a page reload.

### Submit Form

Fill in the answers to the question in the `Respond` tab. Login with Metamask and click `Send Form`. You will be requested to approve the siging of the to be submitted message by Metamask. The form will be submitted using [Waku](https://wakuconnect.dev/). The current state of the waku network does not guarantee instant submission. The backend will decript, verify the signature and log the transmitted message as soon as it is forwarded by the waku network.

## How to Use

Set up a .local.env variable in `/frontend` containing:

```
NEXT_PUBLIC_PUBLIC_ENCRYPTION_KEY=<Public encryption key for metamask based enryption has to derived from ETH_PRIVATE_KEY used in backend>
NEXT_PUBLIC_OWNER=<Ethereum Address of the owner of the submission form>
```

Install frontend libraries:

```
npm i
```

Start frontend:

```
npm run dev
```

Open new terminal and navigate to `/backend`:

Start backend:

```
ETH_PRIVATE_KEY=<Ethereum private key for decryption> go run main.go
```

### Backlog

- Implement submission filter based on soulbound tokens
- Implement interfaces to forward submissions
- Persist submission form changes
- Extend submission form elements to feature parity with google forms
