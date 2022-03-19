const Header = (props) => {

    const Sidebar = () => {
        coreui.Sidebar.getInstance(document.querySelector('#sidebar')).toggle();
    }

    return (
        <header className="header header-sticky mb-4">
            <div className="container-fluid">
                <button className="header-toggler px-md-0 me-md-3" type="button" onClick={Sidebar}>
                    <svg className="icon icon-lg">
                        <use xlinkHref="static/vendors/@coreui/icons/svg/free.svg#cil-menu" />
                    </svg>
                </button><a className="header-brand d-md-none" href="#">
                    <svg width={118} height={46} alt="CoreUI Logo">
                        <use xlinkHref="static/assets/brand/coreui.svg#full" />
                    </svg></a>
                <ul className="header-nav d-none d-md-flex">
                    <li className="nav-item"><a className="nav-link" href="#">Dashboard</a></li>
                    <li className="nav-item"><a className="nav-link" href="#">Users</a></li>
                    <li className="nav-item"><a className="nav-link" href="#">Settings</a></li>
                </ul>
                <ul className="header-nav ms-auto">
                    <li className="nav-item">
                        <a className="nav-link" href="#">
                            <svg className="icon icon-lg"><use xlinkHref="static/vendors/@coreui/icons/svg/free.svg#cil-envelope-open" /></svg>
                        </a>
                    </li>
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
                                <svg className="icon me-2">
                                    <use xlinkHref="static/vendors/@coreui/icons/svg/free.svg#cil-user" />
                                </svg> Profile
                            </a>
                            <a className="dropdown-item" href="#">
                                <svg className="icon me-2">
                                    <use xlinkHref="static/vendors/@coreui/icons/svg/free.svg#cil-settings" />
                                </svg> Settings
                            </a>
                            <div className="dropdown-divider" />
                            <a className="dropdown-item" href="/logout">
                                <svg className="icon me-2">
                                    <use xlinkHref="static/vendors/@coreui/icons/svg/free.svg#cil-account-logout" />
                                </svg> Logout
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
