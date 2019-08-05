# Most Active GitHub Users Counter

This tool, written on golang, allows to fetch GitHub users from a specified location and sort them by the number of contributions.

## Usage

Project is written on Go programming language, so you have to follow general recommendations and guides about development using this language.

To start to clone the repository to your $GOPATH folder.

### Get GitHub Token

For having an ability to run requests to the GitHub API you have to create personal token [here](https://github.com/settings/tokens) and grant `read:org` and `read:user` 


### How to Run, Example

```
go run *.go \
   --token paste-your-token-here \
   --preset worldwide \
   --amount 500 \
   --consider 1000 \
   --output csv \
   --file ./output.csv
```