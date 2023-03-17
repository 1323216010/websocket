import asyncio
import websockets
from fastapi import FastAPI, Request
import uvicorn
from multiprocessing import Process

list = []

async def echo(websocket):
    msg = await websocket.recv()
    await websocket.send(msg)
    list.append(websocket)
    print(list)

async def main():
    async with websockets.serve(echo, "localhost", 2021):
        await asyncio.Future()  # run forever

app = FastAPI()
@app.get("/send")
async def send(request: Request):
    print(list)
    list[int(request.query_params.get("id"))].send(request.query_params.get("msg"))
def ws():
    asyncio.run(main())

def h():
    uvicorn.run(app="main:app", port=8084, reload=True)

def f(name):
    print('hello', name)

if __name__ == '__main__':
    p1 = Process(target=ws)
    p2 = Process(target=h)
    p1.start()
    p2.start()
