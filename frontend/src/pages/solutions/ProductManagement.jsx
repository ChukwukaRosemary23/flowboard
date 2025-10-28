import { Link } from 'react-router-dom';
import { ArrowLeft, Layers, GitBranch, Zap, CheckCircle } from 'lucide-react';
import Navbar from '../../components/Navbar';

export default function ProductManagement() {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <Link to="/" className="inline-flex items-center text-blue-600 hover:text-blue-700 mb-8">
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Home
        </Link>

        <div className="text-center mb-16">
          <h1 className="text-5xl font-bold text-gray-900 mb-6">
            FlowBoard for Product Management
          </h1>
          <p className="text-xl text-gray-600 max-w-3xl mx-auto">
            Build better products with clear roadmaps, prioritized backlogs, and seamless team collaboration
          </p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Layers className="w-8 h-8 text-blue-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Product Roadmap</h3>
            <p className="text-gray-600">Visualize your product strategy and timeline</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <GitBranch className="w-8 h-8 text-purple-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Feature Tracking</h3>
            <p className="text-gray-600">Track features from idea to launch</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Zap className="w-8 h-8 text-green-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Sprint Planning</h3>
            <p className="text-gray-600">Plan and manage agile sprints efficiently</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <CheckCircle className="w-8 h-8 text-orange-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Release Management</h3>
            <p className="text-gray-600">Coordinate product releases across teams</p>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-xl p-12">
          <h2 className="text-3xl font-bold text-gray-900 mb-8 text-center">Product Management Workflow</h2>
          <div className="space-y-8">
            <div className="flex items-start">
              <div className="bg-blue-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-blue-600">1</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Roadmap Board</h3>
                <p className="text-gray-600">Create quarterly boards with lists for Now, Next, Later, and Icebox features.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-purple-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-purple-600">2</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">User Stories & Requirements</h3>
                <p className="text-gray-600">Document user stories, acceptance criteria, and technical requirements in cards.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-green-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-green-600">3</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Cross-Team Collaboration</h3>
                <p className="text-gray-600">Work with engineering, design, and marketing teams in one place.</p>
              </div>
            </div>
          </div>
        </div>

        <div className="mt-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg p-12 text-center text-white">
          <h2 className="text-3xl font-bold mb-4">Ship products customers love</h2>
          <p className="text-xl mb-8">Join product teams building with FlowBoard</p>
          <Link
            to="/register"
            className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-bold text-lg hover:bg-gray-100 transition"
          >
            Get Started Free
          </Link>
        </div>
      </div>
    </div>
  );
}