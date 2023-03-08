import { WebSocketServer } from 'ws';

const wss = new WebSocketServer({ port: 2020 });

wss.on('connection', function connection(ws) {
  ws.on('error', console.error);

  ws.on('message', function message(data) {
    ws.send(String(data));
  });
});
