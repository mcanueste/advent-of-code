{
  description = "A Nix-flake-based Go 1.19 development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      goVersion = 19;
      overlays = [(self: super: {go = super."go_1_${toString goVersion}";})];
      pkgs = import nixpkgs {inherit overlays system;};
    in {
      devShells.default = pkgs.mkShellNoCC {
        packages = with pkgs; [
          # go 1.19 (specified by overlay)
          go

          # goimports, godoc, etc.
          gotools

          # https://github.com/golangci/golangci-lint
          golangci-lint
        ];

        shellHook = ''
          ${pkgs.go}/bin/go version
        '';
      };
    });
}
