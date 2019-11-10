var express = require("express");
var userRepo = require("./../data/user-repo");
var tokenUtil = require("./../data/token-util");
var moment = require("moment");
var router = express.Router();

router.get("/health", (req, res, next) => {
  res.status(200).send("OK");
});

router.post("/signin", async function(req, res, next) {
  let { username, password } = req.body;
  let user = await userRepo.getUser({ name: username, password });
  if (user.dob != undefined) {
    user.dob = moment(user.dob).format("YYYY/MM/DD");
    user.password = undefined;
  }
  res.status(200).json(user);
});

router.post("/signup", async function(req, res, next) {
  let { username, password, dob, firstName, lastName } = req.body;
  let user = await userRepo.getUser({ name: username });
  if (user != null) {
    res.status(400).send("User Already Exists");
    return;
  }
  if (dob != undefined && moment(dob).isValid) {
    dob = moment(dob).format("YYYY-MM-DD");
  } else {
    dob = null;
  }
  user = await userRepo.createUser(username, password, dob);
  res.status(200).json({ id: user._id });
});

router.get("/token", async function(req, res, next) {
  let user, username, error;
  let refreshToken = req.header("x-refresh");
  try {
    if (req.query.type != undefined && req.query.type === "USPWD") {
      username = req.header("username");
      let password = req.header("password");
      user = await userRepo.getUser({ name: username, password });
    } else if (refreshToken != undefined) {
      username = await tokenUtil.extractToken(refreshToken);
      user = await userRepo.getUser({ name: username});
    }
    if (user != null && user != undefined) {
      let tokens = tokenUtil.genToken(user);
      await userRepo.updateToken(
        username,
        tokens.accessToken,
        tokens.idToken,
        tokens.refreshToken
      );
      res.status(200).json(tokens);
    } else if (!error) {
      throw new Error("Credintials missing");
    }
  } catch (err) {
    res.status(400).send(err);
  }
});

module.exports = router;
