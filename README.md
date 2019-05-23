# yamlyaml

convert yaml with yaml in values to plain yaml


## description

translates something like this:

    ---
    apiVersion: v1
    data:
      file.yaml: "\"key\": \"some value\"\n"

to this:

    ---
    apiVersion: v1
    data:
      file.yaml:
        "key": "some value"


## install

    go install github.com/nordicdyno/yamlyaml

## usage

    ~/go/bin/yamlyaml your-file.yaml

or

    cat your-file.yaml | ~/go/bin/yamlyaml


## ideas of improvents

* preserve original indents
* detect or enable/disable of `---` preambula in output
* add tests
