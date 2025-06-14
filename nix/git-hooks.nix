{ treefmtEval }:
{
  src = ../.;
  hooks = {
    treefmt = {
      enable = true;
      package = treefmtEval.config.build.wrapper;
    };
    golangci-lint.enable = true;
    gotest.enable = true;
  };
}
