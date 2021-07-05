# TEdi (Text Editor)

TEdi is a CLI tool writen in go for extract text and replace it.

## Command extract

**tedi extract [options]**

| option | type | description | default |
|--------|------|-------------|---------|
| --path | string list separated by comma | folder path to scan | `./` |
| --extension | string list separated by comma | file extension to search | `.go` |
| --ignore | string list separated by comma |glob pattern to ignore folders or files | |
| --pattern | string list separated by comma | regular expression to extract texts | `((?:'\|\").*(?:'|\"))` |
| --output-type | string | output type, stdout or csv | `stdout` |