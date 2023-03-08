import express from 'express'
import arr from './json/poet300.json' assert {type: 'json'}
import { WebSocketServer } from 'ws';

var app = express()

app.get('/poet', function (req, res) {
    let poet = []
    for (let i = 0; i < arr.length; i++) {
        if (arr[i].author === req.query.author) {
            poet.push(arr[i])
        }
    }
    res.setHeader('Content-Type', 'application/json; charset=utf8');
    res.end(JSON.stringify(poet))
})

var server = app.listen(8081, function () {
    var port = server.address().port
    console.log("访问端口为%s", port)
})

const wss = new WebSocketServer({ port: 2020 });

wss.on('connection', function connection(ws) {
  ws.on('error', console.error);

  ws.on('message', function message(data) {
    console.log('received: %s', data);
  });

  ws.send('something');
});
