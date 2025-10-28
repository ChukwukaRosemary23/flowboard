import { Link } from 'react-router-dom';
import { ArrowLeft, Code, Bug, Rocket, GitMerge } from 'lucide-react';
import Navbar from '../../components/Navbar';

export default function EngineeringTeams() {
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
            FlowBoard for Engineering Teams
          </h1>
          <p className="text-xl text-gray-600 max-w-3xl mx-auto">
            Ship better code faster with agile workflows, sprint planning, and seamless team coordination
          </p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Code className="w-8 h-8 text-blue-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Sprint Management</h3>
            <p className="text-gray-600">Plan and track sprints with customizable workflows</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Bug className="w-8 h-8 text-purple-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Bug Tracking</h3>
            <p className="text-gray-600">Identify, prioritize, and resolve bugs efficiently</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <GitMerge className="w-8 h-8 text-green-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Code Reviews</h3>
            <p className="text-gray-600">Coordinate pull requests and code reviews</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Rocket className="w-8 h-8 text-orange-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Release Planning</h3>
            <p className="text-gray-600">Manage deployments and release schedules</p>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-xl p-12">
          <h2 className="text-3xl font-bold text-gray-900 mb-8 text-center">Engineering Workflow with FlowBoard</h2>
          <div className="space-y-8">
            <div className="flex items-start">
              <div className="bg-blue-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-blue-600">1</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Sprint Board Setup</h3>
                <p className="text-gray-600">Create boards with lists for Backlog, To Do, In Progress, Code Review, Testing, and Done.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-purple-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-purple-600">2</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Task & Bug Management</h3>
                <p className="text-gray-600">Create cards for features, bugs, and technical debt with labels for priority and type.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-green-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-green-600">3</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Documentation & Assets</h3>
                <p className="text-gray-600">Attach architecture diagrams, API docs, and technical specifications to cards.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-orange-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-orange-600">4</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Real-Time Collaboration</h3>
                <p className="text-gray-600">See updates instantly as teammates move cards and add comments during standups.</p>
              </div>
            </div>
          </div>
        </div>

        <div className="mt-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg p-12 text-center text-white">
          <h2 className="text-3xl font-bold mb-4">Build software that scales</h2>
          <p className="text-xl mb-8">Trusted by engineering teams worldwide</p>
          <Link
            to="/register"
            className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-bold text-lg hover:bg-gray-100 transition"
          >
            Start Building Today
          </Link>
        </div>
      </div>
    </div>
  );
}
