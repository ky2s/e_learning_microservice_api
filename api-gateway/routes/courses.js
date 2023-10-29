var express = require('express')
var router = express.Router();

const verifyToken = require("../middlewares/verifytoken")

var courseHandler = require("./handler/courses")

router.post('/create', verifyToken, courseHandler.create)
router.get('/getall', verifyToken,courseHandler.getAll)
router.get('/getdetail/:courseId', verifyToken, courseHandler.getDetail)
router.put('/update/:courseId', verifyToken, courseHandler.update)
router.delete('/destroy/:courseId', verifyToken, courseHandler.destroy)

module.exports = router;