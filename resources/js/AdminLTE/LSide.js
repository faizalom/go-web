import { Link, NavLink } from "react-router-dom"

const LSide = (props) => {
    return (
        <aside className="main-sidebar sidebar-dark-primary elevation-4">
            <Link to="/" className="brand-link">
                <img src="/dist/img/AdminLTELogo.png" alt="AdminLTE Logo" className="brand-image img-circle elevation-3" style={{ opacity: '.8' }} />
                <span className="brand-text font-weight-light">OCS</span>
            </Link>
            <div className="sidebar">
                <nav className="mt-2">
                    <ul className="nav nav-pills nav-sidebar flex-column nav-flat nav-legacy nav-child-indent" data-widget="treeview" role="menu" data-accordion="false">
                        <li className="nav-item">
                            <NavLink to="/staff" className="nav-link">
                                <i className="nav-icon fas fa-users"></i>
                                <p>Staff</p>
                            </NavLink>
                        </li>
                    </ul>
                </nav>
            </div>
        </aside>
    )
}

export default LSide
