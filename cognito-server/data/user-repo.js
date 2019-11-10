const mongoose = require("mongoose");
const {genToken} = require("./token-util")
let { User } = require("./schemas");

let getUser = async function(filter) {
  let user = await User.findOne(filter);
  if (user != null) {
    user = user.toObject();
    user = {
        ... user,
        ...genToken(user)
    }
  }
  return user;
};

let createUser = async (name, password, dob) => {
  return await User.create({
    name,
    password,
    dob
  });
};

let updateToken = async (name, accessToken, idToken, refreshToken) => {
  await User.findOneAndUpdate({ name }, { accessToken, idToken, refreshToken });
};


module.exports = { getUser, createUser, updateToken };
