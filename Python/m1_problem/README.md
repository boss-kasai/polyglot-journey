# N±1問題を意図的に起こして検出したり改善したり。

意図的にN＋1問題は起こして、その検出と改善を実践してみます。

データベースは、useRedisのものを使用します。

```bash:
uv run uvicorn app.main:app --reload --log-level debug
```