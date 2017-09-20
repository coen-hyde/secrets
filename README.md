Secrets
=======

Secrets is an experimental tool to help manage your application's development secrets within a Team.

It encrypts your secrets to a `.secrets` directory, which should then be committed to git. This coupling of secrets with code ensures that when the application is executed, it will have the correct secrets configuration. To grant your team mates access to the secrets repository simply do `secrets members add ${keybase_username}`.

The encryption is performed by the local Keybase service running on your machine. This allows Secrets to pass off all the hard work of making sure encryption is done right to Keybase. It also means Secrets doesn't need access to your Keybase authentication details. As long as you are logged into Keybase on your local machine Secrets can communicate to it via a socket.

Installation and Usage
----------------------
Please see https://secrets.team for installation and usage details.

Contributing
------------

### Publishing a Release

#### 1. Create a git tag
Create a git tag eg. `git tag -a v0.1.0` and push the tags to github. The tag used here will be the version name compiled into secrets for `secrets --version`.

#### 2. Build Release binaries
Run `make release`. Binaries for Mac, Windows and Linux will be built.

#### 3. Create a Github Release
Create a github release on the tag version created earlier.

#### 4. Update brew tap to point to new release
Update the [Brew Forumla for Secrets](https://github.com/bugcrowd/homebrew-cartons/blob/master/Formula/secrets.rb). You'll need to calculate the sha256 hash of the mac build of secrets. You can do this via `shasum -a 256 -b {{mac binary}}`

The Future
----------

- Add tests
- Consider supporting other encryption backends such as pgp.
