load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library", "go_prefix")

go_prefix("github.com/deepakr199/go/app")

gazelle(
    name = "gazelle",
)

go_library(
    name = "go_default_library",
    srcs = ["fakemain.go"],
    importpath = "github.com/deepakr199/go/app",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "app",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
