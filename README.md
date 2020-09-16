# dekms
Simple utility for decrypting data with Google KMS when you do not have access to `gcloud`. Data is read from stdin, and decrypted data is writted to stdout.

## Usage
```
usage: dekms --keyring=KEYRING --key=KEY [<flags>]

Flags:
  --help               Show context-sensitive help (also try --help-long and --help-man).
  --keyring=KEYRING    Name of keyring
  --key=KEY            Name of key
  --location="global"  Location of keyring
  --project=PROJECT    GCP project. Attempts read project from metadata server if not given.
```
