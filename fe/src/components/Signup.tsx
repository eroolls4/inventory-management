import { useState } from 'react';
import axios from 'axios';
import { useNavigate, Link } from 'react-router-dom';

const Signup: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:3000/signup', { email, password });
            console.log(response.data);
            navigate('/login');
        } catch (err) {
            setError('Signup failed. Email might be taken or server error.');
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen">
            <div className="bg-white p-6 rounded shadow-md w-96">
                <h2 className="text-2xl font-bold mb-4 text-center">Sign Up</h2>
                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <input
                            type="email"
                            className="w-full border p-2 rounded"
                            placeholder="Email"
                            value={email}
                            onChange={e => setEmail(e.target.value)}
                            required
                        />
                    </div>
                    <div className="mb-4">
                        <input
                            type="password"
                            className="w-full border p-2 rounded"
                            placeholder="Password"
                            value={password}
                            onChange={e => setPassword(e.target.value)}
                            required
                        />
                    </div>
                    {error && <p className="text-red-500 mb-4">{error}</p>}
                    <button type="submit" className="w-full bg-blue-500 text-white p-2 rounded">
                        Sign Up
                    </button>
                </form>
                <p className="mt-4 text-center">
                    Already have an account? <Link to="/login" className="text-blue-500">Login</Link>
                </p>
            </div>
        </div>
    );
};

export default Signup;