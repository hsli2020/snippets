# db commands

## create
    use <db name>
    The above will use the database of the provided name if it exists, and create it if it doesn't

## see current db
    db

## see all db
    show dbs

### insert document
    db.<collection name>.insert({"name":"McLeod"})
    db.dogs.insert({"name":"toby"})

### view documents
    db.<collection name>.find()
    db.cats.insert({"firstname":"coco"})
    db.<collection name>.find().pretty()
    db.cats.find().pretty()

### view collections
    show collections

## drop db
    db.dropDatabase()

# collection commands

### create implicitly
    db.<collection name>.insert({"name":"McLeod"})

### create explicitly
    db.createCollection(<name>, {<optional options>})

#### optional options
| option | type   | description |
| ------ | ------ | ----------- |
| capped | bool   | caps the size |
| size   | number | sets size of cap in bytes |
| max    | bool   | maximum number of documents allowed in capped collection |

#### examples
    db.createCollection("customers")
    db.createCollection("crs",{capped:true, size:65536,max:1000000})

### view collections
    show collections

### drop
    db.<collection name>.drop()

# document commands

### insert
    db.<collection name>.insert({document})

### insert multiple
    db.<collection name>.insert(< [{document}, {document}, ..., {document}] >)

#### example
    use playroom
    db
    show dbs
    db.crayons.insert([
        { "hex": "#EFDECD", "name": "Almond", "rgb": "(239, 222, 205)" }, 
        { "hex": "#FDD9B5", "name": "Apricot", "rgb": "(253, 217, 181)" }, 
        { "hex": "#87A96B", "name": "Asparagus", "rgb": "(135, 169, 107)" }, 
        { "hex": "#FAE7B5", "name": "Banana Mania", "rgb": "(250, 231, 181)" }
    ])

    show collections
    db.crayons.find()
    db.crayons.drop()
    db.dropDatabase()

# find (aka, query)

### setup

    use store
    db
    show dbs
    db.customers.insert([
        {"role":"double-zero","name": "Bond","age": 32},
        {"role":"citizen","name": "Moneypenny","age":32},
        {"role":"citizen","name": "M","age":57},
        {"role":"citizen","name": "Dr. No","age":52}
    ])

### find
    db.<collection name>.find()
    db.customers.find()

### find one
    db.<collection name>.findOne()
    db.customers.findOne()

### find specific
    db.customers.find({"name":"Bond"})
    db.customers.find({name:"Bond"})

    You can do it either way: `"name" or name`.
    JSON specification is to enclose name (object name-value pair) in double qoutes

### and
    db.customers.find({$and: [{name:"Bond"}, {age:32}]})
    db.customers.find({$and: [{name:"Bond"}, {age:{$lt:20}}]})
    db.customers.find({$and: [{name:"Bond"}, {age:{$gt:20}}]})

### or
    db.customers.find({$or: [{name:"Bond"}, {age:67}]})
    db.customers.find({$or: [{name:"Bond"}, {age:{$lt:20}}]})
    db.customers.find({$or: [{name:"Bond"}, {age:{$gt:32}}]})

### and or
    db.customers.find({role:"citizen"})
    db.customers.find({age:52})
    db.customers.find({$and: [{role:"citizen"}, {age:52}]})
    db.customers.find({$or: [{role:"citizen"}, {age:52}]})
    db.customers.find({$or: [{role:"citizen"}, {age:52}, {name:"Bond"}]})

    db.customers.find({$or:[
        { $and : [ { role : "citizen" }, { age : 32 } ] },
        { $and : [ { role : "citizen" }, { age : 67 } ] }
    ]})

### regex
    db.customers.find({name: {$regex: '^M'}})

### operators

| operator            | syntax             | example |
| ------------------- | ------------------ | ------- |
| equality            | {key:value}        | db.customers.find({"name":"Bond"}).pretty() | 
| less than           | {key:{$lt:value}}  | db.customers.find({"age":{$lt:20}}).pretty() | 
| less than equals    | {key:{$lte:value}} | db.customers.find({"age":{$lte:20}}).pretty() | 
| greater than        | {key:{$gt:value}}  | db.customers.find({"age":{$gt:20}}).pretty() | 
| greater than equals | {key:{$gte:value}} | db.customers.find({"age":{$gte:20}}).pretty() | 
| not equals          | {key:{$ne:value}}  | db.customers.find({"age":{$ne:20}}).pretty() | 

# update & save

update will update a record
save will overwrite a record

### update
    db.<collection name>.update(<selection criteria>, <update data>, <optional options>)

example
    db.customers.find()

gives this data
    { "_id" : ObjectId("5891221756867ebff44cc885"), "role" : "double-zero", "name" : "Bond", "age" : 32 }
    { "_id" : ObjectId("5891221756867ebff44cc886"), "role" : "citizen", "name" : "Moneypenny", "age" : 32 }
    { "_id" : ObjectId("5891221756867ebff44cc887"), "role" : "citizen", "name" : "Q", "age" : 67 }
    { "_id" : ObjectId("5891221756867ebff44cc888"), "role" : "citizen", "name" : "M", "age" : 57 }
    { "_id" : ObjectId("5891221756867ebff44cc889"), "role" : "citizen", "name" : "Dr. No", "age" : 52 }

update like this
    db.customers.update({_id:ObjectId("5891221756867ebff44cc886")},{$set:{role:"double-zero"}})
    db.customers.update({name:"Moneypenny"},{$set:{role:"double-zero"}})
    db.customers.update({name:"Moneypenny"},{$set:{role:"citizen", name: "Miss Moneypenny"}})
    db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}})
    db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}}, {multi:true})
    db.customers.update({},{$set:{role:"citizen"}}, {multi:true})

### save

    db.customers.save({"role":"villain","name":"Jaws","age":43})
    db.customers.save({"_id":ObjectId("5891221756867ebff44cc889"),"role":"villain","name":"Goldfinger","age":77})
    db.customers.save({"_id":ObjectId("5893888012acb8ada532a8e4"),"role":"villain","name":"PussyGalore","age":31})

# remove document

    db.<collection name>.remove(<selection criteria>)
    db.customers.remove({role:"double-zero"})
    db.customers.remove({role:"villain"})

    removes all it matches

### remove only 1
    db.customers.remove({role:"citizen"},1)

### remove
    db.customers.remove({role:"citizen"})

### put documents back
    db.customers.insert([
        {"role":"double-zero","name": "Bond","age": 32},
        {"role":"citizen","name": "Moneypenny","age":32},
        {"role":"citizen","name": "M","age":57},
        {"role":"citizen","name": "Dr. No","age":52}
    ])

### remove all
    db.customers.remove({})

### put documents back
    db.customers.insert([
        {"role":"double-zero","name": "Bond","age": 32},
        {"role":"citizen","name": "Moneypenny","age":32},
        {"role":"citizen","name": "Q","age":67},
        {"role":"citizen","name": "M","age":57},
        {"role":"citizen","name": "Dr. No","age":52}
    ])

# projection
    Retrieving part of a document; only some of the fields.

    db.<collection name>.find(<selection criteria>,<list of fields with toggle 0 or 1>)
    db.customers.find({}, {_id:0,name:1,})

        _id is displayed by default; turn off with 0

    db.customers.find({}, {_id:0,name:1,age:1})
    db.customers.find({age:{$gt:32}}, {_id:0,name:1,age:1})

# limit

### setup
    db.crayons.insert([
        { "hex": "#EFDECD", "name": "Almond", "rgb": "(239, 222, 205)" }, 
        { "hex": "#FDD9B5", "name": "Apricot", "rgb": "(253, 217, 181)" }, 
        { "hex": "#87A96B", "name": "Asparagus", "rgb": "(135, 169, 107)" }, 
        { "hex": "#FAE7B5", "name": "Banana Mania", "rgb": "(250, 231, 181)" }
    ])

### limit
    db.<collection name>.find(<selection criteria>).limit(n)
    db.crayons.find().limit(3)
    db.customers.find({age:{$gt:32}}, {_id:0,name:1,age:1}).limit(2)

# sort
    Run **setup** below first

    db.<collection name>.find().sort(<field to sort on>:<1 for ascend, -1 descend>)

    db.oscars.find().limit(10)
    db.oscars.find({},{_id:0,year:1,title:1}).limit(10)
    db.oscars.find({},{_id:0,year:1,title:1}).limit(10).sort({title:1})
    db.oscars.find({},{_id:0,year:1,title:1}).sort({title:1}).limit(10)
    db.oscars.find({},{_id:0,year:1,title:1}).limit(10).sort({title:-1})
    db.oscars.find({releaseYear:{$gt:1970}},{_id:0,year:1,title:1}).limit(10).sort({title:1})
    db.oscars.find({releaseYear:{$gt:1980}},{_id:0,year:1,title:1})

### setup
    db.oscars.insert([
        { "year": "1927", "title": "Wings", "imdbId": "tt0018578", ... },
        { "year": "1929", "title": "The Broadway Melody", "imdbId": "tt0019729", ... },
        { "year": "1930", "title": "All Quiet on the Western Front", "imdbId": "tt0020629", ... },
        ...
        { "year": "2010", "title": "The King's Speech", "imdbId": "tt1504320", ... },
        { "year": "2011", "title": "The Artist", "imdbId": "tt1655442", ... },
        { "year": "2012", "title": "Argo", "imdbId": "tt1024648", ... }
    ])

# create index
    db.<collection name>.createIndex({<field to index>:<1 for ascend, -1 descend>})
    db.oscars.createIndex({title:1})
    db.oscars.getIndexes()

# aggregate

Aggregations operations process data records and return computed results. Aggregation
operations group values from multiple documents together, and can perform a variety of
operations on the grouped data to return a single result. MongoDB provides three ways to
perform aggregation: the aggregation pipeline, the map-reduce function, and single purpose
aggregation methods.

## single purpose aggregation

    db.collection.distinct(field, query, options)

| Parameter | Description |
| --------- | ----------- | 
| field     | The field for which to return distinct values.
| query     | A query that specifies the documents from which to retrieve the distinct values.
| options   | Optional. A document that specifies the options. See Options.

#### examples - count()
    db.oscars.count()
    db.oscars.find().count()

    db.customers.find({role:"citizen"}).count()
    db.customers.find({$or: [{name:"Bond"}, {age:{$gt:32}}]}).count()

#### examples - distinct() - setup
    db.inventory.insert([
        { "_id": 1, "dept": "A", "item": { "sku": "111", "color": "red" }, "sizes": [ "S", "M" ] },
        { "_id": 2, "dept": "A", "item": { "sku": "111", "color": "blue" }, "sizes": [ "M", "L" ] },
        { "_id": 3, "dept": "B", "item": { "sku": "222", "color": "blue" }, "sizes": "S" },
        { "_id": 4, "dept": "A", "item": { "sku": "333", "color": "black" }, "sizes": [ "S" ] }
    ])

#### examples - distinct()
    db.inventory.distinct( "dept" )
    db.inventory.distinct( "item.sku" )
    db.inventory.distinct( "sizes" )

## aggregation pipeline
    db.<collection name>.aggregate([{<match, sort, geoNear>},{<group>}])

MongoDBâ€™s aggregation framework is modeled on the concept of data processing pipelines.
Documents enter a multi-stage pipeline that transforms the documents into an aggregated
result.

The most basic pipeline stages provide filters that operate like queries and document
transformations that modify the form of the output document.

Other pipeline operations provide tools for grouping and sorting documents by specific
field or fields as well as tools for aggregating the contents of arrays, including arrays
of documents. In addition, pipeline stages can use operators for tasks such as calculating
the average or concatenating a string.

The pipeline provides efficient data aggregation using native operations within MongoDB,
and is the preferred method for data aggregation in MongoDB.

#### example - setup
    db.orders.insert([
        {"cust_id":"A123","amount":500,"status":"A"},
        {"cust_id":"A123","amount":250,"status":"A"},
        {"cust_id":"B212","amount":200,"status":"A"},
        {"cust_id":"A123","amount":300,"status":"D"}
    ])

#### example
    db.orders.aggregate([
        {$match:{status:"A"}},
        {$group:{_id: "$cust_id",total: {$sum:"$amount"}}}
    ])

## create admin super user
    use admin
    db.createUser(
        {
            user: "jamesbond",
            pwd: "moneypennyrocks007sworld",
            roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
        }
    )

#### see current user
    db.runCommand({connectionStatus : 1})

## create regular user
    Give this user readwrite permissions on the ```store``` db.

    db.createUser(
        {
            user: "bond",
            pwd: "moneypenny007",
            roles: [ { role: "readWrite", db: "store" } ]
        }
    )

#### see current user
    db.runCommand({connectionStatus : 1})

#### exit mongo & then start again with auth enabled
    mongod --auth
    mongo -u "bond" -p "moneypenny007" --authenticationDatabase "store"
    mongo -u "jamesbond" -p "moneypennyrocks007sworld" --authenticationDatabase "admin"

#### test
    use store
    show collections
    db.customers.find()
    db.customers.insert({"role" : "double-zero", "name" : "Elon Musk", "age" : 47 })

#### test
    launch a new terminal window
    mongo
    should be unauthorized:
    show collections

#### drop user
    db.dropUser("<user name>")
