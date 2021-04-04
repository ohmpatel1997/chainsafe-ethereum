# node:alpine will be our base image to create this image
FROM node:alpine
# Set the /app directory as working directory
WORKDIR /app

# Install ganache-cli globally
RUN npm install -g ganache-cli

CMD ["ganache-cli", "-h", "0.0.0.0", "--mnemonic", "this is the sample assignment for chainsafe"]
