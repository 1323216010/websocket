import WebSocket, { WebSocketServer } from 'ws';
import express from 'express'

const wss = new WebSocketServer({ port: 2021 });

var i = 0
wss.on('connection', function connection(ws) {
  ws.on('error', console.error)
  ws["id"] = i
  i++
  ws.on('message', function message(data) {
    wss.clients.forEach(function each(client) {
      if (client.readyState === WebSocket.OPEN) {
        client.send(String(data));
      }
    });
  });
});

var app = express()

app.get('/send', function (req, res) {

  wss.clients.forEach(function each(client) {
    if (client.readyState === WebSocket.OPEN && client.id == req.query.id) {
      client.send(req.query.msg);
    }
  })
  
})

var server = app.listen(8081, function () {
  var port = server.address().port
  console.log("访问端口为%s", port)
})