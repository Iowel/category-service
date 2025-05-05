#!/bin/sh

GRPC_HOST="localhost:44044"
GRPC_METHOD="categoryservice.CategoryService/GetCategoryById"
# GRPC_METHOD="authv1/GetCategoryById"

payload=$(
    cat <<EOF
{
    "id": 1
}
EOF
)

grpcurl -plaintext -emit-defaults \
-rpc-header 'x-app-name:dev' \
-rpc-header 'x-app-version:1' \
-d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}

