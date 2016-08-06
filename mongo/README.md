# MongoDB

* brew install mongodb
* MongoDBの薄い本 http://www.cuspy.org/diary/2012-04-17/the-little-mongodb-book-ja.pdf

# mongod

* mongod -dbpath /usr/local/var/mongodb

# mongo

* サンプルデータ挿入 db.unicorn.insert(...)
* 確認 db.unicorns.find()
* 取得 db.unicorns.find({gender:"m"})