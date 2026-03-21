import { useState } from 'react'
import { Link } from 'react-router-dom'
import './Auth.css'

export default function Login() {
  const [form, setForm] = useState({ username: '', password: '' })
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
          <h1 className="auth-title">Welcome back</h1>
          <p className="auth-subtitle">Sign in to your monitoring dashboard</p>
        </div>

        <form className="auth-form">
          <div className="auth-field">
            <label className="auth-label">Username</label>
            <div className="auth-input-wrap">
              <span className="auth-input-icon">◉</span>
              <input
                className="auth-input"
                type="text"
                name="username"
                value={form.username}
                onChange={handle}
                placeholder="Enter your username"
                autoComplete="username"
              />
            </div>
          </div>

          <div className="auth-field">
            <div className="auth-label-row">
              <label className="auth-label">Password</label>
              <a href="#" className="auth-forgot">Forgot password?</a>
            </div>
            <div className="auth-input-wrap">
              <span className="auth-input-icon">◈</span>
              <input
                className="auth-input"
                type="password"
                name="password"
                value={form.password}
                onChange={handle}
                placeholder="••••••••••"
                autoComplete="current-password"
              />
            </div>
          </div>

          <button type="submit" className="auth-btn">
            <span>Sign In</span>
            <span className="auth-btn-arrow">→</span>
          </button>
        </form>

        <div className="auth-footer">
          <span>First time here?</span>
          <Link to="/onboard" className="auth-link">Set up your account →</Link>
        </div>
      </div>

      <div className="auth-side">
        <div className="auth-side-content">
          <div className="auth-side-tag">API Monitoring Platform</div>
          <h2 className="auth-side-title">Real-time insights<br />for your APIs</h2>
          <p className="auth-side-desc">Monitor performance, track errors, and analyse traffic patterns across all your services.</p>
          <div className="auth-side-stats">
            <div className="auth-stat">
              <div className="auth-stat-value">99.9%</div>
              <div className="auth-stat-label">Uptime SLA</div>
            </div>
            <div className="auth-stat-divider" />
            <div className="auth-stat">
              <div className="auth-stat-value">&lt; 2ms</div>
              <div className="auth-stat-label">Avg Latency</div>
            </div>
            <div className="auth-stat-divider" />
            <div className="auth-stat">
              <div className="auth-stat-value">1M+</div>
              <div className="auth-stat-label">Events/day</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
