let mongo = require("mongodb").MongoClient
let {database, host, port, username, password } = require("../config/db-config")
const mongoURI = `mongodb://${username}:${password}@${host}:${port}`

let listAll = async () => {
    let client = await mongo.connect(mongoURI)
    let db = client.db(database)
    let docs = await db.collection('users').find({}).toArray()
    client.close()
    return docs
}


module.exports = {listAll}
