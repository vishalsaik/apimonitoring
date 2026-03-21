import { NavLink } from 'react-router-dom'
import './Sidebar.css'

const nav = [
  { to: '/',        icon: '◈', label: 'Overview' },
  { to: '/users',   icon: '◉', label: 'Users' },
  { to: '/clients', icon: '◎', label: 'Clients' },
  { to: '/apikeys', icon: '⬡', label: 'API Keys' },
  { to: '/analytics', icon: '◇', label: 'Analytics' },
]

export default function Sidebar() {
  return (
    <aside className="sidebar">
      <div className="sidebar-logo">
        <div className="sidebar-logo-icon">
          <span>⬡</span>
        </div>
        <div>
          <div className="sidebar-logo-title">API Monitor</div>
          <div className="sidebar-logo-sub">v2.0</div>
        </div>
      </div>

      <div className="sidebar-section-label">Navigation</div>
      <nav className="sidebar-nav">
        {nav.map(item => (
          <NavLink
            key={item.to}
            to={item.to}
            end={item.to === '/'}
            className={({ isActive }) => `sidebar-link ${isActive ? 'active' : ''}`}
          >
            <span className="sidebar-link-icon">{item.icon}</span>
            <span className="sidebar-link-label">{item.label}</span>
            <span className="sidebar-link-dot" />
          </NavLink>
        ))}
      </nav>

      <div className="sidebar-divider" />
      <div className="sidebar-section-label">System</div>
      <nav className="sidebar-nav">
        <NavLink to="/settings" className={({ isActive }) => `sidebar-link ${isActive ? 'active' : ''}`}>
          <span className="sidebar-link-icon">◫</span>
          <span className="sidebar-link-label">Settings</span>
          <span className="sidebar-link-dot" />
        </NavLink>
      </nav>

      <div className="sidebar-footer">
        <div className="sidebar-status">
          <span className="sidebar-status-dot" />
          <span>All systems operational</span>
        </div>
      </div>
    </aside>
  )
}
