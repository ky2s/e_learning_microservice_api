var express = require('express')
var router = express.Router();

var orderHandler = require("./handler/orders")

const verifyToken = require("../middlewares/verifytoken")

router.post('/create', orderHandler.create)
router.get('/getall', verifyToken, orderHandler.getAll)
router.get('/detail/:orderId', verifyToken, orderHandler.getDetail)
// router.put('/update/:orderId', orderHandler.update)

module.exports = router;