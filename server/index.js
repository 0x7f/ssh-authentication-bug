const crypto = require("crypto");
const ssh2 = require("ssh2");

const { privateKey } = crypto.generateKeyPairSync("rsa", {
  modulusLength: 2048,
  publicKeyEncoding: { type: "pkcs1", format: "pem" },
  privateKeyEncoding: { type: "pkcs1", format: "pem" },
});

new ssh2.Server({ hostKeys: [privateKey] }, (client) => {
  console.log("SSH Client connected");

  client.on("authentication", (ctx) => {
    console.log("authentication", ctx.method, ctx.username);

    switch (ctx.method) {
      case "none":
        console.log("Auth request with method none");
        return ctx.reject(["publickey"]);

      case "publickey":
        console.log(
          `User ${ctx.username} successfully authenticated with public key`
        );
        ctx.accept();
        break;

      default:
        console.log("Unsupported auth method", ctx.method);
        return ctx.reject();
    }
  });

  client.on("error", (err) => {
    console.log("Client error", err);
  });

  client.on("ready", () => {
    console.log("Client authenticated!");
  });

  client.on("end", () => {
    console.log("Client disconnected");
  });
}).listen(8022, "127.0.0.1", () => {
  console.log(`SSH server successfully started on ssh://127.0.0.1:8022`);
});
