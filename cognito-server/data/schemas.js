let mongoose = require("mongoose");
let Schema = mongoose.Schema;
let {database, host, port, username, password } = require("../config/db-config")
const mongoURI = `mongodb://${username}:${password}@${host}:${port}`
// const mongoURI = `mongodb://${username}:${password}@${host}:${port}/${database}`

mongoose.connect(
  mongoURI, {useNewUrlParser: true, useUnifiedTopology: true}).then(
    connection => {
      // console.log(connection)
    })
  .catch(console.error)

let userSchema = new Schema({
    name: String,
    password: String,
    dob: Date,
    accessToken: String,
    refreshToken: String
  })

let UserModel = mongoose.model(
  "users",
  userSchema
);

module.exports = { User: UserModel };
