# MPC using Shamir's Secret Sharing

This project implements a secure Multi-Party Computation (MPC) protocol based on Shamir's Secret Sharing scheme using gRPC for communication. The goal of this project is to allow multiple parties to perform computations on encrypted data without revealing the data itself.

## Overview

1. **Splitting the secret into shares**: The user splits their secret into `n` shares with a threshold of `k` using the `SplitSecret` function in the `secret_sharing.go` file.
2. **Distributing shares to nodes**: The user distributes the generated shares across multiple nodes. Each node will get one or more shares.
3. **MPC server implementation**: The `Server.go` file contains the implementation of the MPC server. It contains the `ComputeSum` function that will be called by the nodes to reconstruct the secret from the shares.
4. **Nodes send shares to the MPC server**: When it's time to perform the MPC computation, the nodes send their shares to the MPC server using the gRPC client in the `mpc_grpc.pb.go` file.
5. **Reconstructing the secret and performing computation**: Inside the `ComputeSum` function in `Server.go`, the server receives the shares, pads them with leading zeros if necessary, and then uses the `shamir.Combine` function to combine the shares and reconstruct the secret.

## Usage

To use this project, follow these steps:

1. Split the user's secret into `n` shares with a threshold of `k` using the `SplitSecret` function in `secret_sharing.go`.
2. Distribute the generated shares to multiple nodes.
3. Run the MPC server using the `RunServer` function in `Server.go`.
4. When it's time to perform the MPC computation, have each node send their shares to the MPC server using the gRPC client in `mpc_grpc.pb.go`.
5. The MPC server will reconstruct the secret and perform the desired computation (e.g., comparison) on the reconstructed secrets while keeping them encrypted.

## Security Considerations

To ensure the privacy and security of the data during computation, follow these best practices:

1. **Secure channels**: Use encrypted and authenticated communication channels like TLS to protect data in transit between nodes and the MPC server.
2. **Least privilege principle**: Limit access to the shares and the MPC server to only the necessary parties and systems.
3. **Secure storage**: Store the shares securely on each node, using encryption and access controls to prevent unauthorized access.
4. **Secure Multi-Party Computation (SMPC)**: Use advanced cryptographic techniques like Secure Multi-Party Computation (SMPC) to perform computations on encrypted data without revealing the actual data to other parties.