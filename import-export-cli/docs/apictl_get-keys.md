## apictl get-keys

Generate access token to invoke the API or API Product

### Synopsis

Generate JWT token to invoke the API or API Product by subscribing to a default application for testing purposes

```
apictl get-keys [flags]
```

### Examples

```
apictl get-keys -n TwitterAPI -v 1.0.0 -e dev --provider admin
NOTE: Both the flags (--name (-n) and --environment (-e)) are mandatory.
You can override the default token endpoint using --token (-t) optional flag providing a new token endpoint
```

### Options

```
  -e, --environment string   Key generation environment
  -h, --help                 help for get-keys
  -n, --name string          API or API Product to generate keys
  -r, --provider string      Provider of the API or API Product
  -t, --token string         Token endpoint URL of Environment
  -v, --version string       Version of the API or API Product
```

### Options inherited from parent commands

```
  -k, --insecure   Allow connections to SSL endpoints without certs
      --verbose    Enable verbose mode
```

### SEE ALSO

* [apictl](apictl.md)	 - CLI for Importing and Exporting APIs and Applications

