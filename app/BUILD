load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library")

gazelle(
    name = "gazelle",
    prefix = "github.com/deepakr199/go/app",
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
