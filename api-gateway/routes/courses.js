var express = require('express')
var router = express.Router();

var courseHandler = require("./handler/courses")

const verifyToken = require("../middlewares/verifytoken")

router.post('/create', courseHandler.create)
router.get('/getall', courseHandler.getAll)
router.get('/getdetail/:courseId', courseHandler.getDetail)
router.put('/update/:courseId', courseHandler.update)
router.delete('/destroy/:courseId', courseHandler.destroy)

module.exports = router;