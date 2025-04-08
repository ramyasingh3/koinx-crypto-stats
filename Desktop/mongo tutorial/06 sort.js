employees> db.inventory.find().limit(1) //shows first 1 document
db.inventory.find().skip(1)  //skips 1st document
db.inventory.find().sort({qty:1}) //ascending order of qty
db.inventory.find().sort({qty:-1}) //descending order of qty


//pagination
page 8-16

db.invnetory.find().skip(8).limit(8)