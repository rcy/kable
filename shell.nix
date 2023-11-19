let
  unstable = import (fetchTarball https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz) { };
in
{ nixpkgs ? import <nixpkgs> {} }:
with nixpkgs; mkShell {
  buildInputs = [
    air
    unstable.go_1_21
    golint
    gopls
    sqlite
    flyctl
    golangci-lint
    entr
    google-cloud-sdk
    nodejs
  ];
}
