import './Dashboard.css'

const stats = [
  { label: 'Total Requests', value: '2.4M', change: '+12.5%', up: true, icon: '◈', color: '#4f8eff' },
  { label: 'Active Clients', value: '18', change: '+2 this week', up: true, icon: '◉', color: '#a78bfa' },
  { label: 'Avg Latency', value: '1.8ms', change: '-0.3ms', up: true, icon: '◇', color: '#22d3ee' },
  { label: 'Error Rate', value: '0.04%', change: '-0.01%', up: true, icon: '◎', color: '#34d399' },
]

const recentActivity = [
  { client: 'Zomato', endpoint: 'POST /api/orders', status: 200, latency: '1.2ms', time: '2s ago' },
  { client: 'Swiggy', endpoint: 'GET /api/menu', status: 200, latency: '0.8ms', time: '5s ago' },
  { client: 'Blinkit', endpoint: 'PUT /api/inventory', status: 500, latency: '45ms', time: '12s ago' },
  { client: 'Zepto', endpoint: 'GET /api/products', status: 200, latency: '1.1ms', time: '18s ago' },
  { client: 'Dunzo', endpoint: 'DELETE /api/cart', status: 404, latency: '2.3ms', time: '31s ago' },
  { client: 'Zomato', endpoint: 'GET /api/user', status: 200, latency: '0.6ms', time: '45s ago' },
]

const topEndpoints = [
  { path: 'GET /api/orders', calls: 842340, pct: 92 },
  { path: 'POST /api/auth', calls: 623100, pct: 68 },
  { path: 'GET /api/menu', calls: 512880, pct: 56 },
  { path: 'PUT /api/cart', calls: 298440, pct: 32 },
  { path: 'GET /api/profile', calls: 183200, pct: 20 },
]

function statusColor(s) {
  if (s < 300) return '#34d399'
  if (s < 400) return '#fbbf24'
  if (s < 500) return '#f97316'
  return '#ff6b6b'
}

export default function Dashboard() {
  return (
    <div className="dashboard">
      <div className="dashboard-header">
        <div>
          <h1 className="dashboard-title">Overview</h1>
          <p className="dashboard-desc">Real-time API monitoring across all clients</p>
        </div>
        <div className="dashboard-header-actions">
          <div className="dashboard-time-range">
            {['1H', '24H', '7D', '30D'].map((t, i) => (
              <button key={t} className={`time-btn ${i === 1 ? 'active' : ''}`}>{t}</button>
            ))}
          </div>
          <button className="dashboard-refresh-btn">⟳ Refresh</button>
        </div>
      </div>

      {/* Stats */}
      <div className="dashboard-stats">
        {stats.map((s, i) => (
          <div className="stat-card" key={i} style={{ '--accent-color': s.color }}>
            <div className="stat-card-top">
              <div className="stat-icon" style={{ color: s.color }}>{s.icon}</div>
              <span className={`stat-change ${s.up ? 'up' : 'down'}`}>
                {s.up ? '↑' : '↓'} {s.change}
              </span>
            </div>
            <div className="stat-value">{s.value}</div>
            <div className="stat-label">{s.label}</div>
            <div className="stat-bar">
              <div className="stat-bar-fill" style={{ width: `${60 + i * 10}%`, background: s.color }} />
            </div>
          </div>
        ))}
      </div>

      <div className="dashboard-grid">
        {/* Recent Activity */}
        <div className="dashboard-panel">
          <div className="panel-header">
            <div>
              <div className="panel-title">Live Activity</div>
              <div className="panel-sub">Latest API requests</div>
            </div>
            <div className="live-badge">
              <span className="live-dot" />
              LIVE
            </div>
          </div>
          <div className="activity-list">
            {recentActivity.map((a, i) => (
              <div className="activity-row" key={i}>
                <div className="activity-client">{a.client[0]}</div>
                <div className="activity-info">
                  <div className="activity-endpoint">{a.endpoint}</div>
                  <div className="activity-meta">{a.client} · {a.time}</div>
                </div>
                <div className="activity-latency">{a.latency}</div>
                <div className="activity-status" style={{ color: statusColor(a.status), borderColor: statusColor(a.status) + '33' }}>
                  {a.status}
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Top Endpoints */}
        <div className="dashboard-panel">
          <div className="panel-header">
            <div>
              <div className="panel-title">Top Endpoints</div>
              <div className="panel-sub">By request volume</div>
            </div>
            <button className="panel-action">View all →</button>
          </div>
          <div className="endpoint-list">
            {topEndpoints.map((e, i) => (
              <div className="endpoint-row" key={i}>
                <div className="endpoint-rank">{i + 1}</div>
                <div className="endpoint-info">
                  <div className="endpoint-path">{e.path}</div>
                  <div className="endpoint-bar-wrap">
                    <div className="endpoint-bar">
                      <div className="endpoint-bar-fill" style={{ width: `${e.pct}%` }} />
                    </div>
                    <span className="endpoint-calls">{(e.calls / 1000).toFixed(0)}k</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
