# Releasing

Releases are automated with GitHub Actions and GoReleaser. There are several manual steps to complete:

1. Decide on the version number by reviewing commits since the last release. The provider versions follow [semantic versioning](https://semver.org/).

1. Start the release action:

   Assuming `upstream` is the remote name pointing to the `canonical/gomaasclient` repository:
   1. First, checkout the master branch and pull the latest changes:

      ```shell
      git switch master
      git pull upstream master
      ```
   1. Verify the latest commit is the same as the latest on the remote master branch.

   1. Create a new tag and push it to the remote repository to trigger the release action:

      ```shell
      git tag vX.Y.Z
      git push upstream tag vX.Y.Z
      ```
   The release action can be viewed under [Actions](https://github.com/canonical/gomaasclient/actions/workflows/release.yml).

1. Publish the release:

   The GitHub Action creates a "draft" release. Go to [Releases](https://github.com/canonical/gomaasclient/releases) to open it, select `edit`, and select `Publish release` if you are happy.

1. Verify the release is published:
   1. Check the release is now the latest published under [Releases](https://github.com/canonical/gomaasclient/releases).
   1. Verify the Go module is available by checking it can be fetched:

      ```shell
      go list -m github.com/canonical/gomaasclient@vX.Y.Z
      ```
   1. Check the package is indexed on [pkg.go.dev](https://pkg.go.dev/github.com/canonical/gomaasclient). This could take some hours to update.
