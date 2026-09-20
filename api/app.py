import os
from flask import Flask, jsonify, request
app=Flask(__name__)
print('STARTUP-MARKER multi-python-api',flush=True)
@app.route('/api/<path:p>')
def a(p):
    if p=='log': print('E2E-LOG-MARKER api',flush=True)
    return jsonify(svc='python-api',path=p,env={k:v for k,v in os.environ.items() if k.startswith('E2E_')},email=request.headers.get('X-Gate-Email'))
