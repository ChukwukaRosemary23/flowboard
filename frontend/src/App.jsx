import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import Homepage from './pages/Homepage';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import BoardView from './pages/BoardView';
import FeaturesPage from './pages/FeaturesPage';
import IntegrationsPage from './pages/IntegrationsPage';
import PricingPage from './pages/PricingPage';
import InboxPage from './pages/InboxPage';
import PlannerPage from './pages/PlannerPage';

// Resources Pages
import Blog from './pages/resources/Blog';
import Guides from './pages/resources/Guides';
import Webinars from './pages/resources/Webinars';
import HelpCenter from './pages/resources/HelpCenter';

// Solutions Pages
import MarketingTeams from './pages/solutions/MarketingTeams';
import ProductManagement from './pages/solutions/ProductManagement';
import EngineeringTeams from './pages/solutions/EngineeringTeams';

// Protected Route Component
function ProtectedRoute({ children }) {
  const { user, loading } = useAuth();
  if (loading) {
    return (
      <div className="min-h-screen bg-blue-600 flex items-center justify-center">
        <div className="text-white text-xl">Loading...</div>
      </div>
    );
  }
  return user ? children : <Navigate to="/login" />;
}

// Public Route (redirect if logged in)
function PublicRoute({ children }) {
  const { user, loading } = useAuth();
  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
        <div className="text-white text-xl">Loading...</div>
      </div>
    );
  }
  return user ? <Navigate to="/dashboard" /> : children;
}

function AppRoutes() {
  return (
    <Routes>
      {/* Public Pages */}
      <Route path="/" element={<Homepage />} />
      <Route path="/features" element={<FeaturesPage />} />
      <Route path="/features/inbox" element={<InboxPage />} />
      <Route path="/features/planner" element={<PlannerPage />} />
      <Route path="/integrations" element={<IntegrationsPage />} />
      <Route path="/pricing" element={<PricingPage />} />

      {/* Resources Pages */}
      <Route path="/blog" element={<Blog />} />
      <Route path="/guides" element={<Guides />} />
      <Route path="/webinars" element={<Webinars />} />
      <Route path="/help" element={<HelpCenter />} />

      {/* Solutions Pages */}
      <Route path="/solutions/marketing" element={<MarketingTeams />} />
      <Route path="/solutions/product" element={<ProductManagement />} />
      <Route path="/solutions/engineering" element={<EngineeringTeams />} />

      {/* Auth Routes */}
      <Route
        path="/login"
        element={
          <PublicRoute>
            <Login />
          </PublicRoute>
        }
      />
      <Route
        path="/register"
        element={
          <PublicRoute>
            <Register />
          </PublicRoute>
        }
      />

      {/* Protected Routes */}
      <Route
        path="/dashboard"
        element={
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        }
      />
      <Route
        path="/board/:id"
        element={
          <ProtectedRoute>
            <BoardView />
          </ProtectedRoute>
        }
      />

      {/* Catch all */}
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
}

function App() {
  return (
    <AuthProvider>
      <Router>
        <AppRoutes />
      </Router>
    </AuthProvider>
  );
}

export default App;