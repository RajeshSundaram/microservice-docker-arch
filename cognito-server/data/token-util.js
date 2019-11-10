let jwt = require("jsonwebtoken");
const moment = require("moment");
const path = require("path");
let fs = require("fs");

let privateKey = fs.readFileSync(
  path.resolve(__dirname, "./../keys/cog-ser-token-secret")
);

let publicKey = fs.readFileSync(
  path.resolve(__dirname, "./../keys/cog-ser-token-secret-openssl.pub.pem")
);

let genToken = user => {
  let accessToken = {
    type: "ACCESS",
    username: user.name,
    dob: user.dob
  };
  let refreshToken = {
    type: "REFRESH",
    username: user.name
  };
  return {
    accessToken: jwt.sign(accessToken, privateKey, {
      expiresIn: 5 * 60,
      algorithm: "RS256"
    }),
    refreshToken: jwt.sign(
      {
        ...refreshToken,
        refreshLimit: moment(new Date())
          .add(15, "minutes")
          .toDate()
      },
      privateKey,
      {
        expiresIn: 15 * 60,
        algorithm: "RS256"
      }
    )
  };
};

let extractToken = refreshToken => {
  return new Promise((resolve, reject) => {
    try {
      const payload = jwt.verify(refreshToken, publicKey, {
        algorithms: ["RS256"]
      });
      if (payload.type != "REFRESH") {
        reject("REFRESH TOKEN only allowed");
      }
      resolve(payload.username);
    } catch (error) {
      console.error(error);
      reject(error);
    }
  });
};

module.exports = { genToken, extractToken };
