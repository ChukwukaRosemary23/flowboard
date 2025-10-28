import { useState } from 'react';
import { Link } from 'react-router-dom';
import { ChevronDown } from 'lucide-react';

function Navbar() {
  const [openDropdown, setOpenDropdown] = useState(null);

  const toggleDropdown = (dropdown) => {
    setOpenDropdown(openDropdown === dropdown ? null : dropdown);
  };

  return (
    <nav className="bg-white border-b border-gray-200 sticky top-0 z-50">
      {/* Banner */}
      <div className="bg-gradient-to-r from-blue-50 to-purple-50 border-b border-blue-100">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-2 text-center">
          <p className="text-sm text-gray-700">
            Accelerate your teams' work with FlowBoard Intelligence (AI) features 🤖 now available for all Premium and Enterprise!{' '}
            <a href="#" className="text-blue-600 hover:underline">Learn more</a>
          </p>
        </div>
      </div>

      {/* Main Nav */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo */}
          <Link to="/" className="flex items-center space-x-2">
            <img src="/logo.png" alt="FlowBoard" className="h-8" />
            <span className="text-2xl font-bold text-gray-900">FlowBoard</span>
          </Link>

          {/* Navigation Links with Dropdowns */}
          <div className="hidden md:flex items-center space-x-1">
            {/* Features Dropdown */}
            <div className="relative">
              <button
                onClick={() => toggleDropdown('features')}
                className="flex items-center space-x-1 px-3 py-2 text-gray-700 hover:text-blue-600 transition"
              >
                <span>Features</span>
                <ChevronDown className="w-4 h-4" />
              </button>
              {openDropdown === 'features' && (
                <div className="absolute top-full left-0 mt-2 w-64 bg-white rounded-lg shadow-xl border border-gray-200 py-2 z-50">
                  <Link to="/features/inbox" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">📥 Inbox</div>
                    <div className="text-sm text-gray-600">Capture every idea</div>
                  </Link>
                  <Link to="/features" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">📋 Boards</div>
                    <div className="text-sm text-gray-600">Organize your work</div>
                  </Link>
                  <Link to="/features/planner" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">📅 Planner</div>
                    <div className="text-sm text-gray-600">Sync your calendar</div>
                  </Link>
                  <Link to="/features" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">⚡ Automation</div>
                    <div className="text-sm text-gray-600">Let robots do the work</div>
                  </Link>
                </div>
              )}
            </div>

            {/* Solutions Dropdown */}
            <div className="relative">
              <button
                onClick={() => toggleDropdown('solutions')}
                className="flex items-center space-x-1 px-3 py-2 text-gray-700 hover:text-blue-600 transition"
              >
                <span>Solutions</span>
                <ChevronDown className="w-4 h-4" />
              </button>
              {openDropdown === 'solutions' && (
                <div className="absolute top-full left-0 mt-2 w-64 bg-white rounded-lg shadow-xl border border-gray-200 py-2 z-50">
                  <Link to="/solutions/marketing" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">Marketing Teams</div>
                    <div className="text-sm text-gray-600">Collaborate on campaigns</div>
                  </Link>
                  <Link to="/solutions/product" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">Product Management</div>
                    <div className="text-sm text-gray-600">Track product roadmap</div>
                  </Link>
                  <Link to="/solutions/engineering" className="block px-4 py-3 hover:bg-blue-50 transition" onClick={() => setOpenDropdown(null)}>
                    <div className="font-medium text-gray-900">Engineering Teams</div>
                    <div className="text-sm text-gray-600">Ship better code</div>
                  </Link>
                </div>
              )}
            </div>

            {/* Plans Dropdown */}
            <div className="relative">
              <button
                onClick={() => toggleDropdown('plans')}
                className="flex items-center space-x-1 px-3 py-2 text-gray-700 hover:text-blue-600 transition"
              >
                <span>Plans</span>
                <ChevronDown className="w-4 h-4" />
              </button>
              {openDropdown === 'plans' && (
                <div className="absolute top-full left-0 mt-2 w-48 bg-white rounded-lg shadow-xl border border-gray-200 py-2 z-50">
                  <Link to="/pricing" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Free</Link>
                  <Link to="/pricing" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Standard</Link>
                  <Link to="/pricing" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Premium</Link>
                  <Link to="/pricing" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Enterprise</Link>
                </div>
              )}
            </div>

            {/* Pricing */}
            <Link to="/pricing" className="px-3 py-2 text-gray-700 hover:text-blue-600 transition">
              Pricing
            </Link>

            {/* Resources Dropdown */}
            <div className="relative">
              <button
                onClick={() => toggleDropdown('resources')}
                className="flex items-center space-x-1 px-3 py-2 text-gray-700 hover:text-blue-600 transition"
              >
                <span>Resources</span>
                <ChevronDown className="w-4 h-4" />
              </button>
              {openDropdown === 'resources' && (
                <div className="absolute top-full left-0 mt-2 w-48 bg-white rounded-lg shadow-xl border border-gray-200 py-2 z-50">
                  <Link to="/blog" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Blog</Link>
                  <Link to="/guides" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Guides</Link>
                  <Link to="/webinars" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Webinars</Link>
                  <Link to="/help" className="block px-4 py-3 hover:bg-blue-50 transition text-gray-900" onClick={() => setOpenDropdown(null)}>Help Center</Link>
                </div>
              )}
            </div>
          </div>

          {/* Auth Buttons */}
          <div className="flex items-center space-x-4">
            <Link
              to="/login"
              className="text-blue-600 hover:text-blue-700 font-medium transition"
            >
              Log in
            </Link>
            <Link
              to="/register"
              className="bg-gradient-to-r from-blue-600 to-blue-700 text-white px-5 py-2.5 rounded-md hover:from-blue-700 hover:to-blue-800 transition font-medium shadow-md hover:shadow-lg"
            >
              Get FlowBoard for free
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;