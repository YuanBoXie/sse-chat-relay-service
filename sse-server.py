from flask import Flask, Response, stream_with_context
import time
from flask_cors import CORS
app = Flask(__name__)
CORS(app)

def generate():
    while True:
        time.sleep(1)
        print(f"data: {time.ctime()}\n\n")
        yield f"data: {time.ctime()}\n\n"
    print("closed")

@app.route('/sse')
def sse():
    return Response(stream_with_context(generate()), mimetype='text/event-stream')

if __name__ == '__main__':
    app.run(debug=True)
