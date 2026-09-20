const express = require('express')
const app = express()
const e2e = () => Object.fromEntries(Object.entries(process.env).filter(([k]) => k.startsWith('E2E_')))
app.get('/api/health', (req, res) => res.json({ ok: true, lang: 'node' }))
app.get('/api/env-echo', (req, res) => res.json({ lang: 'node', env: e2e(), has_database_url: !!process.env.DATABASE_URL }))
app.get('/api/whoami', (req, res) => {
  const cookies = (req.headers.cookie || '').split(';').map(s => s.trim().split('=')[0]).filter(Boolean)
  res.json({ lang: 'node', gate_email: req.headers['x-gate-email'] || null, gate_app: req.headers['x-gate-app'] || null,
    cookie_names: cookies, has_gate_session_cookie: cookies.includes('_gate_session') })
})
app.get('/apiary', (req, res) => res.status(200).send('BACKEND-APIARY-LEAK'))   // must never be reachable through the gate
app.listen(3000, '0.0.0.0', () => console.log('node api on 3000'))
