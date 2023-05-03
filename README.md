# aftra-api

Public API go binary for integration with AFTRA

Env Variables:

- AFTRA_API_TOKEN: Token for communicating with the AFTRA api
- AFTRA_COMPANY: Company ID associated with the token (Retrieved using `aftra-api get token company`)

- AFTRA_HOST: Location of the host. Used during testing of the CLI client.

## Rebuilding the openapi-based structs

- go generate ./...

To add additional items to the subset of openapi schema being used, edit `PATHS` in subset_maker.py

## Example usage

| Command                                                          | Description                                                        |
| ---------------------------------------------------------------- | -------------------------------------------------------------------|
| `aftra-api create opportunity`                                   | Create an internal opportunity in Aftra                            |
| `aftra-api submit <scan-type> <scan-name> --message <msg>`       | Submit a raw scan event to the specified scanner                   |
| `aftra-api submit <scan-type> <scan-name> --filename <filename>` | Submit a file of raw scan events to the specified scanner          |
| `aftra-api get token`                                            | Get current token information in json format                       |
| `aftra-api get company`                                          | Get current token company information only                         |
| `aftra-api get config <scan-type> <scan-name> `                  | Get a scan config                                                  |
| `aftra-api log <scan-type> <scan-name> <msg>`                    | Log the contents of msg to Aftra. It will be viewable viat the API |
| `your_command.sh \| aftra-api log <scan-type> <scan-name>`       | Log from stdout to Aftra. It will be viewable viat the API         |

### Create opportunity

- uid: This should uniquely identify the opportunity. Creating with the same uid will result
  in an update to the existing one.
- details: Additional information in the form of key,value pairs. These are presented to the user in Aftra.
- name: The display name for the opportunity.
- score: Risk score (critical, high, medium, low, info, none, unknown)

## Getting started

1.  Export your token as AFTRA_API_TOKEN

    `$ export AFTRA_API_TOKEN=<token>`

2.  Export company id as AFTRA_COMPANY

    `$ export AFTRA_COMPANY=$(aftra-api get token company)`

3.  (Optional) Get any config required, and put somewhere that your script uses. The name is that defined on the
    config via the web UI.

    `$ aftra-api get config syndis-scan myscanner > config.ini`

4.  Create an opportunity

    `$ aftra-api create opportunity --uid=<uid> --name=<name> --score=<score> --details=<details>`

5.  Log out messages from stdin

    `$ ./my_opportunity_finder.sh | aftra-api log syndis-scan myscanner`
