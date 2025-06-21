import base64
import io

import matplotlib

matplotlib.use("Agg")  # GUIなしの描画用バックエンドを使う

import matplotlib.pyplot as plt
import yfinance as yf


def calculate_macd_chart(ticker: str) -> str | None:
    try:
        df = yf.download(ticker, period="6mo", auto_adjust=False)

        if df.empty:
            return None

        # MACD計算
        df["EMA12"] = df["Close"].ewm(span=12, adjust=False).mean()
        df["EMA26"] = df["Close"].ewm(span=26, adjust=False).mean()
        df["MACD"] = df["EMA12"] - df["EMA26"]
        df["Signal"] = df["MACD"].ewm(span=9, adjust=False).mean()
        df["Histogram"] = df["MACD"] - df["Signal"]

        # チャート生成
        fig, ax = plt.subplots(figsize=(10, 4))
        ax.plot(df["MACD"], label="MACD", color="blue")
        ax.plot(df["Signal"], label="Signal", color="red")
        ax.bar(df.index, df["Histogram"], color="gray", alpha=0.5)
        ax.legend()
        plt.tight_layout()

        # PNG → base64変換
        buf = io.BytesIO()
        plt.savefig(buf, format="png")
        buf.seek(0)
        chart_base64 = base64.b64encode(buf.getvalue()).decode("utf-8")
        plt.close(fig)

        return chart_base64

    except Exception as e:
        print(f"Error generating MACD chart for {ticker}: {e}")
        return None
