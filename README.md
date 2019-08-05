[Project homepage](https://commits.top/)

# Most Active GitHub Users Counter

This CLI tool queries the GitHub GraphQL API for users and ranks them according to number of contributions. Several preset locations are provided.

**GitHub Token**

In order to make requests against the GitHub API one needs an access token, which can be created [here](https://github.com/settings/tokens). The token needs `read:org` and `read:user` permissions.

**Example usage (dev environment):**

```
go run *.go \
   --token paste-your-token-here \
   --preset worldwide \
   --amount 500 \
   --consider 1000 \
   --output csv \
   --file ./output.csv
```

## Contribution

Contributions are accepted. Please provide any input as pull request. As a hobby project, my time is limited, but PRs and issues are addressed regularly.

*Please use the provided precommit hooks and run `go fmt`, `go vet` and `go lint` liberally.*
