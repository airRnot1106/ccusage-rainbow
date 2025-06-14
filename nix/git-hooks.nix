{ treefmtEval }:
{
  src = ../.;
  hooks = {
    # Use treefmt for pre-commit formatting (custom to avoid --fail-on-change)
    treefmt-custom = {
      enable = true;
      name = "treefmt";
      entry = "${treefmtEval.config.build.wrapper}/bin/treefmt --no-cache";
      language = "system";
      pass_filenames = false;
      always_run = true;
    };
    # Run golangci-lint on commit
    golangci-lint.enable = true;
    # Run gotest on commit
    gotest.enable = true;
  };
}
