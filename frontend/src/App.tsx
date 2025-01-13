import { useState, useEffect } from 'react';
import { RegisterForm } from './components/auth/RegisterForm';
import { LoginForm } from './components/auth/LoginForm';
import { ArticleForm } from './components/articles/ArticleForm';
import { ArticleList } from './components/articles/ArticleList';
import { Button } from '@/components/ui/button';

function App() {
  const [showLogin, setShowLogin] = useState(false);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [showArticleForm, setShowArticleForm] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    setIsAuthenticated(!!token);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
    setShowArticleForm(false);
  };

  return (
    <div className="min-h-screen bg-gray-100 py-8 px-4">
      <div className="max-w-7xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold">Blog Application</h1>
          {isAuthenticated ? (
            <div className="space-x-4">
              <Button
                variant={showArticleForm ? "default" : "outline"}
                onClick={() => setShowArticleForm(true)}
              >
                New Article
              </Button>
              <Button
                variant="outline"
                onClick={handleLogout}
              >
                Logout
              </Button>
            </div>
          ) : (
            <div className="space-x-4">
              <Button
                variant={!showLogin ? "default" : "outline"}
                onClick={() => setShowLogin(false)}
              >
                Register
              </Button>
              <Button
                variant={showLogin ? "default" : "outline"}
                onClick={() => setShowLogin(true)}
              >
                Login
              </Button>
            </div>
          )}
        </div>
        
        {!isAuthenticated ? (
          showLogin ? <LoginForm /> : <RegisterForm />
        ) : (
          showArticleForm ? <ArticleForm /> : <ArticleList />
        )}
      </div>
    </div>
  )
}

export default App
