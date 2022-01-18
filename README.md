# sample-specs

The aim of this project is to explore the functionality of Furo: https://furo.pro.

It was created following the approach described in the [Furo Quickstart Guide](https://furo.pro/docs/guides/quickstart).

Please note: `furo install` did not work as described in the Quickstart Guide, returning an error message: `git@github.com: Permission denied (publickey)`. This is caused by a new security restriction of GitHub requiring a registered public key for SSH connections. Follow GitHub's instructions for [generating an SSH key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent) and for [adding the key to your GitHub account](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account). After that, `furo install` is working fine.

As described in the Quickstart Guide, running `furo` will pick the Furo muspecs for a sample type and a sample service from /muspecs/sample and generate Furo long-format specs (/specs/sample), protobuf type and gRPC service specs (/dist/protos/sample), Go language stubs (/dist/pb/sample), a Go package for gRPC-Gateway (/pkg/grpc-gateway/autoregister.go) and a long /dist/env.js file to be used for [Furo Web Components](https://furo.pro/docs/web-components).
