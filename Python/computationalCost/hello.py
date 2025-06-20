"""
リストの要素を全てn乗して返す関数をさまざまな形で作成し、Pythonでの計算コストを比較するプログラム
"""

import timeit

import numpy as np
import pandas as pd


# 各手法の実装
def scale_with_list_comprehension(lst, n):
    return [x**n for x in lst]


def scale_with_map(lst, n):
    return list(map(lambda x: x**n, lst))


def scale_with_numpy(lst, n):
    return (np.array(lst) ** n).tolist()


def scale_with_numpy_raw(lst, n):
    arr = np.array(lst)
    return arr * n  # 変換なし


# ベンチマーク関数
def benchmark_list_scaling_methods():
    setup_code = """
from __main__ import scale_with_list_comprehension, scale_with_map, scale_with_numpy, scale_with_numpy_raw
lst = list(range(1000))
n = 5
"""
    test_cases = {
        "List Comprehension": "scale_with_list_comprehension(lst, n)",
        "Map": "scale_with_map(lst, n)",
        "NumPy": "scale_with_numpy(lst, n)",
        "NumPy (Raw)": "scale_with_numpy_raw(lst, n)",
    }

    repeat_counts = [100, 1000, 10000, 100000, 1000000]
    results = []

    for method, stmt in test_cases.items():
        for repeat in repeat_counts:
            exec_time = timeit.timeit(stmt, setup=setup_code, number=repeat)
            results.append(
                {
                    "Method": method,
                    "Repetitions": repeat,
                    "Time (seconds)": round(exec_time, 6),
                }
            )

    return results


results_df = pd.DataFrame(benchmark_list_scaling_methods())
print(results_df.to_string(index=False))  # 表形式できれいに表示

# 小数の表示を整える
pd.set_option("display.float_format", "{:.6f}".format)

# ピボット変換
pivot_df = results_df.pivot(
    index="Method", columns="Repetitions", values="Time (seconds)"
)

# 表示
print(pivot_df.to_string())
