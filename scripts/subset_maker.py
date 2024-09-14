import argparse
import json
from collections import defaultdict


def get_referenced_objects(output):
    referenced_objects = set()
    def get_referenced_objects_recursively(data, path):
        if isinstance(data, dict):
            for key, value in data.items():
                if key == "$ref":
                    referenced_objects.add((value.split('#/')[-1].replace("/", ".")))
                else:
                    get_referenced_objects_recursively(value, f"{path}.{key}")
        elif isinstance(data, list):
            for idx, item in enumerate(data):
                get_referenced_objects_recursively(item, f"{path}[{idx}]")
    get_referenced_objects_recursively(output, "")
    return set(referenced_objects)

def get_included_paths_and_dependents(data, paths, output):
    accounted_for = set()
    path_list = list(paths)
    for path in path_list:
        if path in accounted_for:
            continue
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

        referenced = get_referenced_objects(focus_out[final_key])
        path_list.extend(referenced)
        accounted_for.add(path)


def get_subset(data, paths, exclude_paths, replace_with_raw_paths):
    output = defaultdict(dict)

    get_included_paths_and_dependents(data, paths, output)

    for path in exclude_paths:
        steps = path.split(".")[:-1]
        final_key = path.split(".")[-1]
        focus_out = output
        for step in steps:
            focus_out = focus_out[step]
        del focus_out[final_key]

    # We need to replace some things because the oapi-codegen
    # creates invalid golang code in the case of large merges by enum it seems
    for path, replacement in replace_with_raw_paths.items():
        steps = path.split(".")[:-1]
        final_key = path.split(".")[-1]
        focus_out = output
        for step in steps:
            focus_out = focus_out[step]
        focus_out[final_key] = replacement
    return output


def fix_anyof_null(data):
    match data:
        case dict():
            for key, value in data.items():
                if isinstance(value, dict) and "anyOf" in value:
                    any_of_conf = value["anyOf"]
                    if {"type": "null"} in any_of_conf:
                        final = any_of_conf[::]
                        final.remove({"type": "null"})
                        if len(final) == 1:
                            data[key] = final[0]
                        else:
                            data[key]["anyOf"] = final
            for item in data.values():
                fix_anyof_null(item)
        case list():
            for item in data:
                fix_anyof_null(item)


PATHS = [
    "openapi",
    "info",
    "paths./api/companies/{company_pk}/opportunities/.post",
    "paths./api/companies/{company_pk}/opportunities/external/.post",
    "paths./api/companies/{parent_pk}/syndis-scans.get",
    "paths./api/companies/{parent_pk}/blobs/upload",
    "paths./api/token/.get",
    "paths./api/integrations/syndis-scan/{scan_name}/config.get",
    "paths./api/integrations/syndis-scan/{scan_name}/logs.post",
    "paths./api/integrations/syndis-scan/{scan_name}/scan.post",
    "paths./api/companies/{company_pk}/opportunities/{opportunity_uid}/",
    "paths./api/companies/{company_pk}/opportunities/v3",
]

EXCLUDE_PATHS = [
    "components.schemas.SyndisScanEntity.properties.created",
    "components.schemas.SyndisScanEntity.properties.updated",
    "components.schemas.SyndisScanEntity.properties.pk",
    "components.schemas.SyndisScanEntity.properties.sk",
    "components.schemas.SyndisScanEntity.properties.entityType",
]

PATHS_REPLACED = {
    "components.schemas.SearchedOpportunitiesResponse.properties.opportunities.items": {
        "type": "object"
    }
}

# test_cases()
if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        prog="SubsetMaker",
        description="Get a subset of a json file based on paths",
    )
    parser.add_argument("filename")  # positional argument

    args = parser.parse_args()

    data = json.loads(open(args.filename).read())
    subset = get_subset(data, PATHS, EXCLUDE_PATHS, PATHS_REPLACED)
    fix_anyof_null(subset)
    print(json.dumps(subset, indent=4))
