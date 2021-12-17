import grpc
import uvicorn
import os
from fastapi import FastAPI, Depends, HTTPException
from fastapi.logger import logger
from datetime import datetime
from pydantic import BaseModel
from typing import List
from meterusage_pb2 import MeterUsageRequest
from meterusage_pb2_grpc import MeterUsageServiceStub

RPC_SERVER_HOSTNAME = os.environ.get("RPC_SERVER_HOSTNAME", "localhost")
RPC_SERVER_PORT = os.environ.get("RPC_SERVER_PORT", 50051)
RPC_SERVER_ADDRESS = f"{RPC_SERVER_HOSTNAME}:{RPC_SERVER_PORT}"

app = FastAPI(title="Meter Usage")


class MeterUsage(BaseModel):
    time: datetime
    usage: float


def get_stub():
    channel = grpc.insecure_channel(RPC_SERVER_ADDRESS)
    stub = MeterUsageServiceStub(channel)
    yield stub


def grpc_get_usages(
    stub: MeterUsageServiceStub, limit: int, offset: int
) -> List[MeterUsage]:
    try:
        response = stub.Get(MeterUsageRequest(limit=limit, offset=offset))
        return [
            MeterUsage(time=datetime.fromtimestamp(usage.time), usage=usage.usage)
            for usage in response.usages
        ]
    except grpc.RpcError as e:
        raise HTTPException(status_code=400, detail=e.details())


@app.get("/api/")
def get_usages(
    limit: int = 100, offset: int = 0, stub=Depends(get_stub)
) -> List[MeterUsage]:
    return grpc_get_usages(stub, limit, offset)


if __name__ == "__main__":
    uvicorn.run(app, port=8000)
