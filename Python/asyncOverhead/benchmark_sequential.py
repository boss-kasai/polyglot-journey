import time

import requests


def benchmark_sequential(endpoint: str, n: int = 100):
    print(f"\nBenchmarking: {endpoint} with {n} sequential requests")
    start = time.time()
    for i in range(n):
        r = requests.get(endpoint)
        r.raise_for_status()
    total = time.time() - start
    print(f"Total time: {total:.2f} sec")
    print(f"Average per request: {total / n:.4f} sec")


if __name__ == "__main__":
    benchmark_sequential("http://localhost:8000/async-cpu-bound", 100)
    benchmark_sequential("http://localhost:8000/sync-cpu-bound", 100)
