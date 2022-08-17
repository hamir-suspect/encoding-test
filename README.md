# Problem
How to send encrypted messages over grpc.
Idea: encode proto into binary data, encode binary data and send base64 of encrypted data, do the oposite steps in server.

# Proof of concept
Proof of concept of interoperability between go and elixir proto encoding used to send some msg encrypted with public key encription.

Go client encodes (marshals) some proto message into binary data, which is then encrypted with public key and sent as string to elixir service.
Elixir service decodes base64 encoding on the string, decrypts using private key, and decodes proto into a pb struct.

Problem: in proto file it is not clearly visible how this encrypted value is composed, so it has to be docummented.
