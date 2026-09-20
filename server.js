console.log('STARTUP-MARKER very slow boot (120s)');setTimeout(()=>require('http').createServer((q,r)=>r.end('late')).listen(8080,'0.0.0.0'),120000)
