// swift-tools-version:5.5
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "MongoRPC",
    products: [
        .library(name: "MongoRPC", targets: ["MongoRPC"]),
    ],
    dependencies: [
        .package(url: "https://github.com/grpc/grpc-swift.git", from: "1.5.0"),
    ],
    targets: [
        .target(
            name: "MongoRPC",
            dependencies: [.product(name: "GRPC", package: "grpc-swift")],
            path: "mongorpc-swift"
        ),
    ]
)
