import { RegisterForm } from './components/auth/RegisterForm';

function App() {
  return (
    <div className="min-h-screen bg-gray-100 py-8 px-4">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold text-center mb-8">Blog Application</h1>
        <RegisterForm />
      </div>
    </div>
  )
}

export default App
