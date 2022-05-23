load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

_ALL_CONTENT = """\
filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)
"""

http_archive(
    name = "rules_foreign_cc",
    sha256 = "7caed4039f442fceed5b80a7b247d4bebdf9745b05c620811855397b54d73052",
    strip_prefix = "rules_foreign_cc-324dbd13cf2716af4e3979fdde7e7cd6eea8d41c",
    url = "https://github.com/bazelbuild/rules_foreign_cc/archive/324dbd13cf2716af4e3979fdde7e7cd6eea8d41c.tar.gz",
)

load("@rules_foreign_cc//foreign_cc:repositories.bzl", "rules_foreign_cc_dependencies")

rules_foreign_cc_dependencies()

http_archive(
    name = "cc_libvnc",
    build_file_content = _ALL_CONTENT,
    sha256 = "553ddd8b042c77767213586a0e6b44d99729cb533c31cbe94f941c53641660ec",
    strip_prefix = "libvncserver-97fbbd678b2012e64acddd523677bc55a177bc58",
    urls = [
        "https://github.com/LibVNC/libvncserver/archive/97fbbd678b2012e64acddd523677bc55a177bc58.tar.gz",
    ],
)

http_archive(
    name = "cc_libjpeg",
    build_file_content = _ALL_CONTENT,
    sha256 = "2303a6acfb6cc533e0e86e8a9d29f7e6079e118b9de3f96e07a71a11c082fa6a",
    strip_prefix = "jpeg-9d",
    urls = [
        "https://www.ijg.org/files/jpegsrc.v9d.tar.gz",
    ],
)

http_archive(
    name = "cc_libpng",
    build_file_content = _ALL_CONTENT,
    sha256 = "6d59d6a154ccbb772ec11772cb8f8beb0d382b61e7ccc62435bf7311c9f4b210",
    strip_prefix = "libpng-1.6.35",
    urls = [
        "https://github.com/glennrp/libpng/archive/v1.6.35.tar.gz",
    ],
)

http_archive(
    name = "cc_zlib",
    build_file_content = _ALL_CONTENT,
    sha256 = "629380c90a77b964d896ed37163f5c3a34f6e6d897311f1df2a7016355c45eff",
    strip_prefix = "zlib-1.2.11",
    urls = [
        "https://github.com/bazelregistry/zlib/archive/v1.2.11.tar.gz",
    ],
)

http_archive(
    name = "cc_lzo",
    build_file_content = _ALL_CONTENT,
    sha256 = "c0f892943208266f9b6543b3ae308fab6284c5c90e627931446fb49b4221a072",
    strip_prefix = "lzo-2.10",
    urls = [
        "https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.tar.gz",
    ],
)

http_archive(
    name = "cc_cairo",
    build_file_content = _ALL_CONTENT,
    sha256 = "5e7b29b3f113ef870d1e3ecf8adf21f923396401604bda16d44be45e66052331",
    strip_prefix = "cairo-1.16.0",
    urls = [
        "https://mirror.bazel.build/www.cairographics.org/releases/cairo-1.16.0.tar.xz",
        "https://www.cairographics.org/releases/cairo-1.16.0.tar.xz",
    ],
)

http_archive(
    name = "cc_uuid",
    build_file_content = _ALL_CONTENT,
    sha256 = "3cd8b707fd094a642c1d900ff6b9ca5705b4a7aa6e11e7e3dd7029ec4f92fd0d",
    strip_prefix = "uuid-1.7.2",
    patches = ["//patches:uuid-autogen-executable.patch"],  # keep
    urls = [
        "https://github.com/rse/uuid/archive/refs/tags/1.7.2.tar.gz",
    ],
)

http_archive(
    name = "cc_openssl",
    build_file_content = _ALL_CONTENT,
    sha256 = "9384a2b0570dd80358841464677115df785edb941c71211f75076d72fe6b438f",
    strip_prefix = "openssl-1.1.1o",
    urls = [
        "https://www.openssl.org/source/openssl-1.1.1o.tar.gz",
    ],
)

http_archive(
    name = "cc_ssh2",
    build_file_content = _ALL_CONTENT,
    sha256 = "d5fb8bd563305fd1074dda90bd053fb2d29fc4bce048d182f96eaa466dfadafd",
    strip_prefix = "libssh2-1.9.0",
    type = "tar.gz",
    urls = [
        "https://mirror.bazel.build/github.com/libssh2/libssh2/releases/download/libssh2-1.9.0/libssh2-1.9.0.tar.gz",
        "https://github.com/libssh2/libssh2/releases/download/libssh2-1.9.0/libssh2-1.9.0.tar.gz",
    ],
)

http_archive(
    name = "cc_freerdp",
    build_file_content = _ALL_CONTENT,
    sha256 = "2350097b2dc865e54a3e858bce0b13a99711428d397ee51d60cf91ccb56c0415",
    strip_prefix = "FreeRDP-2.7.0",
    urls = [
        "https://github.com/FreeRDP/FreeRDP/archive/refs/tags/2.7.0.tar.gz",
    ],
)

http_archive(
    name = "cc_libgcrypt",
    build_file_content = _ALL_CONTENT,
    sha256 = "0cba2700617b99fc33864a0c16b1fa7fdf9781d9ed3509f5d767178e5fd7b975",
    strip_prefix = "libgcrypt-1.8.6",
    urls = [
        "https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.8.6.tar.bz2",
    ],
)

http_archive(
    name = "cc_guacamole_server",
    build_file_content = _ALL_CONTENT,
    strip_prefix = "guacamole-server-81300052e01f7e544ead82ece1e8398d95e9653e",
    patches = ["//patches:guacamole-disable-lib-checks.patch"],  # keep
    sha256 = "7f01f6c7b33a0ba9c46119a7a7609f2f2ab3274239b45b5416ac3504421cb248",
    urls = [
        "https://github.com/apache/guacamole-server/archive/81300052e01f7e544ead82ece1e8398d95e9653e.tar.gz"
    ]
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "f2dcd210c7095febe54b804bb1cd3a58fe8435a909db2ec04e31542631cf715c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "5982e5463f171da99e3bdaeff8c0f48283a7a5f396ec5282910b9e8a49c0dd7e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.18.1")

gazelle_dependencies()