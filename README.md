# SSH authentication bug

A sandbox to reproduce an authentication issue when having the SSH server built with https://github.com/mscdex/ssh2 and the SSH client built with https://pkg.go.dev/golang.org/x/crypto/ssh.

## Building

### Server

Requires [Node.js](https://nodejs.org/) being installed on your system.

```bash
$ cd server
$ npm install
$ node index.js
```

It started successfully if it prints the following message:

```
SSH server successfully started on ssh://127.0.0.1:8022
```

### Client

Requires [Golang](https://go.dev/) being installed on your system.

```bash
$ cd client
$ go build
$ ./ssh-client
```

## Issue

When starting the ssh-client while having the provided example server running, it prints the following error message:

```
2024/03/21 09:07:40 Failed to dial: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain
```

The server prints the following messages for the client authentication request:

```
SSH Client connected
authentication none test
Auth request with method none
authentication publickey test
User test successfully authenticated with public key
Client disconnected
```

Even though the server offers and accepts the `publickey` authentication method, the client fails the handshake.

## OpenSSH Client

When using an openssh client, it works:

```bash
$ mkdir -p openssh
$ cd openssh
$ ssh-keygen -t rsa -b 2048 -f id_rsa
$ ssh -v -i id_rsa -p 8022 -l test localhost
```

The OpenSSH client debug logs contain the following lines:

```
debug1: Offering public key: ...
debug1: Server accepts key: ...
Authenticated to localhost ([127.0.0.1]:8022) using "publickey".
```

And the server will print the following lines:

```
SSH Client connected
authentication none test
Auth request with method none
authentication publickey test
User test successfully authenticated with public key
Client authenticated!
Client disconnected
```
