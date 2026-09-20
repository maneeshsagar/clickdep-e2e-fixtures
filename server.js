const http=require('http'),fs=require('fs');
const env=()=>Object.fromEntries(Object.entries(process.env).filter(([k])=>k.startsWith('E2E_')));
console.log('STARTUP-MARKER node');
const PORT=+process.env.PORT_OVERRIDE||8080;
http.createServer((q,r)=>{
  if(q.url==='/log'){console.log('E2E-LOG-MARKER '+Date.now());}
  if(q.url==='/count'){let n=0;try{n=+fs.readFileSync('/data/count','utf8')}catch(e){}n++;try{fs.writeFileSync('/data/count',String(n))}catch(e){}r.setHeader('content-type','application/json');return r.end(JSON.stringify({count:n}))}
  r.setHeader('content-type','application/json');
  r.end(JSON.stringify({lang:'node',path:q.url,env:env(),email:q.headers['x-gate-email']||null,user:process.getuid()}));
}).listen(PORT,'0.0.0.0',()=>console.log('listening on '+PORT));
