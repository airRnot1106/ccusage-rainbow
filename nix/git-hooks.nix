{ treefmtEval }:
{
  src = ./.;
  hooks = {
    # Use treefmt for pre-commit formatting
    treefmt = {
      enable = true;
      package = treefmtEval.config.build.wrapper;
    };
    # Run golangci-lint on commit
    golangci-lint.enable = true;
    # Run gotest on commit
    gotest.enable = true;
  };
}
