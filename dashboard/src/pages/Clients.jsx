import './Clients.css'

const mockClients = [
  { id: 1, name: 'Zomato', domain: 'zomato.com', plan: 'Enterprise', users: 4, apiKeys: 3, requests: '842K', status: 'active', joined: '2026-01-15' },
  { id: 2, name: 'Swiggy', domain: 'swiggy.com', plan: 'Pro', users: 2, apiKeys: 2, requests: '623K', status: 'active', joined: '2026-01-22' },
  { id: 3, name: 'Blinkit', domain: 'blinkit.com', plan: 'Starter', users: 1, apiKeys: 1, requests: '298K', status: 'inactive', joined: '2026-02-08' },
  { id: 4, name: 'Zepto', domain: 'zepto.com', plan: 'Pro', users: 3, apiKeys: 2, requests: '183K', status: 'active', joined: '2026-02-14' },
  { id: 5, name: 'Dunzo', domain: 'dunzo.com', plan: 'Starter', users: 1, apiKeys: 1, requests: '91K', status: 'active', joined: '2026-03-01' },
]

const planColors = {
  Enterprise: { bg: 'rgba(79,142,255,0.12)', color: '#4f8eff', border: 'rgba(79,142,255,0.25)' },
  Pro: { bg: 'rgba(167,139,250,0.12)', color: '#a78bfa', border: 'rgba(167,139,250,0.25)' },
  Starter: { bg: 'rgba(52,211,153,0.12)', color: '#34d399', border: 'rgba(52,211,153,0.25)' },
}

const initials = name => name.slice(0, 2).toUpperCase()

export default function Clients() {
  return (
    <div className="clients-page">
      <div className="clients-header">
        <div>
          <h1 className="clients-title">Clients</h1>
          <p className="clients-desc">{mockClients.length} registered clients · {mockClients.filter(c => c.status === 'active').length} active</p>
        </div>
        <button className="clients-add-btn">
          <span>+</span> Add Client
        </button>
      </div>

      <div className="clients-grid">
        {mockClients.map(client => {
          const pc = planColors[client.plan]
          return (
            <div className="client-card" key={client.id}>
              <div className="client-card-top">
                <div className="client-avatar">{initials(client.name)}</div>
                <div className={`client-status ${client.status}`}>
                  <span className="client-status-dot" />
                  {client.status}
                </div>
              </div>

              <div className="client-name">{client.name}</div>
              <div className="client-domain">{client.domain}</div>

              <span className="client-plan" style={{ background: pc.bg, color: pc.color, borderColor: pc.border }}>
                {client.plan}
              </span>

              <div className="client-stats">
                <div className="client-stat">
                  <div className="client-stat-value">{client.requests}</div>
                  <div className="client-stat-label">Requests</div>
                </div>
                <div className="client-stat-divider" />
                <div className="client-stat">
                  <div className="client-stat-value">{client.users}</div>
                  <div className="client-stat-label">Users</div>
                </div>
                <div className="client-stat-divider" />
                <div className="client-stat">
                  <div className="client-stat-value">{client.apiKeys}</div>
                  <div className="client-stat-label">API Keys</div>
                </div>
              </div>

              <div className="client-footer">
                <span className="client-joined">Joined {client.joined}</span>
                <div className="client-actions">
                  <button className="client-btn">Manage</button>
                  <button className="client-btn client-btn--icon">⋯</button>
                </div>
              </div>
            </div>
          )
        })}
      </div>
    </div>
  )
}
