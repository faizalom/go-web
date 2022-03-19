import LSide from './LSide';
import Nav from './Nav';
import RSide from './RSide';
import Footer from './Footer';
import Content from './Content';
import PageTitle from './PageTitle';
import { Fragment, useContext } from 'react';
// import { ToastContainer } from 'react-toastify';
// import AuthContext from '../store/auth-context';
import { useHistory } from 'react-router';


const AdminLTE = (props) => {
    // const authCtx = useContext(AuthContext);
    //const history = useHistory();

    // if (!authCtx.isLoggedIn) {
    //     history.push("/login")
    // }

    return (
        <Fragment>
            <Nav />
            <LSide />
            <div className="content-wrapper pt-3">
                {/* <ToastContainer /> */}
                {/* <PageTitle /> */}
                <Content>{props.children}</Content>
            </div>
            <RSide />
            <Footer />
        </Fragment>
    );
}

export default AdminLTE
