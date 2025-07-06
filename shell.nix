let
  unstable = import (fetchTarball https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz) { };
in
{ nixpkgs ? import <nixpkgs> {} }:
with nixpkgs; mkShell {
  buildInputs = [
    air
    go
    golint
    gopls
    sqlite
    flyctl
    golangci-lint
    entr
    google-cloud-sdk
    nodejs
  ];

  DOCKER_CLI_HINTS="false";
  PGSERVICEFILE="${builtins.toString ./.}/pg-service.conf";
}
