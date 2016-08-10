Secrets
=======

Secrets is an experimental tool to help manage your application's development secrets within a Team.

It encrypts your secrets to a `.secrets` directory, which should then be committed to git. This coupling of secrets with code ensures that when the application is executed, it will have the correct secrets configuration. To grant your team mates access to the secrets repository simply do `secrets members add ${keybase_username}`.

The encryption is performed by the local Keybase service running on your machine. This allows Secrets to pass off all the hard work of making sure encryption is done right to Keybase. It also means Secrets doesn't need access to your Keybase authentication details. As long as you are logged into Keybase on your local machine Secrets can communicate to it via a socket.

Installation
------------

Keybase 0.17 or later is required. If you used brew to install Keybase do:
`brew update && brew upgrade keybase`.

Install via `go get github.com/coen-hyde/secrets` and ensure `$GOPATH/bin` is in your `$PATH`.

Usage
-----

### Initialize a Secrets repository
`$ secrets init`

### Set a variable
`$ secrets set key=value`

### Get a variable
`$ secret get key`

### export all variables in a human friendly format
`$ secrets export`

### export all variables in JSON format
`$ secrets export -f json`

### list all members
`$ secrets members`

### add member
All members added to a Secrets scope require a Keybase device key.
`$ secrets members add $keybase_name`

The Future
----------

- Allow removal of members
- Add environment variable export
- Add import ability
- Allow multiple scopes. A scope is a container for secrets and member permissions.
- Allow multiple keybase users to be added at once.
- Consider supporting other encryption backends such as pgp.
