if(!process.env.E2E_REQUIRED){console.error('FATAL: E2E_REQUIRED is not set');process.exit(1)}
require('http').createServer((q,r)=>r.end(JSON.stringify({required:process.env.E2E_REQUIRED}))).listen(8080,'0.0.0.0',()=>console.log('STARTUP-MARKER required-ok'));
