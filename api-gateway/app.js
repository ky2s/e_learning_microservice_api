require("dotenv").config()
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');
var mediaRouter = require('./routes/media');
var coursesRouter = require('./routes/courses');
var refreshTokenRouter = require("./routes/refreshToken")

var app = express()

app.use(logger('dev'));
app.use(express.json({limit : "50mb"}));
app.use(express.urlencoded({ extended: false, limit : "50mb" }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/refresh_token', refreshTokenRouter)

app.use('/users', usersRouter);
app.use('/media', mediaRouter);
app.use('/courses', coursesRouter);

var listener = app.listen(8888, function(){
    console.log('Listening on port ' + listener.address().port); //Listening on port 8888
});

module.exports = app;
