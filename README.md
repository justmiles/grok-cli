# grok-cli

Easily test your grok expressions and parse log files from the CLI

## Installation

Download install the grok binary from the [releases](https://github.com/justmiles/grok-cli/releases) page.

```
# Linux install
curl -sfLo ~/bin/grok https://github.com/justmiles/grok-cli/releases/download/v0.0.2/grok-0.0.2-linux
chmod +x ~/bin/grok
```

## Usage

    Usage:
      grok parse <files> [flags]

    Flags:
      -h, --help                        help for parse
      -m, --multi-line-pattern string   pattern to mark the beginning of a multiline grok
          --named-captures-only         only include named capture groups in returned data (default true)
      -o, --output-type string          output type csv or json (default "json")
      -p, --pattern string              pattern to match
      -d, --patterns-dir stringArray    directory to with additional grok patterns
          --remove-empty-values         do not include empty values in returned data (default true)
          --skip-default-patterns       skip default patterns

## Examples

    grok parse \
      -d ~/grok-patterns \
      --multi-line-pattern '%{TS:timestamp}' \
      --pattern '%{TS:timestamp} %{LOGLEVEL:level}%{SPACE}\[%{THREAD:thread}\] %{JAVACLASS:class}:%{SPACE}%{RAW:details}' \
      --output-type json \
      my-log-files

## Roadmap

- [ ] support multiple output types: json, csv, text
