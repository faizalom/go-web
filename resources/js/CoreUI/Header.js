const Header = (props) => {

    const Sidebar = () => {
        coreui.Sidebar.getInstance(document.querySelector('#sidebar')).toggle();
    }

    return (
        <header className="header header-sticky mb-4">
            <div className="container-fluid">
                <button className="header-toggler px-md-0 me-md-3" type="button" onClick={Sidebar}>
                    <i className="fas fa-bars" aria-hidden="true"></i>
                </button>
                <a className="header-brand d-md-none" href="#">
                    <svg width={118} height={46} alt="CoreUI Logo">
                        <use xlinkHref="static/assets/brand/coreui.svg#full" />
                    </svg>
                </a>
                <ul className="header-nav d-none d-md-flex">
                    <li className="nav-item"><a className="nav-link" href="#">Dashboard</a></li>
                    <li className="nav-item"><a className="nav-link" href="#">Users</a></li>
                    <li className="nav-item"><a className="nav-link" href="#">Settings</a></li>
                </ul>
                <ul className="header-nav ms-auto">
                    <li className="nav-item"><i className="fas fa-envelope-open"></i></li>
                </ul>
                <ul className="header-nav ms-3">
                    <li className="nav-item dropdown">
                        <a className="nav-link py-0" data-coreui-toggle="dropdown" href="#" role="button" aria-haspopup="true" aria-expanded="false">
                            <div className="avatar avatar-md">
                                <img className="avatar-img" src={DATA.user.profile_photo} alt="user@email.com" />
                            </div>
                        </a>
                        <div className="dropdown-menu dropdown-menu-end pt-0">
                            <div className="dropdown-header bg-light py-2">
                                <div className="fw-semibold">Settings</div>
                            </div>
                            <a className="dropdown-item" href="#">
                                <i className="fas fa-user me-2"></i> Profile
                            </a>
                            <a className="dropdown-item" href="#">
                                <i className="fas fa-cog me-2"></i> Settings
                            </a>
                            <div className="dropdown-divider" />
                            <a className="dropdown-item" href="/logout">
                                <i className="fas fa-sign-out-alt me-2"></i> Logout
                            </a>
                        </div>
                    </li>
                </ul>
            </div>
            <div className="header-divider" />
            <div className="container-fluid">
                <nav aria-label="breadcrumb">
                    <ol className="breadcrumb my-0 ms-2">
                        <li className="breadcrumb-item">
                            {/* if breadcrumb is single*/}<span>Home</span>
                        </li>
                        <li className="breadcrumb-item active"><span>Blank</span></li>
                    </ol>
                </nav>
            </div>
        </header>
    )
}

export default Header
