## apictl delete-api-product

Delete API Product

### Synopsis

Delete an API Product from an environment

```
apictl delete-api-product (--name <name-of-the-api-product> --provider <provider-of-the-api-product> --environment <environment-from-which-the-api-product-should-be-deleted>) [flags]
```

### Examples

```
apictl delete-api-product -n TwitterAPI -r admin -e dev
apictl delete-api-product -n FacebookAPI -v 2.1.0 -e production
NOTE: Both the flags (--name (-n) and --environment (-e)) are mandatory.
If the --provider (-r) is not specified, the logged-in user will be considered as the provider.
```

### Options

```
  -e, --environment string   Environment from which the API should be deleted
  -h, --help                 help for delete-api-product
  -n, --name string          Name of the API to be deleted
  -r, --provider string      Provider of the API
```

### Options inherited from parent commands

```
  -k, --insecure   Allow connections to SSL endpoints without certs
      --verbose    Enable verbose mode
```

### SEE ALSO

* [apictl](apictl.md)	 - CLI for Importing and Exporting APIs and Applications
