import { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';
import { userApi } from '../services/api';
import toast from 'react-hot-toast';

function Home() {
  const { user, logout } = useAuth();
  const [userData, setUserData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await userApi.getHomeData();
        setUserData(response.data);
      } catch (error) {
        console.error('Erro ao carregar dados:', error);
        toast.error('Erro ao carregar dados');
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const handleLogout = () => {
    toast.success('Até logo! 👋', {
      duration: 2000,
    },
    logout()
);
  };

  if (loading) {
    return <div>Carregando...</div>;
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <nav className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <h1 className="text-xl font-bold">Quiz App</h1>
            </div>
            <div className="flex items-center">
              <span className="text-gray-700 mr-4">Olá, {user?.name}</span>
              <button
                onClick={handleLogout}
                className="bg-red-600 text-white px-4 py-2 rounded-md text-sm font-medium hover:bg-red-700"
              >
                Sair
              </button>
            </div>
          </div>
        </div>
      </nav>

      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <div className="bg-white shadow rounded-lg p-6">
            <h2 className="text-2xl font-bold mb-4">Bem-vindo à sua área</h2>
            {userData && (
              <div className="space-y-4">
                <p><strong>Nome:</strong> {userData.user.name}</p>
                <p><strong>Email:</strong> {userData.user.email}</p>
                <p><strong>Usuário:</strong> {userData.user.username}</p>
              </div>
            )}
          </div>
        </div>
      </main>
    </div>
  );
}

export default Home;