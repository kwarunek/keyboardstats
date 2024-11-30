#! /usr/bin/env python3

import sys
import re

REGEX_C_COMMENT = re.compile(r"/\*.*\*/")

CODES_TO_GET = [
    "KEY_",
    "EV_"
]

GO_TEMPLATE = """
package codes

%s


var KeyCodeToName = map[int]string{
%s
}
"""


def main(file_name: str):
    map = {}
    with open(file_name, "r") as f:
        for raw_line in f:
            line = raw_line.strip()
            if line.startswith("#define"):
                line = REGEX_C_COMMENT.sub("", line)
                try:
                    _, name, value = line.split()
                except ValueError:
                    continue
                for code in CODES_TO_GET:
                    if name.startswith(code) and not value.startswith(code):
                        map[name] = value

    consts = []
    namesMap = []
    for name, value in map.items():
        consts.append(f"const {name} = {value}")
        if name.startswith("KEY_"):
            namesMap.append(f"\t{value}: \"{name}\",")

    print(GO_TEMPLATE % (
        "\n".join(consts),
        "\n".join(namesMap)
    ))


if __name__ == "__main__":
    if len(sys.argv) != 2:
        fn = "/tmp/input-event-codes.h"
    else:
        fn = sys.argv[1]
    main(fn)
