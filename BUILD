#gazelle:exclude basic
#gazelle:exclude common
#gazelle:exclude pkg4
#gazelle:exclude pkg5

load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library", "go_prefix")

go_prefix("github.com/deepakr199/go")

gazelle(
    name = "gazelle",
)

go_library(
    name = "go_default_library",
    srcs = ["fakemain.go"],
    importpath = "github.com/deepakr199/go",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "go",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
