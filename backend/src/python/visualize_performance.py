from flask import Flask, jsonify, request
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import io
import base64
from waitress import serve
from flask_cors import CORS

app = Flask(__name__)
CORS(app)  # CORSを有効にする


@app.route("/api/visualize", methods=["POST"])
def visualize():
    try:
        # クライアントから送信されたJSONデータを辞書型として取得
        data = request.json
        # 取得した辞書型データをPandas DataFrameに変換
        df = pd.DataFrame(data)

        # データの可視化
        plt.figure(figsize=(12, 6))  # グラフのサイズを設定
        ax = sns.barplot(
            x="CreatedAt",
            y="CorrectAnswers",
            data=df,
            color="b",
            label="Correct Answers",
        )  # 棒グラフを作成
        ax2 = ax.twinx()  # 2つ目のy軸を作成
        sns.lineplot(
            x="CreatedAt",
            y="PerformanceIndex",
            data=df,
            ax=ax2,
            color="r",
            label="Performance Index",
        )  # 線グラフを作成

        # グラフのタイトルとラベルを設定
        ax.set_title(f"Performance Data for Employee {df['EmpID'][0]}")
        ax.set_xlabel("Date")
        ax.set_ylabel("Correct Answers")
        ax2.set_ylabel("Performance Index")

        # 画像をバイナリデータとしてエンコード
        img = io.BytesIO()
        plt.savefig(img, format="png")
        img.seek(0)
        img_base64 = base64.b64encode(img.getvalue()).decode("utf8")

        # エンコードされた画像データをJSON形式で返す
        return jsonify({"image": img_base64})
    except Exception as e:
        # エラーが発生した場合、エラーメッセージを返す
        return jsonify({"error": str(e)}), 500


if __name__ == "__main__":
    # 本番環境用: Waitressサーバーでアプリを起動
    # serve(app, host='0.0.0.0', port=5001)

    # 開発環境用: Flaskのデバッグサーバーで起動
    app.run(debug=True, host="0.0.0.0", port=5001)
