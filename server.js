for(let i=1;i<=800;i++) console.log('BULK-LOG-LINE '+i);
console.log('STARTUP-MARKER many-logs done');
require('http').createServer((q,r)=>{console.log('E2E-LOG-MARKER');r.end('ok')}).listen(8080,'0.0.0.0');
