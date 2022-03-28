import './App.css';
import CoreUI from './CoreUI/CoreUI';
import { Routes, Route } from "react-router-dom";
import Staff from './Staff/Staff';
import StaffAdd from './Staff/StaffAdd';
import Dashboard from './Dashboard/Dashboard';
import { ToastContainer } from 'react-toastify';
import XV from './xv/XV';
import Market from './Market/Market';

function App() {
  return (
    <CoreUI>
      <ToastContainer />
      <Routes>
        <Route path="/u" exact element={<Dashboard />} />
        <Route path="/u/staff" exact element={<Staff />} />
        <Route path="/u/staff/add" exact element={<StaffAdd />} />
        <Route path="/u/staff/edit/:id" exact element={<StaffAdd />} />
        <Route path="/u/xv" exact element={<XV />} />
        <Route path="/u/market" exact element={<Market />} />
      </Routes>
    </CoreUI>
  );
}

export default App;
