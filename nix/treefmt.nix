{ pkgs, ... }:
{
  projectRootFile = "flake.nix";
  programs = {
    gofumpt.enable = true;
    nixfmt.enable = true;
  };
}
