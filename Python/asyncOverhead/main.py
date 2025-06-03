from fastapi import FastAPI

app = FastAPI()


@app.get("/async-cpu-bound")
async def async_cpu_bound():
    # CPUバウンド処理を無駄にasync関数にしている
    total = 0
    for i in range(10_000_000):
        total += i
    return {"total": total}


@app.get("/sync-cpu-bound")
def sync_cpu_bound():
    # 同じCPUバウンド処理だが、通常の関数として実装
    total = 0
    for i in range(10_000_000):
        total += i
    return {"total": total}
