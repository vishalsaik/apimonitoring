import { useState } from 'react'
import { Link } from 'react-router-dom'
import './Auth.css'
import './Onboard.css'

export default function Onboard() {
  const [form, setForm] = useState({ username: '', email: '', password: '' })
  const handle = e => setForm(f => ({ ...f, [e.target.name]: e.target.value }))

  return (
    <div className="auth-page">
      <div className="auth-bg-orb auth-bg-orb--1" />
      <div className="auth-bg-orb auth-bg-orb--2" />
      <div className="auth-bg-orb auth-bg-orb--3" />

      <div className="auth-card">
        <div className="auth-card-glow" />

        <div className="auth-header">
          <div className="auth-logo">⬡</div>
          <div className="onboard-badge">First Time Setup</div>
          <h1 className="auth-title">Create Super Admin</h1>
          <p className="auth-subtitle">This account will have full control over the platform. Set it up carefully.</p>
        </div>

        <form className="auth-form">
          <div className="auth-field">
            <label className="auth-label">Username</label>
            <div className="auth-input-wrap">
              <span className="auth-input-icon">◉</span>
              <input className="auth-input" type="text" name="username" value={form.username} onChange={handle} placeholder="superadmin" />
            </div>
          </div>
          <div className="auth-field">
            <label className="auth-label">Email</label>
            <div className="auth-input-wrap">
              <span className="auth-input-icon">◎</span>
              <input className="auth-input" type="email" name="email" value={form.email} onChange={handle} placeholder="admin@company.com" />
            </div>
          </div>
          <div className="auth-field">
            <label className="auth-label">Password</label>
            <div className="auth-input-wrap">
              <span className="auth-input-icon">◈</span>
              <input className="auth-input" type="password" name="password" value={form.password} onChange={handle} placeholder="••••••••••" />
            </div>
          </div>

          <div className="onboard-warning">
            <span>⚠</span>
            <span>This page is only accessible once. After setup you will be redirected to login.</span>
          </div>

          <button type="submit" className="auth-btn">
            <span>Initialize Platform</span>
            <span className="auth-btn-arrow">→</span>
          </button>
        </form>

        <div className="auth-footer">
          <span>Already set up?</span>
          <Link to="/login" className="auth-link">Sign in →</Link>
        </div>
      </div>

      <div className="auth-side">
        <div className="auth-side-content">
          <div className="auth-side-tag">Getting Started</div>
          <h2 className="auth-side-title">One account<br />to rule them all</h2>
          <p className="auth-side-desc">The super admin account is the root of your monitoring system. You'll use it to register clients, create users, and manage API keys.</p>
          <div className="onboard-steps">
            {['Create super admin', 'Register clients', 'Assign users', 'Monitor APIs'].map((step, i) => (
              <div className="onboard-step" key={i}>
                <div className={`onboard-step-num ${i === 0 ? 'active' : ''}`}>{i + 1}</div>
                <span className={`onboard-step-label ${i === 0 ? 'active' : ''}`}>{step}</span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
