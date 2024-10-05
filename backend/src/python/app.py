from flask import Flask, jsonify, request
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import io
import base64
from waitress import serve  # ここでwaitressのserveをインポート
from flask_cors import CORS  # CORSを有効にするためのモジュールをインポート

app = Flask(__name__)
CORS(app)  # CORSを有効にする

@app.route('/api/visualize', methods=['POST'])
def visualize():
    try:
        # クライアントから送信されたJSONデータを辞書型として取得
        data = request.json

        # 取得した辞書型データをPandas DataFrameに変換し、データの可視化や分析が容易にする
        df = pd.DataFrame(data)

        # データの可視化
        plt.figure(figsize=(10, 6))  # グラフのサイズを設定：インチ
        sns.barplot(x='employee_id', y='score_indicator', data=df)  # Seabornを使用して棒グラフを作成
        plt.title('Employee Scores')  # グラフのタイトルを設定

        # 画像をバイナリデータとしてエンコード
        img = io.BytesIO()  # バッファを作成
        plt.savefig(img, format='png')  # 画像をバッファに保存
        img.seek(0)  # バッファの先頭に移動することが推奨される
        img_base64 = base64.b64encode(img.getvalue()).decode('utf8')  # 画像をBase64エンコード

        # エンコードされた画像データをJSON形式で返す
        return jsonify({'image': img_base64})
    except Exception as e:
        # エラーが発生した場合、エラーメッセージを返す
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    # 本番環境用のコード：本番環境では、waitressサーバーを使用
    # serve(app, host='0.0.0.0', port=5001)  # Waitressサーバーを使用してアプリケーションを起動

    # 開発環境用のコード：開発環境では、Flaskの組み込み開発サーバーを使用
    # デバッグモード、ホットリロード
    app.run(debug=True, host='0.0.0.0', port=5001)

    # サーバ起動時コマンド : python app.py
