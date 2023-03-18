import asyncio
import multiprocessing

import websockets
from fastapi import FastAPI, Request
import uvicorn

list = []

async def echo(websocket):
    msg = await websocket.recv()
    await websocket.send(msg)
    list.append(websocket)

async def ws(list):
    async with websockets.serve(echo, "localhost", 2021):
        await asyncio.Future()  # run forever

app = FastAPI()
@app.get("/send")
async def send(request: Request):
    list[int(request.query_params.get("id"))].send(request.query_params.get("msg"))

def f(list):
    asyncio.run(ws(list))

def g(list):
    uvicorn.run(app="main:app", port=8084, reload=True)

if __name__ == '__main__':
    list1 = multiprocessing.Array('i', list)
    print(list1[:])
    p1 = multiprocessing.Process(target=f, args=(list1,))
    p2 = multiprocessing.Process(target=g, args=(list1,))
    p1.start()
    p2.start()
    p1.join()
    p2.join()
