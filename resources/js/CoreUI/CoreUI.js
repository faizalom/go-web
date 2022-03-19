import { Fragment } from "react";
import Footer from "./Footer";
import Header from "./Header";
import LSide from "./LSide";

const CoreUI = (props) => {

    return (
        <Fragment>
            <LSide />
            <div className="wrapper d-flex flex-column min-vh-100 bg-light">
                <Header />
                <div className="body flex-grow-1 px-3">
                    {props.children}
                </div>
                <Footer />
            </div>
        </Fragment>
    );
}

export default CoreUI
