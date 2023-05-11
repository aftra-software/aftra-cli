import argparse
import json
from collections import defaultdict


def get_subset(data, paths, exclude_paths):
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

    for path in exclude_paths:
        steps = path.split(".")[:-1]
        final_key = path.split(".")[-1]
        focus_out = output
        for step in steps:
            focus_out = focus_out[step]
        del focus_out[final_key]

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
    "paths./api/companies/{parent_pk}/syndis-scans.get",
    "paths./api/token/.get",
    "paths./api/integrations/syndis-scan/{scan_name}/config.get",
    "paths./api/integrations/syndis-scan/{scan_name}/logs.post",
    "paths./api/integrations/syndis-scan/{scan_name}/scan.post",
    "components.schemas.MaskedToken",
    "components.schemas.SubmitLogEvent",
    "components.schemas.CreateOpportunity",
    "components.schemas.HTTPValidationError",
    "components.schemas.OpportunityScore",
    "components.schemas.PaginatedEntityCollection_SyndisScanEntity_",
    "components.schemas.SyndisScanTypes",
    "components.schemas.SyndisScanConfig",
    "components.schemas.SyndisScanEntity",
    "components.schemas.SyndisInternalScanEvent",
    "components.schemas.SyndisRiskScore",
    "components.schemas.ValidationError",
]

EXCLUDE_PATHS = [
    "components.schemas.SyndisScanEntity.properties.created",
    "components.schemas.SyndisScanEntity.properties.updated",
    "components.schemas.SyndisScanEntity.properties.pk",
    "components.schemas.SyndisScanEntity.properties.sk",
    "components.schemas.SyndisScanEntity.properties.entityType",
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
    print(json.dumps(get_subset(data, PATHS, EXCLUDE_PATHS), indent=4))
