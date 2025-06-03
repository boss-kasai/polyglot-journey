import asyncio
import time


# 共通の重い処理（CPUバウンド）
def cpu_bound():
    total = 0
    for i in range(10_000_000):
        total += i
    return total


# 非同期関数として定義（awaitなし）
async def async_cpu_bound():
    total = 0
    for i in range(10_000_000):
        total += i
    return total


# async を同期的に動かすためのラッパー
def run_async(func):
    return asyncio.run(func())


# ベンチマーク関数
def benchmark(func, n=10, label=""):
    start = time.perf_counter()
    for _ in range(n):
        func()
    duration = time.perf_counter() - start
    print(f"{label:<20}: {duration:.4f} sec (avg: {duration / n:.4f} sec)")


if __name__ == "__main__":
    print("CPU-bound function benchmark (pure Python)")
    benchmark(cpu_bound, n=10, label="sync def")

    benchmark(lambda: run_async(async_cpu_bound), n=10, label="async def")
