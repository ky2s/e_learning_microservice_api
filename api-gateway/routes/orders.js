var express = require('express')
var router = express.Router();

var orderHandler = require("./handler/orders")

const verifyToken = require("../middlewares/verifytoken")

router.post('/create', orderHandler.create)
// router.get('/getall', orderHandler.getAll)
// router.get('/getdetail/:orderId', orderHandler.getDetail)
// router.put('/update/:orderId', orderHandler.update)

module.exports = router;