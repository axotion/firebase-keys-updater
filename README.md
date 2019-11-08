# Firebase JWT public tokens updater

Package for dealing with new JWT public tokens every N minutes

## Why?

Firebase uses multiple private keys to generate the idToken, therefore the token we receive from Firebase is decodable with one of the Firebase public keys.

The header part of every the Firebase JWT token is:

```json
{
  "alg": "RS256",
  "kid": "19f07ad8152b2fc4e427cb25e9306edaca41a635"
}
```


Where the key kid corresponds to the id of the public key that we must use to verify JWT token


```json
{
  "5dce7e41add121b86ed404b84da75739467ed2bc": "-----BEGIN CERTIFICATE--- ...",
  "250811cdc609d90f98151191b2bc9bd0ceb9c004": "-----BEGIN CERTIFICATE--- ..."
}
```

## Ok, but why we must update keys every few hours?

The public keys are subject to change/rotation and the whole idea of publishing them on a URL is the probability of "rotation". Google, periodically rotates its keys to conform their security standards.

## Usage

```bash
go build main.go ; ./main -every 50
```

50 means minutes

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)