from flask import Flask
from flask import jsonify

app = Flask(__name__)


@app.route('/say-hello/<name>', methods=['GET'])
def say_hello(name):
    response = {
        'message': f'Hello {name}'
    }
    return jsonify(response)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
