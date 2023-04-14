# iris-api

Public API go binary for integration with IRIS

Env Variables:

- IRIS_API_TOKEN: Token for communicating with the IRIS api
- IRIS_COMPANY: Company ID associated with the token (Retrieved using `iris-api get token company`)

- IRIS_HOST: Location of the host. Used during testing of the CLI client.


## Rebuilding the openapi-based structs

- go generate ./...

To add additional items to the subset of openapi schema being used, edit `PATHS` in subset_maker.py

## Example usage

| Command                                                            | Description                                                       |
| ------------------------------------------------------------------ | ----------------------------------------------------------------- |
| `iris-api create opportunity`                                      | Create an internal opportunity in Iris                            |
| `iris-api get token`                                               | Get current token information in json format                      |
| `iris-api get company`                                             | Get current token company information only                        |
| `iris-api get config <scan-type> <scan-name> `                     | Get a scan config                                                 |
| `iris-api log <scan-type> <scan-name> <msg>`                       | Log the contents of msg to Iris. It will be viewable viat the API |
| `your_command.sh \| iris-api log <scan-type> <scan-name>`          | Log from stdout to Iris. It will be viewable viat the API         |
        
### Create opportunity

- uid: This should uniquely identify the opportunity. Creating with the same uid will result
  in an update to the existing one.
- details: Additional information in the form of key,value pairs. These are presented to the user in Iris.
- name: The display name for the opportunity.
- score: Risk score (critical, high, medium, low, info, none, unknown)

## Getting started

1.  Export your token as IRIS_API_TOKEN

    `$ export IRIS_API_TOKEN=<token>`

2.  Export company id as IRIS_COMPANY

    `$ export IRIS_COMPANY=$(iris-api get token company)`

3.  (Optional) Get any config required, and put somewhere that your script uses. The name is that defined on the
    config via the web UI.

    `$ iris-api get config syndis-scan myscanner > config.ini`

4.  Create an opportunity

    `$ iris-api create opportunity --uid=<uid> --name=<name> --score=<score> --details=<details>`

5.  Log out messages from stdin

    `$ ./my_opportunity_finder.sh | iris-api log syndis-scan myscanner`

