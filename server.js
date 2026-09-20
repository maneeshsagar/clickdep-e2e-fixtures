console.log('STARTUP-MARKER slow booting (45s)...');
setTimeout(()=>require('http').createServer((q,r)=>r.end('slow-ok')).listen(8080,'0.0.0.0',()=>console.log('listening after 45s')),45000);
