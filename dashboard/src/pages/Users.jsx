import { useState } from 'react'
import './Users.css'

const mockUsers = [
  { id: 1, username: 'superadmin', email: 'admin@company.com', role: 'super_admin', isActive: true, createdAt: '2026-03-20', clientId: null },
  { id: 2, username: 'john.zomato', email: 'john@zomato.com', role: 'client_admin', isActive: true, createdAt: '2026-03-18', clientId: 'Zomato' },
  { id: 3, username: 'priya.swiggy', email: 'priya@swiggy.com', role: 'client_viewer', isActive: true, createdAt: '2026-03-17', clientId: 'Swiggy' },
  { id: 4, username: 'dev.blinkit', email: 'dev@blinkit.com', role: 'client_editor', isActive: false, createdAt: '2026-03-15', clientId: 'Blinkit' },
  { id: 5, username: 'ops.zepto', email: 'ops@zepto.com', role: 'client_viewer', isActive: true, createdAt: '2026-03-12', clientId: 'Zepto' },
]

const roleColors = {
  super_admin: { bg: 'rgba(79,142,255,0.12)', color: '#4f8eff', border: 'rgba(79,142,255,0.25)' },
  client_admin: { bg: 'rgba(167,139,250,0.12)', color: '#a78bfa', border: 'rgba(167,139,250,0.25)' },
  client_editor: { bg: 'rgba(34,211,238,0.12)', color: '#22d3ee', border: 'rgba(34,211,238,0.25)' },
  client_viewer: { bg: 'rgba(52,211,153,0.12)', color: '#34d399', border: 'rgba(52,211,153,0.25)' },
}

export default function Users() {
  const [showModal, setShowModal] = useState(false)
  const [search, setSearch] = useState('')
  const [form, setForm] = useState({ username: '', email: '', password: '', role: 'client_viewer' })
  const handle = e => setForm(f => ({ ...f, [e.target.name]: e.target.value }))

  const filtered = mockUsers.filter(u =>
    u.username.includes(search) || u.email.includes(search)
  )

  return (
    <div className="users-page">
      <div className="users-header">
        <div>
          <h1 className="users-title">Users</h1>
          <p className="users-desc">{mockUsers.length} users · {mockUsers.filter(u => u.isActive).length} active</p>
        </div>
        <button className="users-add-btn" onClick={() => setShowModal(true)}>
          <span>+</span> Register User
        </button>
      </div>

      {/* Filters */}
      <div className="users-filters">
        <div className="users-search">
          <span className="users-search-icon">⌕</span>
          <input
            type="text"
            placeholder="Search users..."
            className="users-search-input"
            value={search}
            onChange={e => setSearch(e.target.value)}
          />
        </div>
        <div className="users-filter-chips">
          {['All', 'Super Admin', 'Client Admin', 'Active', 'Inactive'].map((f, i) => (
            <button key={f} className={`filter-chip ${i === 0 ? 'active' : ''}`}>{f}</button>
          ))}
        </div>
      </div>

      {/* Table */}
      <div className="users-table-wrap">
        <table className="users-table">
          <thead>
            <tr>
              <th>User</th>
              <th>Role</th>
              <th>Client</th>
              <th>Status</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {filtered.map(user => {
              const rc = roleColors[user.role] || roleColors.client_viewer
              return (
                <tr key={user.id} className="users-table-row">
                  <td>
                    <div className="user-cell">
                      <div className="user-avatar" style={{ background: `linear-gradient(135deg, ${rc.color}40, ${rc.color}20)`, border: `1px solid ${rc.border}` }}>
                        <span style={{ color: rc.color }}>{user.username[0].toUpperCase()}</span>
                      </div>
                      <div>
                        <div className="user-name">{user.username}</div>
                        <div className="user-email">{user.email}</div>
                      </div>
                    </div>
                  </td>
                  <td>
                    <span className="role-badge" style={{ background: rc.bg, color: rc.color, borderColor: rc.border }}>
                      {user.role.replace(/_/g, ' ')}
                    </span>
                  </td>
                  <td>
                    <span className="user-client">{user.clientId ?? '—'}</span>
                  </td>
                  <td>
                    <div className="status-cell">
                      <span className={`status-dot ${user.isActive ? 'active' : 'inactive'}`} />
                      <span className={`status-label ${user.isActive ? 'active' : 'inactive'}`}>
                        {user.isActive ? 'Active' : 'Inactive'}
                      </span>
                    </div>
                  </td>
                  <td><span className="user-date">{user.createdAt}</span></td>
                  <td>
                    <div className="user-actions">
                      <button className="user-action-btn">✎</button>
                      <button className="user-action-btn danger">✕</button>
                    </div>
                  </td>
                </tr>
              )
            })}
          </tbody>
        </table>
      </div>

      {/* Register Modal */}
      {showModal && (
        <div className="modal-overlay" onClick={() => setShowModal(false)}>
          <div className="modal" onClick={e => e.stopPropagation()}>
            <div className="modal-header">
              <div>
                <div className="modal-title">Register New User</div>
                <div className="modal-sub">User will be able to login after registration</div>
              </div>
              <button className="modal-close" onClick={() => setShowModal(false)}>✕</button>
            </div>
            <div className="modal-body">
              {['username', 'email', 'password'].map(field => (
                <div className="modal-field" key={field}>
                  <label className="modal-label">{field.charAt(0).toUpperCase() + field.slice(1)}</label>
                  <input
                    className="modal-input"
                    type={field === 'password' ? 'password' : 'text'}
                    name={field}
                    value={form[field]}
                    onChange={handle}
                    placeholder={field === 'email' ? 'user@company.com' : ''}
                  />
                </div>
              ))}
              <div className="modal-field">
                <label className="modal-label">Role</label>
                <select className="modal-input" name="role" value={form.role} onChange={handle}>
                  <option value="client_admin">Client Admin</option>
                  <option value="client_editor">Client Editor</option>
                  <option value="client_viewer">Client Viewer</option>
                </select>
              </div>
            </div>
            <div className="modal-footer">
              <button className="modal-cancel" onClick={() => setShowModal(false)}>Cancel</button>
              <button className="modal-submit">Register User →</button>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
