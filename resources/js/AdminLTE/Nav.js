import { useContext } from 'react';
import { Link } from 'react-router-dom';

//import AuthContext from '../store/auth-context';

const Nav = (props) => {
    //const authCtx = useContext(AuthContext);

    const logoutHandler = () => {
        //authCtx.logout();
        // optional: redirect the user
    };

    return (
        <nav className="main-header navbar navbar-expand navbar-white navbar-light">
            <ul className="navbar-nav">
                <li className="nav-item">
                    <div className="nav-link" data-widget="pushmenu" href="#" role="button"><i className="fas fa-bars" /></div>
                </li>
                <li className="nav-item d-none d-sm-inline-block">
                    <Link to="/staff" className="nav-link">Staff</Link>
                </li>
            </ul>
            <ul className="navbar-nav ml-auto">
                <li className="nav-item">
                    <a className="nav-link" data-widget="fullscreen" href="#" role="button">
                        <i className="fas fa-expand-arrows-alt" />
                    </a>
                </li>
                <li className="nav-item">
                    <a className="nav-link" href="#" role="button" title="LogOut" onClick={logoutHandler}>
                        <i className="fas fa-power-off"></i>
                    </a>
                </li>
            </ul>

        </nav>
    );
}

export default Nav;
