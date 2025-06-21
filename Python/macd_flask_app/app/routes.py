from flask import Blueprint, render_template, request

from .utils.macd import calculate_macd_chart

main = Blueprint("main", __name__)


@main.route("/", methods=["GET", "POST"])
def index():
    chart = None
    symbol = ""
    if request.method == "POST":
        symbol = request.form.get("symbol", "").upper()
        chart = calculate_macd_chart(symbol)
    return render_template("index.html", chart=chart, symbol=symbol)
