import argparse
import json
from collections import defaultdict


def get_subset(data, paths):
    output = defaultdict(dict)
    for path in paths:
        focus_in = data
        focus_out = output
        steps = path.split(".")[:-1]
        final_key = path.split(".")[-1]

        for step in steps:
            focus_in = focus_in[step]
            out = focus_out.get(step, {})
            focus_out[step] = out
            focus_out = out
        focus_out[final_key] = focus_in[final_key]

    return output


def test_cases():
    basic_data = {"a": 1, "b": {"c": 3, "d": 4, "e": {"e1": "1"}}}
    openapi_like_data = {
        "paths": {"/some/url": {"get": {"get-stuff": 1}, "post": {"post-stuff": 2}}}
    }
    for input, paths, expected in [
        (
            basic_data,
            ["a", "b.d"],
            {"a": 1, "b": {"d": 4}},
        ),
        (basic_data, ["b.e"], {"b": {"e": {"e1": "1"}}}),
        (basic_data, ["b.c", "b.d"], {"b": {"c": 3, "d": 4}}),
        (
            openapi_like_data,
            ["paths./some/url.post"],
            {"paths": {"/some/url": {"post": {"post-stuff": 2}}}},
        ),
    ]:
        result = get_subset(input, paths)
        assert result == expected, f"Failure. Actual {result}; Expected {expected}"


PATHS = [
    "openapi",
    "info",
    "paths./api/companies/{company_pk}/opportunities/.post",
    "paths./api/integrations/syndis-scan/{config_name}.get",
    "paths./api/integrations/syndis-scan/{config_name}/logs.post",
    "components.schemas.MaskedToken",
    "components.schemas.SubmitLogEvent",
    "components.schemas.CreateOpportunity",
    "components.schemas.HTTPValidationError",
    "components.schemas.OpportunityScore",
    "components.schemas.ValidationError",
    "components.schemas.SyndisScanTypes",
    "components.schemas.SyndisScanConfig",
]


# test_cases()
if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        prog="SubsetMaker",
        description="Get a subset of a json file based on paths",
    )
    parser.add_argument("filename")  # positional argument

    args = parser.parse_args()

    data = json.loads(open(args.filename).read())
    print(json.dumps(get_subset(data, PATHS), indent=4))
