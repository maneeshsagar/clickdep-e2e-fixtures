const express = require('express')
const app = express()
const crypto = require('crypto')
// Verbatim from docs/gate-proof.md — the point of this branch is to prove that snippet works.
const pubB64 = process.env.CLICKDEP_GATE_PUBKEY
const pub = pubB64 ? crypto.createPublicKey({
  key: Buffer.concat([Buffer.from('302a300506032b6570032100', 'hex'), Buffer.from(pubB64, 'base64')]),
  format: 'der', type: 'spki',
}) : null
function requireGate(req, res, next) {
  if (!pub) return res.status(500).json({ error: 'CLICKDEP_GATE_PUBKEY not set' })
  const proof = req.get('X-Gate-Proof') || ''
  const [ts, sig] = proof.split('.')
  const now = Math.floor(Date.now() / 1000)
  if (!ts || !sig || Math.abs(now - Number(ts)) > 60) return res.sendStatus(401)
  const path = req.originalUrl.split('?')[0]
  const msg = ['clickdep-gate-proof-v1', req.get('X-Gate-App'), req.method, path, req.get('X-Gate-Subject') || '', ts].join('\n')
  const ok = crypto.verify(null, Buffer.from(msg), pub, Buffer.from(sig, 'base64'))
  return ok ? next() : res.sendStatus(401)
}
app.use(requireGate)
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
