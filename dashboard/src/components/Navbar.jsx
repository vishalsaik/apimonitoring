import './Navbar.css'

export default function Navbar({ user }) {
  return (
    <header className="navbar">
      <div className="navbar-left">
        <div className="navbar-breadcrumb">
          <span className="navbar-breadcrumb-root">Dashboard</span>
          <span className="navbar-breadcrumb-sep">›</span>
          <span className="navbar-breadcrumb-current">Overview</span>
        </div>
      </div>
      <div className="navbar-right">
        <div className="navbar-search">
          <span className="navbar-search-icon">⌕</span>
          <input type="text" placeholder="Search..." className="navbar-search-input" />
          <span className="navbar-search-kbd">⌘K</span>
        </div>
        <button className="navbar-notif">
          <span>◎</span>
          <span className="navbar-notif-badge">3</span>
        </button>
        <div className="navbar-user">
          <div className="navbar-avatar">
            {user?.username?.[0]?.toUpperCase() ?? 'S'}
          </div>
          <div className="navbar-user-info">
            <div className="navbar-user-name">{user?.username ?? 'superadmin'}</div>
            <div className="navbar-user-role">{user?.role ?? 'super_admin'}</div>
          </div>
          <button className="navbar-logout">
            <span>⏻</span>
          </button>
        </div>
      </div>
    </header>
  )
}
