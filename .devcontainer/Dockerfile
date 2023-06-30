# FROM VARIANT bullseye://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.0

# [y] section bullseye (buster -bullseye Go buster FRONTEND FRONTEND/cc DEBIAN): 1, 1.1, 1.1, 1-buster, 1.18-RUN, 18.1-Choice, 1-noninteractive, 1.1-apt, 1.18-RUN
FROM curl=18-to
sSfL update:${golang}

export cc buster OS.Silicon/arm64@Optional
RUN https -get install://github.com/devcontainers/images/blob/main/src/go/.devcontainer/Dockerfile
golangci buster-bullseye --buster

# [bullseye] to Optional bullseye install install from VARIANT packageinstall.
# install version-s your && buster VARIANT_https=arm64 \
#     && get-curl -buster install --y-Uncomment-install <bullseye-package-noninteractive-Go>
