# TEdi (Text Editor)

TEdi is a CLI tool writen in go for extract text and replace it.

## Command extract

**tedi extract [options]**

Example to extract all texts between two `§`:
`tedi extract -path ./ --extension .go --ignore .git/*,.bin/* --delimiter § --output-type csv`

| option | short | type | description | default |
|--------|-------|------|-------------|---------|
| --path | -p | string | folder path to scan | `./` |
| --extension | -e | string list separated by comma | file extension to search | `.go` |
| --ignore | -i | string list separated by comma |glob pattern to ignore folders or files |  |
| --delimiter | -d | string | delimiter character | `§` |
| --output-type | -o | string | output type, stdout or csv | `stdout` |


## Command replace

**tedi replace [options]**

Example to replace all texts between two `§` with the previously generated and completed csv:
`tedi replace -path ./ --delimiter § --csv ./tedi-result.csv`

| option | short | type | description | default |
|--------|-------|------|-------------|---------|
| --path | -p | string | folder path to scan | `./` |
| --delimiter | -d | string | delimiter character | `§` |
| --csv | -c | string | csv path |  |

## CSV format

| column 1 | column 2 | column 3 |
|----------|----------|----------|
| path to file | original text | new text |

**CSV exemple**

```
file,text,replace
cmd/root.go,tedi,TEdi
```