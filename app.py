import os,sys
from flask import Flask, jsonify, request
app=Flask(__name__)
print('STARTUP-MARKER python',flush=True)
@app.route('/',defaults={'p':''})
@app.route('/<path:p>')
def a(p):
    if p=='log': print('E2E-LOG-MARKER',flush=True)
    return jsonify(lang='python',path='/'+p,env={k:v for k,v in os.environ.items() if k.startswith('E2E_')},email=request.headers.get('X-Gate-Email'))
