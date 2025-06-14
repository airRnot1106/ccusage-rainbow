{
  description = "ccusage-rainbow - CLI tool for ASCII art text display";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      treefmt-nix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        treefmtEval = treefmt-nix.lib.evalModule pkgs ./nix/treefmt.nix;
      in
      {
        formatter = treefmtEval.config.build.wrapper;
        checks.formatting = treefmtEval.config.build.check self;

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            golangci-lint
            gofumpt
            treefmtEval.config.build.wrapper
          ];

          shellHook = ''
            echo "🚀 Go development environment loaded!"
            echo "Available commands:"
            echo "  nix run - Run the CLI tool"
            echo "  golangci-lint run - Run linter"
            echo "  nix fmt - Format code"
          '';
        };

        packages.default = pkgs.buildGoModule {
          pname = "ccusage-rainbow";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-LWY1Tnh4iyNAV7dNjlKdT9IwPJRN25HkEAGSkQIRe9I=";
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/ccusage-rainbow";
        };
      }
    );
}
