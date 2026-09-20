const http=require('http');const {Pool}=require('pg');
console.log('STARTUP-MARKER single-pg');
const pool=new Pool({connectionString:process.env.DATABASE_URL});
http.createServer(async(q,r)=>{
  r.setHeader('content-type','application/json');
  try{
    await pool.query('create table if not exists hits(id serial primary key, at timestamptz default now())');
    if(q.url==='/log') console.log('E2E-LOG-MARKER');
    await pool.query('insert into hits default values');
    const c=await pool.query('select count(*)::int as n from hits');
    r.end(JSON.stringify({db:'ok',hits:c.rows[0].n,has_database_url:!!process.env.DATABASE_URL}));
  }catch(e){ console.error('DB-ERROR',e.message); r.statusCode=500; r.end(JSON.stringify({db:'error',msg:e.message})) }
}).listen(8080,'0.0.0.0');
