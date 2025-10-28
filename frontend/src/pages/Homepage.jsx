import { Link } from 'react-router-dom';
import Navbar from '../components/Navbar';
import { CheckCircle, Users, Zap, Layout, Calendar, Bell, Sparkles, Target, TrendingUp } from 'lucide-react';
import { useEffect, useState } from 'react';

function Homepage() {
  const [isVisible, setIsVisible] = useState(false);

  useEffect(() => {
    // Trigger animations on page load
    setIsVisible(true);
  }, []);

  return (
    <div className="min-h-screen bg-white">
      <Navbar />

      {/* Hero Section - WITH ANIMATIONS */}
      <section className="relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-blue-50 via-purple-50 to-pink-50"></div>
        <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
          <div className="grid md:grid-cols-2 gap-16 items-center">
            
            {/* Left: Text - FADE IN FROM LEFT */}
            <div 
              className={`space-y-8 transition-all duration-1000 ${
                isVisible ? 'opacity-100 translate-x-0' : 'opacity-0 -translate-x-20'
              }`}
            >
              <h1 className="text-5xl md:text-6xl font-bold text-gray-900 leading-tight">
                Capture, organize, and tackle your{' '}
                <span className="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
                  to-dos
                </span>{' '}
                from anywhere.
              </h1>
              <p className="text-xl text-gray-600 leading-relaxed">
                Escape the clutter and chaos—unleash your productivity with FlowBoard.
              </p>
              <Link
                to="/register"
                className="inline-block bg-gradient-to-r from-blue-600 to-blue-700 text-white px-8 py-4 rounded-lg text-lg font-semibold hover:from-blue-700 hover:to-blue-800 transition shadow-lg hover:shadow-xl transform hover:-translate-y-1 hover:scale-105"
              >
                Sign up - it's free!
              </Link>
            </div>

            {/* Right: Phone/Visual - SLIDE IN FROM RIGHT */}
            <div 
              className={`relative transition-all duration-1000 delay-300 ${
                isVisible ? 'opacity-100 translate-x-0' : 'opacity-0 translate-x-20'
              }`}
            >
              <div className="bg-gradient-to-br from-blue-500 via-purple-600 to-pink-500 rounded-2xl p-1 shadow-2xl transform hover:scale-105 hover:rotate-1 transition duration-500">
                <div className="bg-white rounded-xl p-6">
                  <div className="flex items-center justify-between mb-6">
                    <h3 className="text-lg font-bold text-gray-900">📬 Inbox</h3>
                    <span className="text-sm text-gray-500">3 items</span>
                  </div>
                  
                  {/* Cards SLIDE UP one by one */}
                  <div className="space-y-3">
                    <div className={`bg-gradient-to-r from-blue-50 to-blue-100 p-4 rounded-lg border-l-4 border-blue-600 hover:shadow-md transition cursor-pointer transform ${
                      isVisible ? 'translate-y-0 opacity-100' : 'translate-y-10 opacity-0'
                    } transition-all duration-500 delay-500`}>
                      <div className="flex items-center justify-between">
                        <div>
                          <p className="font-semibold text-gray-900">Finish project proposal</p>
                          <p className="text-sm text-gray-600 mt-1">Due today</p>
                        </div>
                        <div className="w-6 h-6 border-2 border-gray-300 rounded"></div>
                      </div>
                    </div>

                    <div className={`bg-gradient-to-r from-purple-50 to-purple-100 p-4 rounded-lg border-l-4 border-purple-600 hover:shadow-md transition cursor-pointer transform ${
                      isVisible ? 'translate-y-0 opacity-100' : 'translate-y-10 opacity-0'
                    } transition-all duration-500 delay-700`}>
                      <div className="flex items-center justify-between">
                        <div>
                          <p className="font-semibold text-gray-900">Review marketing deck</p>
                          <p className="text-sm text-gray-600 mt-1">Tomorrow</p>
                        </div>
                        <div className="w-6 h-6 border-2 border-gray-300 rounded"></div>
                      </div>
                    </div>

                    <div className={`bg-gradient-to-r from-green-50 to-green-100 p-4 rounded-lg border-l-4 border-green-600 hover:shadow-md transition cursor-pointer transform ${
                      isVisible ? 'translate-y-0 opacity-60' : 'translate-y-10 opacity-0'
                    } transition-all duration-500 delay-900`}>
                      <div className="flex items-center justify-between">
                        <div>
                          <p className="font-semibold text-gray-900 line-through">Team meeting</p>
                          <p className="text-sm text-gray-600 mt-1">Completed</p>
                        </div>
                        <div className="w-6 h-6 bg-green-600 rounded flex items-center justify-center">
                          <CheckCircle className="w-4 h-4 text-white" />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              {/* Floating elements - ANIMATED */}
              <div className="absolute -top-4 -right-4 bg-yellow-400 rounded-full p-3 shadow-lg animate-bounce">
                <Sparkles className="w-6 h-6 text-white" />
              </div>
              <div className={`absolute -bottom-4 -left-4 bg-pink-500 rounded-full p-3 shadow-lg transition-all duration-1000 delay-1000 ${
                isVisible ? 'scale-100 rotate-0' : 'scale-0 -rotate-180'
              }`}>
                <Target className="w-6 h-6 text-white" />
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Company Logos Section */}
      <section className="py-12 bg-white border-y border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <p className="text-center text-gray-600 mb-8 font-medium">Join over 2,000,000 teams worldwide that are using FlowBoard</p>
          <div className="grid grid-cols-2 md:grid-cols-6 gap-8 items-center opacity-60 grayscale hover:grayscale-0 transition duration-500">
            <div className="text-center hover:scale-110 transition">
              <div className="text-3xl font-bold text-gray-800">VISA</div>
            </div>
            <div className="text-center hover:scale-110 transition">
              <div className="text-2xl font-bold text-gray-800">coinbase</div>
            </div>
            <div className="text-center hover:scale-110 transition">
              <div className="text-2xl font-bold text-blue-600">John Deere</div>
            </div>
            <div className="text-center hover:scale-110 transition">
              <div className="text-2xl font-bold text-blue-500">zoom</div>
            </div>
            <div className="text-center hover:scale-110 transition">
              <div className="text-xl font-serif text-gray-800">GRAND HYATT</div>
            </div>
            <div className="text-center hover:scale-110 transition">
              <div className="text-2xl font-script text-gray-800">Fender</div>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section id="features" className="py-24 bg-gradient-to-b from-white to-blue-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
              Explore the features that help your team succeed
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Everything you need to organize work, track progress, and collaborate with your team
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            {/* Feature cards */}
            <FeatureCard
              icon={<Bell className="w-6 h-6" />}
              title="Inbox"
              description="Capture every idea. Saved from emails, Slack messages, and turning into tasks that appear in your FlowBoard inbox."
              gradient="from-blue-500 to-blue-600"
              bgGradient="from-blue-50 to-blue-100"
              link="/features/inbox"
            />
            <FeatureCard
              icon={<Calendar className="w-6 h-6" />}
              title="Planner"
              description="Sync your calendar and allocate your time to tasks with productivity focus."
              gradient="from-purple-500 to-purple-600"
              bgGradient="from-purple-50 to-purple-100"
              link="/features/planner"
            />
            <FeatureCard
              icon={<Zap className="w-6 h-6" />}
              title="Automation"
              description="Automate tasks and workflows with Butler. No-code automation is built into every board."
              gradient="from-yellow-500 to-yellow-600"
              bgGradient="from-yellow-50 to-yellow-100"
              link="/features"
            />
            <FeatureCard
              icon={<Layout className="w-6 h-6" />}
              title="Power-Ups"
              description="Power up your workflow by linking other tools directly to your FlowBoard boards."
              gradient="from-green-500 to-green-600"
              bgGradient="from-green-50 to-green-100"
              link="/features"
            />
            <FeatureCard
              icon={<CheckCircle className="w-6 h-6" />}
              title="Templates"
              description="Give your team a blueprint for success with easy-to-use templates from industry leaders."
              gradient="from-pink-500 to-pink-600"
              bgGradient="from-pink-50 to-pink-100"
              link="/features"
            />
            <FeatureCard
              icon={<Users className="w-6 h-6" />}
              title="Integrations"
              description="Find the apps your team is already using or discover new ways to get work done in FlowBoard."
              gradient="from-indigo-500 to-indigo-600"
              bgGradient="from-indigo-50 to-indigo-100"
              link="/integrations"
            />
          </div>
        </div>
      </section>

      {/* Integrations/Extensions Section - FIXED WITH CLICKABLE LINKS! */}
      <section id="integrations" className="py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
              Work smarter with powerful integrations
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Connect FlowBoard with the tools you already use to streamline your workflow
            </p>
          </div>

          <div className="grid md:grid-cols-4 gap-6">
            {/* Integration Cards - NOW WITH WORKING URLS! */}
            <IntegrationCard name="Slack" icon="💬" color="from-purple-500 to-pink-500" url="https://slack.com" />
            <IntegrationCard name="Google Drive" icon="📁" color="from-blue-500 to-green-500" url="https://drive.google.com" />
            <IntegrationCard name="Jira" icon="🎯" color="from-blue-600 to-blue-700" url="https://www.atlassian.com/software/jira" />
            <IntegrationCard name="GitHub" icon="🐙" color="from-gray-700 to-gray-900" url="https://github.com" />
            <IntegrationCard name="Dropbox" icon="📦" color="from-blue-400 to-blue-600" url="https://www.dropbox.com" />
            <IntegrationCard name="Microsoft Teams" icon="👥" color="from-purple-600 to-blue-600" url="https://www.microsoft.com/en/microsoft-teams" />
            <IntegrationCard name="Zoom" icon="📹" color="from-blue-500 to-blue-600" url="https://zoom.us" />
            <IntegrationCard name="Gmail" icon="✉️" color="from-red-500 to-red-600" url="https://mail.google.com" />
          </div>

          <div className="text-center mt-12">
            <Link
              to="/integrations"
              className="inline-block bg-gradient-to-r from-blue-600 to-purple-600 text-white px-8 py-4 rounded-lg text-lg font-semibold hover:from-blue-700 hover:to-purple-700 transition shadow-lg hover:shadow-xl transform hover:-translate-y-1"
            >
              Explore all integrations →
            </Link>
          </div>
        </div>
      </section>

      {/* Productivity Powerhouse Section */}
      <section className="py-24 bg-gradient-to-br from-blue-600 via-purple-600 to-pink-600 relative overflow-hidden">
        <div className="absolute inset-0 opacity-10">
          <div className="absolute top-0 left-0 w-96 h-96 bg-white rounded-full blur-3xl animate-pulse"></div>
          <div className="absolute bottom-0 right-0 w-96 h-96 bg-white rounded-full blur-3xl animate-pulse"></div>
        </div>
        <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-4xl md:text-5xl font-bold text-white mb-6">
            Your productivity powerhouse
          </h2>
          <p className="text-xl text-blue-100 mb-12 max-w-3xl mx-auto leading-relaxed">
            Stay organized and efficient with Inbox, Boards, and Planner. Every to-do, idea, or responsibility—no matter how small—finds its place, keeping you at the top of your game.
          </p>
          <div className="grid md:grid-cols-3 gap-8 mt-16">
            <div className="bg-white/10 backdrop-blur-lg rounded-xl p-8 hover:bg-white/20 transition transform hover:scale-105 cursor-pointer">
              <div className="text-white mb-4">
                <TrendingUp className="w-12 h-12 mx-auto" />
              </div>
              <h3 className="text-xl font-bold text-white mb-2">Inbox</h3>
              <p className="text-blue-100">When it's on your mind, it goes in your Inbox</p>
            </div>
            <div className="bg-white/10 backdrop-blur-lg rounded-xl p-8 hover:bg-white/20 transition transform hover:scale-105 cursor-pointer">
              <div className="text-white mb-4">
                <Layout className="w-12 h-12 mx-auto" />
              </div>
              <h3 className="text-xl font-bold text-white mb-2">Boards</h3>
              <p className="text-blue-100">Your to-do list may be long, but it can be manageable</p>
            </div>
            <div className="bg-white/10 backdrop-blur-lg rounded-xl p-8 hover:bg-white/20 transition transform hover:scale-105 cursor-pointer">
              <div className="text-white mb-4">
                <Calendar className="w-12 h-12 mx-auto" />
              </div>
              <h3 className="text-xl font-bold text-white mb-2">Planner</h3>
              <p className="text-blue-100">See your day clearly with time blocking</p>
            </div>
          </div>
        </div>
      </section>

      {/* Pricing Preview Section */}
      <section id="pricing" className="py-24 bg-gradient-to-b from-blue-50 to-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
              Choose the plan that's right for you
            </h2>
            <p className="text-xl text-gray-600">Start free, upgrade when you need more</p>
          </div>

          <div className="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto">
            {/* Free Plan */}
            <div className="bg-white rounded-2xl p-8 shadow-lg hover:shadow-2xl transition border-2 border-gray-200 hover:border-blue-500 transform hover:-translate-y-2">
              <h3 className="text-2xl font-bold text-gray-900 mb-2">Free</h3>
              <p className="text-4xl font-bold text-gray-900 mb-6">#0<span className="text-lg text-gray-600">/month</span></p>
              <ul className="space-y-3 mb-8">
                <li className="flex items-center text-gray-700">
                  <CheckCircle className="w-5 h-5 text-green-500 mr-2" />
                  10 boards per workspace
                </li>
                <li className="flex items-center text-gray-700">
                  <CheckCircle className="w-5 h-5 text-green-500 mr-2" />
                  Unlimited cards
                </li>
                <li className="flex items-center text-gray-700">
                  <CheckCircle className="w-5 h-5 text-green-500 mr-2" />
                  Up to 10 team members
                </li>
              </ul>
              <Link to="/register" className="block w-full text-center bg-gray-200 text-gray-900 py-3 rounded-lg font-semibold hover:bg-gray-300 transition">
                Get started
              </Link>
            </div>

            {/* Premium Plan */}
            <div className="bg-gradient-to-br from-blue-600 to-purple-600 rounded-2xl p-8 shadow-2xl transform scale-105 relative">
              <div className="absolute -top-4 left-1/2 transform -translate-x-1/2 bg-yellow-400 text-gray-900 px-4 py-1 rounded-full text-sm font-bold">
                POPULAR
              </div>
              <h3 className="text-2xl font-bold text-white mb-2">Premium</h3>
              <p className="text-4xl font-bold text-white mb-6">#10000<span className="text-lg text-blue-100">/month</span></p>
              <ul className="space-y-3 mb-8">
                <li className="flex items-center text-white">
                  <CheckCircle className="w-5 h-5 text-yellow-300 mr-2" />
                  Unlimited boards
                </li>
                <li className="flex items-center text-white">
                  <CheckCircle className="w-5 h-5 text-yellow-300 mr-2" />
                  Advanced features
                </li>
                <li className="flex items-center text-white">
                  <CheckCircle className="w-5 h-5 text-yellow-300 mr-2" />
                  Priority support
                </li>
              </ul>
              <Link to="/register" className="block w-full text-center bg-white text-blue-600 py-3 rounded-lg font-semibold hover:bg-gray-100 transition">
                Start free trial
              </Link>
            </div>

            {/* Enterprise Plan */}
            <div className="bg-white rounded-2xl p-8 shadow-lg hover:shadow-2xl transition border-2 border-gray-200 hover:border-purple-500 transform hover:-translate-y-2">
              <h3 className="text-2xl font-bold text-gray-900 mb-2">Enterprise</h3>
              <p className="text-4xl font-bold text-gray-900 mb-6">Custom</p>
              <ul className="space-y-3 mb-8">
                <li className="flex items-center text-gray-700">
                  <CheckCircle className="w-5 h-5 text-green-500 mr-2" />
                  Everything in Premium
                </li>
                <li className="flex items-center text-gray-700">
                  <CheckCircle className="w-5 h-5 text-green-500 mr-2" />
                  Advanced security
                </li>
                <li className="flex items-center text-gray-700">
                  <CheckCircle className="w-5 h-5 text-green-500 mr-2" />
                  Dedicated support
                </li>
              </ul>
              <a href="#" className="block w-full text-center bg-gray-900 text-white py-3 rounded-lg font-semibold hover:bg-gray-800 transition">
                Contact sales
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-24 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Get started with FlowBoard today
          </h2>
          <div className="max-w-md mx-auto">
            <div className="flex gap-3 mb-4">
              <input
                type="email"
                placeholder="Email"
                className="flex-1 px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
              />
              <Link
                to="/register"
                className="bg-gradient-to-r from-blue-600 to-blue-700 text-white px-8 py-3 rounded-lg font-semibold hover:from-blue-700 hover:to-blue-800 transition shadow-lg hover:shadow-xl whitespace-nowrap transform hover:scale-105"
              >
                Sign up - it's free!
              </Link>
            </div>
            <p className="text-sm text-gray-600">
              By entering my email, I acknowledge the Privacy Policy
            </p>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-white py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-4 gap-12 mb-12">
            <div>
              <h3 className="font-bold text-lg mb-4">About FlowBoard</h3>
              <ul className="space-y-2 text-gray-400">
                <li><a href="#" className="hover:text-white transition">What's behind the boards</a></li>
              </ul>
            </div>
            <div>
              <h3 className="font-bold text-lg mb-4">Jobs</h3>
              <ul className="space-y-2 text-gray-400">
                <li><a href="#" className="hover:text-white transition">Learn about open roles</a></li>
              </ul>
            </div>
            <div>
              <h3 className="font-bold text-lg mb-4">Apps</h3>
              <ul className="space-y-2 text-gray-400">
                <li><a href="#" className="hover:text-white transition">Download for Desktop</a></li>
                <li><a href="#" className="hover:text-white transition">Download for Mobile</a></li>
              </ul>
            </div>
            <div>
              <h3 className="font-bold text-lg mb-4">Contact us</h3>
              <ul className="space-y-2 text-gray-400">
                <li><a href="#" className="hover:text-white transition">Need anything? Get in touch</a></li>
              </ul>
            </div>
          </div>
          <div className="border-t border-gray-800 pt-8 flex flex-col md:flex-row justify-between items-center">
            <p className="text-gray-400 text-sm">&copy; 2025 Rosemary. All rights reserved.</p>
            <div className="flex space-x-6 mt-4 md:mt-0">
              <a href="#" className="text-gray-400 hover:text-white transition">Privacy Policy</a>
              <a href="#" className="text-gray-400 hover:text-white transition">Terms</a>
              <a href="#" className="text-gray-400 hover:text-white transition">Cookie Settings</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

// Feature Card Component
function FeatureCard({ icon, title, description, gradient, bgGradient, link }) {
  return (
    <Link 
      to={link}
      className={`block bg-gradient-to-br ${bgGradient} p-8 rounded-2xl hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2 cursor-pointer`}
    >
      <div className={`bg-gradient-to-r ${gradient} w-14 h-14 rounded-xl flex items-center justify-center mb-6 shadow-lg`}>
        <div className="text-white">{icon}</div>
      </div>
      <h3 className="text-xl font-bold text-gray-900 mb-3">{title}</h3>
      <p className="text-gray-600 leading-relaxed">{description}</p>
      <div className="mt-4 text-sm font-semibold text-blue-600 flex items-center">
        Learn more →
      </div>
    </Link>
  );
}

// Integration Card Component - NOW WITH WORKING URLS!
function IntegrationCard({ name, icon, color, url }) {
  const handleClick = () => {
    if (url) {
      window.open(url, '_blank', 'noopener,noreferrer');
    }
  };

  return (
    <div 
      onClick={handleClick}
      className="bg-white rounded-xl p-6 shadow-md hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2 cursor-pointer border-2 border-gray-100 hover:border-blue-500 group"
    >
      <div className={`bg-gradient-to-r ${color} w-16 h-16 rounded-xl flex items-center justify-center mb-4 mx-auto text-3xl group-hover:scale-110 transition-transform`}>
        {icon}
      </div>
      <h3 className="text-center font-bold text-gray-900 mb-2">{name}</h3>
      <p className="text-center text-xs text-blue-600 font-medium">Click to visit →</p>
    </div>
  );
}

export default Homepage;