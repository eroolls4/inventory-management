import { Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Items from './components/Items';
import Signup from "./components/Signup.tsx";


const App: React.FC = () => {
    return (
        <div className="min-h-screen ">
            <Routes>
                <Route path="/login" element={<Login />} />
                  <Route path="/signup" element={<Signup />} />

                <Route
                    path="/items"
                    element={<Items />}
                />
                <Route path="/" element={<Navigate to={ '/login'} />} />
            </Routes>
        </div>
    );
};

export default App;