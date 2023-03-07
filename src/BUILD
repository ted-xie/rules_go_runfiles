load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

genrule(
    name = "a_txt",
    outs = ["a.txt"],
    cmd = "echo hello world > $@",
)

go_library(
    name = "a_lib",
    srcs = ["a.go"],
    importpath = "github.com/ted-xie/rules_go_runfiles",
)

go_test(
    name = "a_test",
    srcs = ["a_test.go"],
    data = [":a_txt"],
    embed = [":a_lib"],
)

go_library(
    name = "src",
    srcs = ["a.go"],
    importpath = "github.com/ted-xie/rules_go_runfiles/src",
    visibility = ["//visibility:public"],
)

go_test(
    name = "src_test",
    srcs = ["a_test.go"],
    embed = [":src"],
    deps = ["@io_bazel_rules_go//go/runfiles:runfiles"],
)
