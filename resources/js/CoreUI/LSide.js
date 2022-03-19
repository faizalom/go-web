import { NavLink } from "react-router-dom"

const LSide = (props) => {
    return (
        <div className="sidebar sidebar-dark sidebar-fixed" id="sidebar">
            <div className="sidebar-brand d-none d-md-flex">
                <svg className="sidebar-brand-full" width={118} height={46} alt="CoreUI Logo">
                    <use xlinkHref="static/assets/brand/coreui.svg#full" />
                </svg>
                <svg className="sidebar-brand-narrow" width={46} height={46} alt="CoreUI Logo">
                    <use xlinkHref="static/assets/brand/coreui.svg#signet" />
                </svg>
            </div>
            <ul className="sidebar-nav" data-coreui="navigation" data-simplebar>
                <li className="nav-title">{DATA.user.firstName} {DATA.user.lastName}</li>
                {/* <li className="nav-item px-3 d-narrow-none">
                    <div className="text-uppercase mb-1"><small><b>CPU Usage</b></small></div>
                    <div className="progress progress-thin">
                        <div className="progress-bar bg-info-gradient" role="progressbar" style={{ width: '25%' }} aria-valuenow={25} aria-valuemin={0} aria-valuemax={100} />
                    </div><small className="text-medium-emphasis-inverse">348 Processes. 1/4 Cores.</small>
                </li> */}

                {MENU.map(menu => {

                    // <li className="nav-item"><a className="nav-link" href="colors.html">
                    // <svg className="nav-icon">
                    // <use xlink:href="vendors/@coreui/icons/svg/free.svg#cil-drop"></use>
                    // </svg> Colors</a></li>

                    return (<li className={menu.Children ? "nav-group" : "nav-item"} key={menu.Link + menu.Text}>
                        <NavLink to={menu.Link} className={menu.Children ? "nav-link nav-group-toggle" : "nav-link"} >
                            <i className={menu.Icon + " nav-icon" }></i> {menu.Text}
                        </NavLink>
                        {menu.Children && (
                            <ul class="nav-group-items">
                                {menu.Children.map(m =>
                                    <li class="nav-item"><NavLink to={m.Link} className="nav-link">{m.Text}</NavLink></li>
                                )}
                            </ul>
                        )}
                    </li>)
                })}

            </ul>
            <button className="sidebar-toggler" type="button" data-coreui-toggle="unfoldable" />
        </div>

    )
}

export default LSide
