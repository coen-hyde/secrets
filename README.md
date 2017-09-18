Secrets
=======

Secrets is an experimental tool to help manage your application's development secrets within a Team.

It encrypts your secrets to a `.secrets` directory, which should then be committed to git. This coupling of secrets with code ensures that when the application is executed, it will have the correct secrets configuration. To grant your team mates access to the secrets repository simply do `secrets members add ${keybase_username}`.

The encryption is performed by the local Keybase service running on your machine. This allows Secrets to pass off all the hard work of making sure encryption is done right to Keybase. It also means Secrets doesn't need access to your Keybase authentication details. As long as you are logged into Keybase on your local machine Secrets can communicate to it via a socket.

Installation
------------

Keybase 0.18 or later is required.

`brew tap bugcrowd/cartons`

`brew install bugcrowd/cartons/secrets`

Usage
-----

### Initialize a Secrets repository
`$ secrets init`

### Set a variable
`$ secrets set key=value`

### Get a variable
`$ secret get key`

### Export all variables in a human friendly format
`$ secrets export`

### Export all variables in JSON format
`$ secrets export -f json`

### List all members
`$ secrets members`

### Add one or more members
`$ secrets members add ${keybase_username}`

### Remove one or more members
`$ secrets members remove ${keybase_username}`

All members added to a Secrets scope require a Keybase device key.

The Future
----------

- Add environment variable export
- Add import ability
- Allow multiple scopes. A scope is a container for secrets and member permissions.
- Consider supporting other encryption backends such as pgp.
