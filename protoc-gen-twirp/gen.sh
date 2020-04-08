#!/usr/bin/env bash

PROTO_SRC_PATH=pb/
IMPORT_MAPPING="common/commo.proto=twirp/protoc-gen-twirp/pb/common"

find pb -name '*.proto' -exec protoc --plugin=protoc-gen-twirp=protoc-gen-twirp \
--proto_path=$PROTO_SRC_PATH \
--twirp_out=prefix=api,M$IMPORT_MAPPING:$PROTO_SRC_PATH \
--go_out=M$IMPORT_MAPPING:$PROTO_SRC_PATH  {} \;

# protoc --plugin=protoc-gen-markdown=protoc-gen-markdown --markdown_out=path_prefix=/vocaldh,M$IMPORT_MAPPING:. hello.proto
