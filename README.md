## Merge Branch

> Version: v0.0.2

------

## How to use it?

By default `merge-branch` uses the env-variable `GITHUB_REF` as a ref base to create a merge. You only need set the variable `base` with the branch name.


For more information you can check this [link](https://developer.github.com/v3/repos/merging/#merge-a-branch)

```yaml
jobs:
  job-id:
    runs-on: ubuntu-latest
    steps:
      - name: Merge Branch
        uses: Hatzelencio/merge-branch@v0.0.2
        with:
          base: "master"
        env:
          GITHUB_TOKEN: ${{secrets.GITHUB_TOKEN}}
```

If you prefer, you can set the head ref or the commit message to override it.  

```yaml
jobs:
  job-id:
    runs-on: ubuntu-latest
    steps:
      - name: Merge Branch
        uses: Hatzelencio/merge-branch@v0.0.2
        with:
          base: "master"
          head: "develop"
          commitMessage: "Merge pull request #${{ github.payload.pull_request.number }} from {{ github.ref }}"
        env:
          GITHUB_TOKEN: ${{secrets.GITHUB_TOKEN}}
```
