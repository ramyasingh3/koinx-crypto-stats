db.inventory.find()   //fetch all queries
db.inventory.find({qty:85})
db.inventory.find( { tags: { $in: [ "gray", "red" ] } } )

AND
db.inventory.find( { status: "A", qty: { $lt: 30 } } )

OR 
db.inventory.find( { $or: [ { status: "A" }, { qty: { $lt: 30 } } ] } )

db.inventory.findOne( { status: "A", qty: { $lt: 30 } } )