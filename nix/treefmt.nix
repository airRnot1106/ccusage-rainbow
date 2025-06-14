{ pkgs, ... }:
{
  projectRootFile = "flake.nix";
  programs.nixfmt.enable = true;

  settings.formatter.go = {
    command = pkgs.writeShellApplication {
      name = "go-formatter";
      runtimeInputs = [
        pkgs.gotools
        pkgs.gofumpt
      ];
      text = ''
        # ファイルをgoimportsで処理してからgofumptで処理
        for file in "$@"; do
          goimports "$file" | gofumpt > "$file.tmp" && mv "$file.tmp" "$file"
        done
      '';
    };
    includes = [ "*.go" ];
  };
}
