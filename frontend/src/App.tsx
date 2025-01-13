import { useState } from 'react';
import { RegisterForm } from './components/auth/RegisterForm';
import { LoginForm } from './components/auth/LoginForm';
import { Button } from '@/components/ui/button';

function App() {
  const [showLogin, setShowLogin] = useState(false);

  return (
    <div className="min-h-screen bg-gray-100 py-8 px-4">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold text-center mb-8">Blog Application</h1>
        <div className="flex justify-center space-x-4 mb-8">
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
        {showLogin ? <LoginForm /> : <RegisterForm />}
      </div>
    </div>
  )
}

export default App
