## Using Vault API with Go

Assuming you have a secret in Vault, writing on path, eg.:

``` shell
kv/customer-service
```

When you are using the cli, you need to use:

```shell
$ vault kv get kv/customer-service
```

Now, using the API, you go to use:

```go
...
results, err := client.Logical().Read("kv/data/customer-service")
...

```

Now, you have a map[] with your secrets from Vault.
