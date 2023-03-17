# iris-api

Public API go binary for integration with IRIS

Env Variables:

IRIS_COMPANY
IRIS_API_TOKEN
IRIS_HOST

## Rebuilding the openapi-based structs

- Copy the full openapi.json spec into scripts
- Run:

  $ make upgrade

To add additional items to the subset, edit `PATHS` in subset_maker.py

TODO:

- [ ] add tests for create path
- [ ] add API for getting company back from token and use.

---

- [ ] add command for getting config for a token
- [ ] add command for submission of logs
- [ ] API for logs (cloudwatch log group per api key)
