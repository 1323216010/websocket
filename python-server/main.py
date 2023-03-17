import asyncio
import websockets
from fastapi import FastAPI, Request
import uvicorn

list = []

async def echo(websocket):
    msg = await websocket.recv()
    await websocket.send(msg)
    list.append(websocket)

async def main():
    async with websockets.serve(echo, "localhost", 2021):
        await asyncio.Future()  # run forever

asyncio.run(main())

app = FastAPI()
@app.get("/send")
async def send(request: Request):
    list[request.query_params.get("id")].send(request.query_params.get("msg"))

if __name__ == "__main__":
    uvicorn.run(app="main:app", port=8084, reload=True)