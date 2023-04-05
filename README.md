# 4chan General [![Go](https://github.com/fluteds/4ch-gen/actions/workflows/go.yml/badge.svg)](https://github.com/fluteds/4ch-gen/actions/workflows/go.yml)

Self-updating bot running on a 6 hourly CRON job to match board generals to it's correct thread location. Forked from [Bingosss420.](https://github.com/bingsoo420/4ch-general)

Example

```json
{
  "g": {
    "wdg": "https://boards.4chan.org/g/thread/91798719"
  }
}
```

The access patterns will always remain as `board.general`, currently the only boards being scraped are:

- fit
- g
- biz
- ck
- vg
- vr

Which can be modified in `main.go` file.

## Usage

### As a bash script

```sh
curl https://raw.githubusercontent.com/fluteds/4ch-gen/master/output/mappings.json | jq .g.wdg
```

### As a simple JavaScript fetch request

```js
fetch(
  "https://raw.githubusercontent.com/fluteds/4ch-gen/master/output/mappings.json"
)
  .then((r) => r.json())
  .then((r) => {
    if (r.g?.wdg) {
      window.location.href = r.g.wdg;
    }
  });
```

The source will refresh every 6 hours. Files may be updated to be granular later as `output/g.json` or `output/fit.json` if the 6 hours interval was found to be too little for certain boards.

### As a program

```sh
git clone
cd 4ch-general
go run ./
```

See the output JSON file in `output/mappings.json`

### As a fork

- Set your repo's workflow settings to have read/write permissions. `Repo / Settings / Actions / General / Workflow Permissions / Read and Write Permissions`

- Set up a fine grained personal access token in GitHub by visiting `Profile / Settings / Developer settings / Personal access tokens / Fine-grained tokens / Generate new token`
  
- Set the token's permission to read/write under `Contents`. Copy the access token output and put that into a secret called `PAT`. `Repo / Settings / Secrets and variables / Actions / New repository secret`

- Add another secret with the name being `EMAIL` and the value being your GitHub account email address. This is to sign off commits.
