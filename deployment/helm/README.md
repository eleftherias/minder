# Helm charts for Minder server

(These are a work in progress)

Stacklok runs Minder on Kubernetes, using these helm charts. In order to use
these helm charts, you will need the following:

- Postgres running in your cluster. In particular, you need a `postgres` Service
  in the `postgres` Namespace. The [`k8s-dev-setup`](../k8s-dev-setup/)
  directory has a sample postgres for local development and testing purposes.
  Note that the Postgres in that setup is **NOT** durable.

- The following Kubernetes secrets:

  - `minder-auth-secrets`: Needs to contain public and private keys for access
    and refresh tokens. The keys must be named as follows:

    - `access_token_rsa`, `access_token_rsa.pub`
    - `refresh_token_rsa`, `refresh_token_rsa.pub`

  - `minder-github-secrets`: Needs to contain API credentials for a GitHub
    app. In particular, the following keys are required:
    - `client_id`: The GitHub client ID to be used by Minder
    - `client_secret`: The GitHub client secret to be used by Minder

  - `minder-identity-secrets`: Needs to contain the OAuth 2 client secret for Minder
    server when authenticating with Keycloak. In particular, the following keys are required:
    - `identity_client_secret`: The Keycloak client secret to be used by Minder server

- In addition, if you are using Minder images which require authentication,
  you will want to create a `docker-registry` type credential with the name
  `minder-pull-secret`

## Building and running

You can build a (local) helm chart with `make helm` at the top-level of the
Minder repository. You can then run it with:

```helm
helm install minder config/helm/minder-0.1.0.tgz
```

Note that the helm chart does not specify a namespace, so Minder will be
installed in the namespace specified by your current Kubernetes context.
