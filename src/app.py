from flask import Flask
from flask import jsonify
app = Flask(__name__)


@app.route('/')
def doRequest():
    data = {
        "serverName": "backend-server",
        "version": "master",
        "success": "true"
    }
    return jsonify(data)


if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", port=80)
