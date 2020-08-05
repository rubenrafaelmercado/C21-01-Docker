from flask import Flask, request
from flask import jsonify


app = Flask(__name__)


alumnos = [
    {'id': 1, 'name': 'Juan Garcia'},
    {'id': 2, 'name': 'Natalia Bertola'},
    {'id': 3, 'name': 'Santiago Perez'},
    {'id': 4, 'name': 'Vanesa Olmo'}
]


@app.route('/say-hello/<name>', methods=['GET'])
def say_hello(name):
    response = {
        'message': f'Hello {name}'
    }
    return jsonify(response)


@app.route('/alumnos', methods=['GET'])
def get_alumnos():
    response = alumnos
    return jsonify(response)


@app.route('/alumnos', methods=['POST'])
def add_alumnos():
    json = request.get_json(force=True)
    if json.get('name') is None:
        return jsonify({'message': 'Bad request'}), 400
    s = {'id': len(alumnos) + 1,
         'name': json.get('name')}
    alumnos.append(s)
    response = alumnos
    return jsonify(s), 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
