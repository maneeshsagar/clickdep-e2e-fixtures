import React, { useEffect, useState } from 'react'
import { createRoot } from 'react-dom/client'
function App() {
  const [api, setApi] = useState('loading')
  const [env, setEnv] = useState(null)
  useEffect(() => {
    fetch('/api/health').then(r => r.text()).then(t => setApi(t.slice(0, 120))).catch(e => setApi('ERR ' + e))
    fetch('/env.json').then(r => r.json()).then(setEnv).catch(() => setEnv({}))
  }, [])
  return (<div>
    <h1 id="ui">E2E-REACT-UI</h1>
    <p>path: {location.pathname}</p>
    <pre id="api">{api}</pre>
    <pre id="env">{JSON.stringify(env)}</pre>
  </div>)
}
createRoot(document.getElementById('root')).render(<App />)
